package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"net/http"
	"os"

	"time"
)

// TODO 要整理整个注册流程  日志理应在 global 里 这里却又重复实现 有问题
// Log 日志接口咯
// var Log = logrus.New()

// Logger logrus gin 中间件
func Logger(log *logrus.Logger) gin.HandlerFunc {
	return LoggerWithWriter(log)
}

// LoggerWithWriter 中间件的实现
func LoggerWithWriter(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		// TODO 这里的时间戳实际上可以合并到 openTraining
		start := time.Now()
		span, ctx := opentracing.StartSpanFromContext(c, c.Request.URL.Path+c.Request.URL.RawQuery)
		defer span.Finish()
		ctx = ctx

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		uuidst:= uuid.NewV4()
		requestID := uuidst.String()
		c.Set("RequestID", requestID)

		var bodyBytes []byte
		if c.Request.Body != nil {
			var err error
			bodyBytes, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.WithFields(logrus.Fields{
					"request_id": requestID,
					"err":        err.Error(),
				}).Warn("请求中的 body 解析错误")
			}
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Use the content
		body := string(bodyBytes)

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		c.Request.ParseForm()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		statusLog := log.WithFields(logrus.Fields{
			"request_id":   requestID,
			"latency":      latency,
			"end":          end.Format("2006/01/02 - 15:04:05"),
			"clientIP":     clientIP,
			"method":       method,
			"path":         path,
			"comment":      comment,
			"statusCode":   statusCode,
			"err":          c.Errors.String(),
			"errors":       c.Errors.Errors(),
			"query":        c.Request.URL.Query(),
			"PostForm":     c.Request.PostForm,
			"body":         string(body),
			"Content-Type": c.Request.Header.Get("Content-Type"),
		})

		// @since 0.0.4 静态文件不打印日志
		// TODO 从配置文件中获取  把中间件移动到 global 中
		mode := "release"
		if mode == "release" {
			if strings.HasPrefix(path, "/js") {
				return
			}
			if strings.HasPrefix(path, "/css") {
				return
			}
		}

		if statusCode == http.StatusOK || statusCode == http.StatusFound {
			statusLog.Info("路由日志")
			return
		}
		statusLog.Error("路由日志")

	}

}

// DefaultLogger 初始化日志
func DefaultLogger() *logrus.Logger {
	Log := logrus.New()
	fmt.Println("[GIN-Mode] " + gin.Mode())
	// 设置日志格式
	if gin.Mode() == gin.ReleaseMode {
		Log.SetLevel(logrus.InfoLevel)
		Log.Formatter = &logrus.JSONFormatter{}
		file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Panic(err)
		}
		Log.Out = file
		Log.Println("log 写出到 log.log")
		Log.Info("logo to log.log")
	} else {
		Log.SetLevel(logrus.DebugLevel)
	}

	return Log
}
