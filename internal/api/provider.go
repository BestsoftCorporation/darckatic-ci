package api

import (
	"darkatic-ci/internal/db"
	"darkatic-ci/internal/provider"
	"darkatic-ci/internal/source"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getSourceRepositories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("source_id"))
	var src source.Source
	if err := db.DB.First(&src, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Source not found"})
		return
	}

	var hubProvider provider.ProjectProvider

	if src.SourceType == 0 {
		hubProvider = &provider.GitHubProvider{}
	} else if src.SourceType == 1 {
		hubProvider = &provider.GitLabProvider{}
	}

	result, err := hubProvider.FetchProjects(src.Name, src.Token)
	if err != nil {
		fmt.Println("An error occuered: " + err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
