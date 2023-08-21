package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hutaochu/web-app-demo/myserver/utils"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		utils.SetTraceID(c, traceID)
	}
}
