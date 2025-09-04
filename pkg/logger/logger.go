// pkg/logger/logger.go

package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
    var err error
    log, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
    zap.ReplaceGlobals(log)
}

func L() *zap.Logger {
    return log
}

func WithContextFields(c *gin.Context) *zap.Logger {
    return log.With(
        zap.String("request_id", c.GetString("request_id")),
        zap.String("user_id", c.GetString("user_id")),
    )
}
