package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func getServersHandler(c *gin.Context) {
	servers := server.GetServers()
	c.JSON(http.StatusOK, servers)
}

func getServerByHostnameHandler(c *gin.Context) {
	hostname := c.Param("hostname")
	server, err := server.GetServerByHostname(hostname)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, server)
}

func addServerHandler(c *gin.Context) {
	var remoteServer server.RemoteServer
	if err := c.ShouldBindJSON(&remoteServer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&remoteServer)
	c.JSON(http.StatusOK, gin.H{"message": "Server added successfully"})
}
