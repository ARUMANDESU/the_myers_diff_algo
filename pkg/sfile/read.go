package sfile

import (
	"context"
	"os"
	"strings"
)

func ReadRunes(ctx context.Context, path string) ([]rune, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return []rune(string(content)), nil
}

func ReadLines(ctx context.Context, path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
