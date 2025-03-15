package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/diff"
	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/sfile"
)

func main() {
	var isDebug bool

	flag.BoolVar(&isDebug, "d", false, "enable debug mode")

	flag.Parse()

	l := NewLogger(isDebug)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l.Debug("flag", "args", flag.Args())

	var (
		text1 []rune
		text2 []rune
		err   error
	)
	text1, err = sfile.ReadTextFile(ctx, flag.Arg(0))
	if err != nil {
		l.Error("read first file", "err", err)
		os.Exit(1)
	}

	text2, err = sfile.ReadTextFile(ctx, flag.Arg(1))
	if err != nil {
		l.Error("read second file", "err", err)
		os.Exit(1)
	}

	l.Debug("input", "text1", string(text1), "text2", string(text2))

	edits := diff.Myers(text1, text2)

	fmt.Print(edits)
}

func NewLogger(isDebugMode bool) *slog.Logger {
	opts := slog.HandlerOptions{}
	if isDebugMode {
		opts.Level = slog.LevelDebug
	}

	l := slog.New(slog.NewTextHandler(os.Stdin, &opts))
	slog.SetDefault(l)

	return l
}
