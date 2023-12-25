package api

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// CRUD endpoints for Source
	r.POST("/add-source", addSourceHandler)
	r.GET("/get-sources", getSourcesHandler)
	r.GET("/get-source/:id", getSourceByIDHandler)
	r.PUT("/update-source/:id", updateSourceHandler)
	r.DELETE("/delete-source/:id", deleteSourceHandler)

	// Server
	r.POST("/add-server", addServerHandler)
	r.GET("/get-servers", getServersHandler)
	r.GET("/get-server/:hostname", getServerByHostnameHandler)

	// Operations
	r.POST("/add-repository", addRepositoryHandler)
	r.GET("/get-repositories", getRepositoriesHandler)
	r.GET("/get-repository/:id", getRepositoryByIDHandler)
	r.PUT("/update-repository/:id", updateRepositoryHandler)
	r.DELETE("/delete-repository/:id", deleteRepositoryHandler)

	//
	r.POST("/deploy", deployHandler)

	r.Run(":8080")
}
