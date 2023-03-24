package logging

import "go.uber.org/zap"

const (
	KeyError = "error"
	KeyID 	 = "id"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
