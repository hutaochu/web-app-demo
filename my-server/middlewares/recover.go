package middlewares

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// avoid broken pipe crash here
				// if connect is broken, don't response
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						serr := strings.ToLower(se.Error())
						if strings.Contains(serr, "broken pipe") ||
							strings.Contains(serr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				errResp := fmt.Errorf("panic error: %v", err)

				url := getRequestURL(ctx)

				klog.Infof("http_request, url=%s, err=%+v, brokenPipe=%v", url, errResp, brokenPipe)
			}
		}()

		ctx.Next()
	}
}
