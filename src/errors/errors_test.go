package errors

import (
	"testing"
)

func TestNew(t *testing.T) {
	msg := "test error"
	err := New(msg)
	if err == nil {
		t.Fatal("New() returned nil")
	}
	if err.Error() != msg {
		t.Errorf("Error() = %q, want %q", err.Error(), msg)
	}
}

func TestErrorString(t *testing.T) {
	msg := "another error"
	err := &errorString{errs: msg}
	if err.Error() != msg {
		t.Errorf("Error() = %q, want %q", err.Error(), msg)
	}
}
