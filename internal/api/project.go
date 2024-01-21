package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/project"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	db.DB.AutoMigrate(&project.Project{})
}

// createProject creates a new project
func createProject(c *gin.Context) {
	var project project.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the existing server based on the provided ID
	var existingServer server.RemoteServer
	if err := db.DB.First(&existingServer, project.Server.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	// Replace the Server field with the fetched server
	project.Server = existingServer

	// Extract repository IDs from the provided repositories
	var repositoryIDs []uint
	for _, repo := range project.Repository {
		repositoryIDs = append(repositoryIDs, repo.ID)
	}

	// Fetch existing repositories based on the provided IDs
	var existingRepositories []repository.Repository
	if err := db.DB.Where("id IN (?)", repositoryIDs).Find(&existingRepositories).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "One or more repositories not found"})
		return
	}

	// Replace the Repository field with the fetched repositories
	project.Repository = existingRepositories

	db.DB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

// getProject retrieves a project by ID
func getProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var project project.Project

	if err := db.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// getAllProjects retrieves all projects
func getAllProjects(c *gin.Context) {
	var projects []project.Project
	db.DB.Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// updateProject updates a project by ID
func updateProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var project project.Project

	if err := db.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&project)
	c.JSON(http.StatusOK, project)
}

// deleteProject deletes a project by ID
func deleteProject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var project project.Project

	if err := db.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}
