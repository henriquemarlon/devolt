package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"
	"github.com/Khan/genqlient/graphql"
)

type NotifyWriter struct {
	io.Writer
	ready   chan struct{}
	lookFor []byte
	found   bool
}

func NewNotifyWriter(w io.Writer, lookFor string) *NotifyWriter {
	return &NotifyWriter{
		Writer:  w,
		ready:   make(chan struct{}, 1),
		lookFor: []byte(lookFor),
	}
}

func (w *NotifyWriter) Ready() <-chan struct{} {
	return w.ready
}

func (w *NotifyWriter) Write(b []byte) (int, error) {
	if !w.found && bytes.Contains(b, w.lookFor) {
		w.found = true
		w.ready <- struct{}{}
	}
	return w.Writer.Write(b)
}

func WaitForInput(ctx context.Context, client graphql.Client) error {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	for {
		result, err := getInputStatus(ctx, client, 0)
		if err != nil && !strings.Contains(err.Error(), "input not found") {
			return fmt.Errorf("failed to get input status: %w", err)
		}
		if result.Input.Status == CompletionStatusAccepted {
			return nil
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
