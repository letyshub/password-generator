package generator

import (
	"password-generator/config"
	"strings"
	"testing"
	"unicode"
)

func TestGeneratePassword_WithOnlyLowercase(t *testing.T) {
	// Arrange
	cfg := config.Config{
		Length:    12,
		Digits:    false,
		Symbols:   false,
		Uppercase: false,
	}

	// Act
	password, err := GeneratePassword(cfg)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(password) != cfg.Length {
		t.Errorf("expected password length %d, got %d", cfg.Length, len(password))
	}
	for _, ch := range password {
		if !unicode.IsLower(ch) {
			t.Errorf("expected only lowercase letters, found '%c'", ch)
		}
	}
}

func TestGeneratePassword_WithDigits(t *testing.T) {
	// Arrange
	cfg := config.Config{
		Length:    16,
		Digits:    true,
		Symbols:   false,
		Uppercase: false,
	}

	// Act
	password, err := GeneratePassword(cfg)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(password) != cfg.Length {
		t.Errorf("expected length %d, got %d", cfg.Length, len(password))
	}
	foundDigit := false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			foundDigit = true
			break
		}
	}
	if !foundDigit {
		t.Error("expected password to contain at least one digit")
	}
}

func TestGeneratePassword_WithSymbols(t *testing.T) {
	// Arrange
	cfg := config.Config{
		Length:    14,
		Digits:    false,
		Symbols:   true,
		Uppercase: false,
	}

	// Act
	password, err := GeneratePassword(cfg)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	foundSymbol := false
	for _, ch := range password {
		if strings.ContainsRune(symbolChars, ch) {
			foundSymbol = true
			break
		}
	}
	if !foundSymbol {
		t.Error("expected password to contain at least one symbol")
	}
}

func TestGeneratePassword_WithUppercase(t *testing.T) {
	// Arrange
	cfg := config.Config{
		Length:    10,
		Digits:    false,
		Symbols:   false,
		Uppercase: true,
	}

	// Act
	password, err := GeneratePassword(cfg)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	foundUpper := false
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			foundUpper = true
			break
		}
	}
	if !foundUpper {
		t.Error("expected password to contain at least one uppercase character")
	}
}

func TestGeneratePassword_InvalidLength_ReturnsError(t *testing.T) {
	// Arrange
	cfg := config.Config{
		Length:    0,
		Digits:    true,
		Symbols:   true,
		Uppercase: true,
	}

	// Act
	password, err := GeneratePassword(cfg)

	// Assert
	if err == nil {
		t.Fatal("expected error for zero length, got nil")
	}
	if password != "" {
		t.Errorf("expected empty password, got: %s", password)
	}
}
