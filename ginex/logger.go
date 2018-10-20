package ginex

import (
	"fmt"
	"io"
	"time"

	"github.com/chouandy/goex/httpex"
	"github.com/gin-gonic/gin"
)

var notLoggedPaths map[string]struct{}

// logger logger struct
type logger struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Status    int    `json:"status"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Latency   string `json:"latency"`
	ClientIP  string `json:"client_ip"`
}

// Logger logger
func Logger() gin.HandlerFunc {
	return LoggerWithWriter(gin.DefaultWriter)
}

// LoggerWithNotLogged logger with not logged
func LoggerWithNotLogged(paths ...string) gin.HandlerFunc {
	// Set not logged
	if length := len(paths); length > 0 {
		notLoggedPaths = make(map[string]struct{}, length)
		for _, path := range paths {
			notLoggedPaths[path] = struct{}{}
		}
	}

	return LoggerWithWriter(gin.DefaultWriter)
}

// LoggerWithWriter logger with writer
func LoggerWithWriter(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Not logged
		if _, ok := notLoggedPaths[c.Request.URL.Path]; ok {
			return
		}

		// New logger
		data, err := jsonex.Marshal(&logger{
			Timestamp: start.Format(time.RFC3339),
			Level:     httpex.GetLogLevel(c.Writer.Status()),
			Status:    c.Writer.Status(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Latency:   fmt.Sprintf("%v", time.Now().UTC().Sub(start)),
			ClientIP:  c.ClientIP(),
		})
		// Print
		if err == nil {
			fmt.Fprintln(out, string(data))
		}
	}
}