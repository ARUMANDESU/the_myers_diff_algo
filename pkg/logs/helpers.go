package logs

import (
	"context"
	"fmt"
	"log/slog"
)

const resultSeparator = "----------------------------------------"

func isDebug() bool {
	return slog.Default().Handler().Enabled(context.Background(), slog.LevelDebug)
}

func Debugln() {
	if isDebug() {
		fmt.Println()
	}
}

func Debugf(format string, args ...interface{}) {
	if isDebug() {
		fmt.Printf(format, args...)
	}
}

func ResultSeparator() {
	if isDebug() {
		fmt.Println(resultSeparator)
	}
}
