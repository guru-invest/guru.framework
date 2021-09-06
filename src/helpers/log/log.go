package log

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger
var serviceName string

type GuruLog struct {
	HTTPHeader *http.Header
	/*
		ServiceName   string
		DeviceID      string
		CorrelationID string
		SessionID     string
		UserAgent     map[string]interface{}
	*/
}
type Fields map[string]interface{}

type LogWithFields struct {
	CustomerCode string
	Caller       string
	InfoMessage  Fields
}

func InitLog(pLogLevel string, pServiceName string) {
	var logLevel zapcore.Level
	switch pLogLevel {
	case "INFO":
		logLevel = zapcore.InfoLevel
	case "DEBUG":
		logLevel = zapcore.DebugLevel
	case "ERROR":
		logLevel = zapcore.ErrorLevel
	case "FATAL":
		logLevel = zapcore.FatalLevel
	case "WARN":
		logLevel = zapcore.WarnLevel
	case "PANIC":
		logLevel = zapcore.WarnLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	serviceName = pServiceName
	logger, _ := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()
	sugar = logger.Sugar()

}

func (t GuruLog) Info(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.Infow(message, fields...)
}

func (t GuruLog) Error(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.Errorw(message, fields...)
}

func (t GuruLog) Debug(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.With(fields...).Debugw(message)
}

func (t GuruLog) Fatal(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.Fatalw(message, fields...)
}

func (t GuruLog) Panic(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.Panicw(message, fields...)
}

func (t GuruLog) Warning(pFields *LogWithFields, pMessage string) {
	message, fields := t.createMessage(pFields, pMessage)
	sugar.With(fields...).Warnw(message)
}

func (t GuruLog) createMessage(fields *LogWithFields, message string) (string, []interface{}) {

	var header []interface{}
	if t.HTTPHeader != nil {
		header = []interface{}{
			"device-id", t.HTTPHeader.Get("device-id"),
			"correlation-id", t.HTTPHeader.Get("correlation-id"),
			"session-id", t.HTTPHeader.Get("session-id"),
			"user-agent", t.HTTPHeader.Get("user-agent"),
			"client-ip", t.HTTPHeader.Get("X-Forwarded-For"),
			"customer-code", fields.CustomerCode,
			"service-name", serviceName,
			"caller", fields.Caller,
			"info-message", fields.InfoMessage,
		}
	}

	if fields != nil {
		return message, header
	}

	return message, []interface{}{}
}
