package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"darkatic-ci/internal/source"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init() {
	db.DB.AutoMigrate(&repository.Repository{})
	db.DB.AutoMigrate(&repository.EnvVars{})
}

func addRepositoryHandler(c *gin.Context) {
	var repo repository.Repository
	if err := c.ShouldBindJSON(&repo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the existing server based on the provided ID
	var existingServer server.RemoteServer
	if err := db.DB.First(&existingServer, repo.Server).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	// Fetch the existing source based on the provided ID
	var existingSource source.Source
	if err := db.DB.First(&existingSource, repo.Source).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Source not found"})
		return
	}

	// Replace the Server and Source fields with the fetched server and source
	repo.Server = existingServer
	repo.Source = existingSource

	// Create the repository
	db.DB.Create(&repo)
	c.JSON(http.StatusOK, gin.H{"message": "Repository added successfully"})
}

func getRepositoriesHandler(c *gin.Context) {
	var repositories []repository.Repository
	db.DB.Preload("EnvVars").Preload("Source").Preload("Server").Find(&repositories)
	c.JSON(http.StatusOK, repositories)
}

func getRepositoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	repository, err := repository.GetRepositoryById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}
	c.JSON(http.StatusOK, repository)
}

func updateRepositoryHandler(c *gin.Context) {
	id := c.Param("id")
	var repo repository.Repository
	if err := db.DB.First(&repo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	var updatedRepository repository.Repository
	if err := c.ShouldBindJSON(&updatedRepository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&repo).Updates(updatedRepository)
	c.JSON(http.StatusOK, gin.H{"message": "Repository updated successfully"})
}

func deleteRepositoryHandler(c *gin.Context) {
	id := c.Param("id")
	var repo repository.Repository
	if err := db.DB.First(&repo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repository not found"})
		return
	}

	db.DB.Delete(&repo)
	c.JSON(http.StatusOK, gin.H{"message": "Repository deleted successfully"})
}
