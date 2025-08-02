package main

import (
	"testing"
)

func TestParseFlags_ParsesLength(t *testing.T) {
	// Arrange
	args := []string{"--length=20"}

	// Act
	config, err := parseFlags(args)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := 20
	if config.Length != expected {
		t.Errorf("expected length %d, got %d", expected, config.Length)
	}
}

func TestParseFlags_SetsDigitsToTrue(t *testing.T) {
	// Arrange
	args := []string{"--digits"}

	// Act
	config, err := parseFlags(args)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !config.Digits {
		t.Error("expected digits to be true")
	}
}

func TestParseFlags_SetsSymbolsToTrue(t *testing.T) {
	// Arrange
	args := []string{"--symbols"}

	// Act
	config, err := parseFlags(args)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !config.Symbols {
		t.Error("expected symbols to be true")
	}
}

func TestParseFlags_SetsUppercaseToTrue(t *testing.T) {
	// Arrange
	args := []string{"--uppercase"}

	// Act
	config, err := parseFlags(args)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !config.Uppercase {
		t.Error("expected uppercase to be true")
	}
}
