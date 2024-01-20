package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/environment"
	"darkatic-ci/internal/project"
	"darkatic-ci/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	db.DB.AutoMigrate(&environment.Environment{})
}

func createEnvironment(c *gin.Context) {
	var environment environment.Environment
	if err := c.ShouldBindJSON(&environment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the existing server based on the provided ID
	var existingServer server.RemoteServer
	if err := db.DB.First(&existingServer, environment.Server.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	// Replace the Server field with the fetched server
	environment.Server = existingServer

	// Extract project IDs from the provided projects
	var projectIDs []uint
	for _, proj := range environment.Project {
		projectIDs = append(projectIDs, proj.ID)
	}

	// Fetch existing projects based on the provided IDs
	var existingProjects []project.Project
	if err := db.DB.Where("id IN (?)", projectIDs).Find(&existingProjects).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "One or more projects not found"})
		return
	}

	// Replace the Project field with the fetched projects
	environment.Project = existingProjects

	db.DB.Create(&environment)
	c.JSON(http.StatusCreated, environment)
}

// getEnvironment retrieves an environment by ID
func getEnvironment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var environment environment.Environment

	if err := db.DB.Preload("Project").Preload("Server").First(&environment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, environment)
}

// getAllEnvironments retrieves all environments
func getAllEnvironments(c *gin.Context) {
	var environments []environment.Environment

	// Retrieve all environments with associated data (Server and Project)
	if err := db.DB.Preload("Server").Preload("Project").Find(&environments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, environments)
}

// updateEnvironment updates an environment by ID
func updateEnvironment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var environment environment.Environment

	if err := db.DB.First(&environment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&environment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Additional logic to handle project structures
	// For example, you can fetch the projects by IDs from the database and assign them to the environment

	db.DB.Save(&environment)
	c.JSON(http.StatusOK, environment)
}

// deleteEnvironment deletes an environment by ID
func deleteEnvironment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var environment environment.Environment

	if err := db.DB.First(&environment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&environment)
	c.JSON(http.StatusOK, gin.H{"message": "Environment deleted successfully"})
}
