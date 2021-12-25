package college

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	FormatDateLayout = "2006-01-02 15:04:05"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		start := time.Now()

		ctx.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		referer := ctx.Request.Referer()
		method := ctx.Request.Method

		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		dataLength := ctx.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			"path":       path,
			"latency":    latency,
			"statusCode": statusCode,
			"clientIP":   clientIP,
			"userAgent":  userAgent,
			"referer":    referer,
			"method":     method,
			"hostname":   hostname,
			"dataLength": dataLength,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s] %s %s [%d]", time.Now().Format(FormatDateLayout), ctx.Request.Method, path, statusCode)

			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}

	}
}
