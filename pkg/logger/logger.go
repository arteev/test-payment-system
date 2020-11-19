package logger

import (
	"path/filepath"
	"runtime"
	"strings"
	"test-payment-system/internal/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func GetLogLevel(l string) zapcore.Level {
	for level := zapcore.DebugLevel; level < zapcore.FatalLevel; level++ {
		if strings.EqualFold(level.String(), l) {
			return level
		}
	}

	return zapcore.InfoLevel
}

func SetupLogger(mode string, configLogger *config.Logger) error {
	zapConfig := zap.NewProductionConfig()
	if mode == "development" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.DisableStacktrace = false
	}
	level := GetLogLevel(configLogger.Level)
	zapConfig.Level.SetLevel(level)
	zapConfig.EncoderConfig.EncodeCaller = createCallerFunc(5)



	log, err := zapConfig.Build()
	if err != nil {
		return err
	}
	Logger = log.Sugar()
	return nil
}

func SetLogger(logger *zap.SugaredLogger) {
	Logger = logger
}

func getCallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		name := details.Name()
		index := strings.LastIndex(name, ".") + 1
		if index > 0 && index < len(name) {
			return name[index:] // extract only function name
		}
		return name
	}
	return ""
}

func createCallerFunc(skip int) func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	return func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(filepath.Base(caller.FullPath()) + " " + getCallerName(skip))
	}
}


