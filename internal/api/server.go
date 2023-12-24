package api

import (
	"darkatic-ci/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB

func init() {
	var err error
	// Replace the connection string with your PostgreSQL database connection details.
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&server.RemoteServer{})
}

func ServerServe() {
	r := gin.Default()
	r.POST("/add-server", addServerHandler)
	r.Run(":8080")
}

func addServerHandler(c *gin.Context) {
	var remoteServer server.RemoteServer
	if err := c.ShouldBindJSON(&remoteServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&remoteServer)
	c.JSON(http.StatusOK, gin.H{"message": "Server added successfully"})
}
