package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/source"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SourceType represents the source type.
type SourceType int

const (
	GitHub SourceType = iota
	GitLab
)

func addSourceHandler(c *gin.Context) {
	var src source.Source
	if err := c.ShouldBindJSON(&src); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&src)
	c.JSON(http.StatusOK, gin.H{"message": "Source added successfully"})
}

func getSourcesHandler(c *gin.Context) {
	var sources []source.Source
	db.DB.Find(&sources)
	c.JSON(http.StatusOK, sources)
}

func getSourceByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var src source.Source
	if err := db.DB.First(&src, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Source not found"})
		return
	}
	c.JSON(http.StatusOK, src)
}

func updateSourceHandler(c *gin.Context) {
	id := c.Param("id")
	var src source.Source
	if err := db.DB.First(&src, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Source not found"})
		return
	}

	var updatedSource source.Source
	if err := c.ShouldBindJSON(&updatedSource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&src).Updates(updatedSource)
	c.JSON(http.StatusOK, gin.H{"message": "Source updated successfully"})
}

func deleteSourceHandler(c *gin.Context) {
	id := c.Param("id")
	var src source.Source
	if err := db.DB.First(&src, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Source not found"})
		return
	}

	db.DB.Delete(&src)
	c.JSON(http.StatusOK, gin.H{"message": "Source deleted successfully"})
}
