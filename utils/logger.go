package utils

import (
	"log/slog"
	"github.com/lmittmann/tint"
	"os"
)

func InitLogger(pretty bool) *slog.Logger {
	var handler slog.Handler

	if pretty {
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelInfo,
			TimeFormat: "2006-01-02 15:04:05",
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}


