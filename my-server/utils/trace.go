package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hutaochu/web-app-demo/myserver/constants"
)

func SetTraceID(c *gin.Context, traceID string) {
	c.Set(constants.TraceIDKey, traceID)
}

func GetTraceID(c *gin.Context) string {
	traceID, ok := c.Get(constants.TraceIDKey)
	if !ok {
		traceID = uuid.New().String()
		c.Set(constants.TraceIDKey, traceID)
	}
	return traceID.(string)
}
