package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	// CRUD endpoints for Source
	r.POST("/source", addSourceHandler)
	r.GET("/sources", getSourcesHandler)
	r.GET("/source/:id", getSourceByIDHandler)
	r.PUT("/source/:id", updateSourceHandler)
	r.DELETE("/source/:id", deleteSourceHandler)

	// Server
	r.POST("/server", addServerHandler)
	r.GET("/servers", getServersHandler)
	r.GET("/server/:hostname", getServerByHostnameHandler)

	// Operations
	r.POST("/repository", addRepositoryHandler)
	r.GET("/repositories", getRepositoriesHandler)
	r.GET("/repository/:id", getRepositoryByIDHandler)
	r.PUT("/repository/:id", updateRepositoryHandler)
	r.DELETE("/repository/:id", deleteRepositoryHandler)

	//
	r.POST("/deploy", deployHandler)

	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
