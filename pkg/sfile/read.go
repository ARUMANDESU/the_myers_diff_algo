package sfile

import (
	"context"
	"os"
)

func ReadTextFile(ctx context.Context, path string) ([]rune, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return []rune(string(content)), nil
}
