package log

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelInfo  = "INFO"
	LevelDebug = "DEBUG"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
	LevelWarn  = "WARN"
	LevelPanic = "PANIC"
)

type Fields map[string]interface{}
type GuruLogOption func(*GuruLog)

type GuruLog struct {
	HTTPHeader *http.Header
	logger     *zap.SugaredLogger
}

type loggableError struct {
	err        error
	logDetails Fields
}

func (l *loggableError) Error() string {
	return l.err.Error()
}

func NewLoggedError(err error, logDetails Fields) error {
	return &loggableError{
		err:        err,
		logDetails: logDetails,
	}
}

func NewGuruLog(serviceName string, logLevel string, options ...GuruLogOption) *GuruLog {
	level := parseLogLevel(logLevel)
	coreLogger, err := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "severity",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		},
	}.Build()

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	log := &GuruLog{
		logger: coreLogger.Sugar().With("service-name", serviceName),
	}

	for _, opt := range options {
		opt(log)
	}

	return log
}

func WithHTTPHeader(httpHeader http.Header) GuruLogOption {
	httpHeader.Add("correlation-id", uuid.NewString())
	return func(g *GuruLog) {
		g.HTTPHeader = &httpHeader
	}
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelPanic:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func (g *GuruLog) log(level, message string, fields Fields) {
	extraFields := make([]interface{}, 0, len(fields)*2)

	for k, v := range fields {
		extraFields = append(extraFields, k, v)
	}

	switch level {
	case LevelDebug:
		g.logger.Debugw(message, extraFields...)
	case LevelInfo:
		g.logger.Infow(message, extraFields...)
	case LevelError:
		g.logger.Errorw(message, extraFields...)
	case LevelFatal:
		g.logger.Fatalw(message, extraFields...)
	case LevelWarn:
		g.logger.Warnw(message, extraFields...)
	case LevelPanic:
		g.logger.Panicw(message, extraFields...)
	}
}

func (g *GuruLog) Info(message string, fields Fields) {
	g.log(LevelInfo, message, fields)
}

func (g *GuruLog) Debug(message string, fields Fields) {
	g.log(LevelDebug, message, fields)
}

func (g *GuruLog) Error(message string, fields Fields) {
	g.log(LevelError, message, fields)
}
