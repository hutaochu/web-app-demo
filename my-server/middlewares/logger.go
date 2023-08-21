package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hutaochu/web-app-demo/myserver/utils"
	"github.com/hutaochu/web-app-demo/myserver/utils/byteutils"
	"k8s.io/klog/v2"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTimestamp := time.Now()
		bodyString := getBodyString(c)
		requestURI := getRequestURL(c)
		traceID := utils.GetTraceID(c)

		c.Next()

		statusCode := c.Writer.Status()
		elapsed := time.Since(startTimestamp)
		klog.Infof("http_request, traceID=%s, uri=%s, requestBody=%s, statusCode=%d, elapsed=%v",
			traceID, requestURI, bodyString, statusCode, elapsed)
	}
}

func getRequestURL(c *gin.Context) string {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	url := path
	if raw != "" {
		url += "?" + raw
	}
	return url
}

func getBodyString(c *gin.Context) string {
	var body []byte
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil && c.Request.Body != nil {
		body, _ = c.GetRawData()
	}

	return byteutils.ToString(body)
}
