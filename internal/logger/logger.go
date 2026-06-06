package core_logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	*slog.Logger
	file *os.File
}

type loggerContextKey struct{}

var (
	loggerKey = loggerContextKey{}
)

func ContextWithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func NewLogger(config Config) (*Logger, error) {
	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2026-01-02T15-04-05.000000")
	logFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s. log", timestamp),
	)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(config.LogLevel),
	}

	loggerHandler := slog.NewTextHandler(multiWriter, opts)
	logger := slog.New(loggerHandler)
	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

func (l *Logger) Close() {
	if l.file == nil {
		return
	}

	if err := l.file.Close(); err != nil {
		fmt.Println("failed to close log file:", err)
	}
}

func (l *Logger) With(fields ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		file:   l.file,
	}
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(loggerKey).(*Logger)
	if !ok {
		fmt.Println("logger not found in context")
		return &Logger{Logger: slog.New(slog.NewTextHandler(os.Stdout, nil))}
	}

	return logger
}
