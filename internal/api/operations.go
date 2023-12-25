package api

import (
	"darkatic-ci/internal/deploy"
	"darkatic-ci/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deployHandler(c *gin.Context) {
	id := c.Param("id")
	repo, err := repository.GetRepositoryById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = deploy.Deploy(repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repository deployed successfully"})
}
