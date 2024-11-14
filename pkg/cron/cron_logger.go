package cronJob

import "github.com/diki-haryadi/go-micro-template/pkg/logger"

type CronLogger struct{}

func NewLogger() *CronLogger {
	return &CronLogger{}
}

func (l *CronLogger) Info(msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Infow(msg, keysAndValues)
}

func (l *CronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Zap.Sugar().Errorw(msg, keysAndValues)
}
