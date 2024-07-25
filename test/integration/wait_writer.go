package integration

import (
	"bytes"
	"io"
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
