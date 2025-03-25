package logs

import (
	"context"
	"fmt"
	"log/slog"
)

const resultSeparator = "----------------------------------------"

func IsDebug() bool {
	return slog.Default().Handler().Enabled(context.Background(), slog.LevelDebug)
}

func Debugln() {
	if IsDebug() {
		fmt.Println()
	}
}

func Debugf(format string, args ...interface{}) {
	if IsDebug() {
		fmt.Printf(format, args...)
	}
}

func ResultSeparator() {
	if IsDebug() {
		fmt.Println(resultSeparator)
	}
}
