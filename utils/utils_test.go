package utils

import "testing"

func TestContainsAny_WhenStringContainsAtLeastOneChar_ReturnsTrue(t *testing.T) {
	// Arrange
	input := "password123"
	chars := "123"

	// Act
	result := ContainsAny(input, chars)

	// Assert
	if result != true {
		t.Error("expected true, got false")
	}
}

func TestContainsAny_WhenStringDoesNotContainAnyChars_ReturnsFalse(t *testing.T) {
	// Arrange
	input := "password"
	chars := "XYZ"

	// Act
	result := ContainsAny(input, chars)

	// Assert
	if result != false {
		t.Error("expected false, got true")
	}
}

func TestContainsAny_WithEmptySearchChars_ReturnsFalse(t *testing.T) {
	// Arrange
	input := "password"
	chars := ""

	// Act
	result := ContainsAny(input, chars)

	// Assert
	if result != false {
		t.Error("expected false for empty chars string")
	}
}

func TestContainsAny_WithEmptyInputString_ReturnsFalse(t *testing.T) {
	// Arrange
	input := ""
	chars := "abc"

	// Act
	result := ContainsAny(input, chars)

	// Assert
	if result != false {
		t.Error("expected false for empty input string")
	}
}

func TestContainsAny_WithExactMatch_ReturnsTrue(t *testing.T) {
	// Arrange
	input := "a"
	chars := "a"

	// Act
	result := ContainsAny(input, chars)

	// Assert
	if result != true {
		t.Error("expected true for exact match")
	}
}
