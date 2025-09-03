package v2

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
    r.GET("/status", func(c *gin.Context) {
        c.JSON(200, gin.H{"version": "v2", "status": "ok"})
    })
}
