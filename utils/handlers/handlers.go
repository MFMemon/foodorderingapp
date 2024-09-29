package handlers

import (
	"context"
	"log/slog"
)

func HandleErr(err error, level slog.Level) {
	if err != nil {
		slog.Log(context.Background(), level, err.Error(), err)
	}
}
