package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func addRepositoryHandler(c *gin.Context) {
	var repository repository.Repository
	if err := c.ShouldBindJSON(&repository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&repository)
	c.JSON(http.StatusOK, gin.H{"message": "Repository added successfully"})
}

func getRepositoriesHandler(c *gin.Context) {
	var repositories []repository.Repository
	db.DB.Preload("Source").Find(&repositories)
	c.JSON(http.StatusOK, repositories)
}

func getRepositoryByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var repository repository.Repository
	if err := db.DB.Preload("Source").First(&repository, id).Error; err != nil {
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
