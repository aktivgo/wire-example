package dependency

import (
	"wire/pkg/logging"
)

func ProvideLogger() logging.Logger {
	return logging.NewLogger().WithLevel(logging.LevelInfo)
}
