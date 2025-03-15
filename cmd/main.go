package main

import (
	"flag"
	"log/slog"
)

func main() {
	slog.Info("flag", "args", flag.Args())
}
