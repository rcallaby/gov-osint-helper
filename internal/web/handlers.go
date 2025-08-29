package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gov-osint-tool/internal/search"
)

func DashboardHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func SearchHandler(c *gin.Context) {
	query := c.PostForm("query")

	results := search.RunQuery(query) // This should use your OSINT logic
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Results": results,
		"Query":   query,
	})
}
