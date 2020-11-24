package service

import (
	"fmt"
	"io"
	"os"
	"test-payment-system/pkg/requestid"
	"time"

	"go.uber.org/zap"

	"github.com/gorilla/handlers"

	"net/http"
)

// RequestScope middleware request context initialization
func RequestScope(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := requestid.NewOrFromRequest(r)
		ctx := requestID.ToContext(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func logFormatter(log *zap.SugaredLogger) handlers.LogFormatter {
	return func(writer io.Writer, params handlers.LogFormatterParams) {
		duration := time.Since(params.TimeStamp)
		remoteIP := params.Request.Header.Get("X-Real-Ip")
		if remoteIP == "" {
			remoteIP = params.Request.RemoteAddr
		}
		log.Debugw(
			fmt.Sprintf("%s %s %d %s", duration, params.Request.Method, params.StatusCode, params.URL.Path),
			zap.String("remote_ip", remoteIP),
			zap.String("method", params.Request.Method),
			zap.Int("status_code", params.StatusCode),
			zap.String("path", params.URL.Path),
			zap.Duration("duration", duration),
		)
	}
}

func loggingHandler(log *zap.SugaredLogger, h http.Handler) http.Handler {
	return handlers.CustomLoggingHandler(os.Stdout, h, logFormatter(log))
}
