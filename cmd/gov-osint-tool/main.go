package main

import (
	"github.com/gin-gonic/gin"
	"gov-osint-tool/internal/web"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./webui/static")
	r.LoadHTMLGlob("webui/templates/*")

	// Routes
	r.GET("/", web.DashboardHandler)
	r.POST("/search", web.SearchHandler)

	// Start server
	r.Run(":8080")
}
