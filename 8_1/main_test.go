package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestMainPrintsGreeting is the smoke test this repository had none of: it
// captures what main() writes to standard output and fails if it stops
// saying "Hello World!".
func TestMainPrintsGreeting(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	want := "Hello World!"
	if got != want {
		t.Fatalf("main() printed %q, want %q", got, want)
	}
}
