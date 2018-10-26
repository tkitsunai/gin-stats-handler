package gin_stats_handler

import (
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
)

func Stats() func(g *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, stats_api.GetStats())
	}
}
