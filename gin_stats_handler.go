package gin_stats_handler

import (
	"github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Stats() func(g *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, stats_api.GetStats())
	}
}
