package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/project"
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
