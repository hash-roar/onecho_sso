package logging

import (
	"context"
	"onecho_sso_backend/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type Logger struct {
	logger *zap.Logger
	fields map[string]string
	ctx    context.Context
}

var innerLogger *Logger

func init() {
	innerLogger = &Logger{
		ctx:    context.Background(),
		fields: make(map[string]string),
		logger: zap.NewExample(),
	}
}

func Ctx(ctx context.Context) *Logger {
	return innerLogger.Ctx(ctx)
}

func (log *Logger) Ctx(ctx context.Context) *Logger {
	l := log.With(log.DecodeCtx(ctx)...)
	l.ctx = ctx
	return l
}

func (log *Logger) clone() *Logger {
	l := *log.logger
	f := make(map[string]string)
	for i, v := range log.fields {
		f[i] = v
	}
	return &Logger{
		logger: &l,
		fields: f,
		ctx:    log.ctx,
	}

}

func (log *Logger) DecodeCtx(ctx context.Context) []zap.Field {
	var f []zap.Field
	reqctx := ctx.Value("reqctx").(utils.RequestContext)
	f = append(f, zap.String("RequestId", reqctx.RequestId))
	return f
}

func (log *Logger) With(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return log
	}
	l := log.clone()
	l.logger = l.logger.With(fields...)
	return l
}

func (log *Logger) Debug(msg string, fields ...zap.Field) {
	log.logger.Debug(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	innerLogger.Debug(msg, fields...)
}

func (log *Logger) Debugf(template string, args ...interface{}) {
	log.logger.Sugar().Debugf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	innerLogger.Debugf(template, args...)
}

func (log *Logger) Info(msg string, fields ...zap.Field) {
	log.logger.Info(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	innerLogger.Info(msg, fields...)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	log.logger.Sugar().Infof(template, args...)
}

func Infof(template string, args ...interface{}) {
	innerLogger.Infof(template, args...)
}

func (log *Logger) Warn(msg string, fields ...zap.Field) {
	log.logger.Warn(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	innerLogger.Warn(msg, fields...)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	log.logger.Sugar().Warnf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	innerLogger.Warnf(template, args...)
}

func (log *Logger) Error(msg string, fields ...zap.Field) {
	log.logger.Error(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	innerLogger.Error(msg, fields...)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	log.logger.Sugar().Errorf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	innerLogger.Errorf(template, args...)
}
