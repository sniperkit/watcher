package status

import (
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func JobHtml(c *gin.Context) {
	c.HTML(200, "", jobrunner.StatusPage())
}
