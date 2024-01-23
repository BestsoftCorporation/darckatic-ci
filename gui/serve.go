package gui

import (
	"github.com/gin-gonic/gin"
)

func ServeGUI() {
	router := gin.Default()
	router.Static("/", "./gui/static")
	router.NoRoute(func(c *gin.Context) {
		c.File("./gui/static/index.html")
	})
	router.Run(":8082")
}
