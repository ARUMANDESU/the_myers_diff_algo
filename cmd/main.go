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
	var (
		isDebug  bool
		isByLine bool
	)

	flag.BoolVar(&isDebug, "d", false, "enable debug mode")
	flag.BoolVar(&isByLine, "l", false, "diff by line")

	flag.Parse()

	l := newLogger(isDebug)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l.Debug("flag", "args", flag.Args())

	if isByLine {
		diffByLine(ctx, l, flag.Arg(0), flag.Arg(1))
	} else {
		diffByRune(ctx, l, flag.Arg(0), flag.Arg(1))
	}
}

func diffByRune(ctx context.Context, l *slog.Logger, file1, file2 string) {
	var (
		text1 []rune
		text2 []rune
		err   error
	)
	text1, err = sfile.ReadRunes(ctx, file1)
	if err != nil {
		l.Error("read first file", "err", err)
		os.Exit(1)
	}

	text2, err = sfile.ReadRunes(ctx, file2)
	if err != nil {
		l.Error("read second file", "err", err)
		os.Exit(1)
	}

	l.Debug("input", "text1", string(text1), "text2", string(text2))

	edits := diff.Myers(text1, text2)
	fmt.Print(edits)
}

func diffByLine(ctx context.Context, l *slog.Logger, file1, file2 string) {
	var (
		text1 []string
		text2 []string
		err   error
	)

	text1, err = sfile.ReadLines(ctx, file1)
	if err != nil {
		l.Error("read first file", "err", err)
		os.Exit(1)
	}

	text2, err = sfile.ReadLines(ctx, file2)
	if err != nil {
		l.Error("read second file", "err", err)
		os.Exit(1)
	}

	l.Debug("input", "text1", text1, "text2", text2)

	edits := diff.MyersByline(text1, text2)
	fmt.Print(edits)
}

func newLogger(isDebugMode bool) *slog.Logger {
	opts := slog.HandlerOptions{}
	if isDebugMode {
		opts.Level = slog.LevelDebug
	}

	l := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(l)

	return l
}
