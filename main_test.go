package main

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	args := []string{
		"--length=20",
		"--digits",
		"--symbols",
		"--uppercase",
	}

	config, err := parseFlags(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if config.Length != 20 {
		t.Errorf("expected length 20, got %d", config.Length)
	}
	if !config.Digits {
		t.Errorf("expected digits to be true")
	}
	if !config.Symbols {
		t.Errorf("expected symbols to be true")
	}
	if !config.Uppercase {
		t.Errorf("expected uppercase to be true")
	}
}
