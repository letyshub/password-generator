package generator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"password-generator/config"
	"password-generator/utils"
	"strings"
)

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars     = "0123456789"
	symbolChars    = "!@#$%^&*()-_=+[]{}<>?/|"
	ambiguousChars = "0O1lI"
)

func GeneratePassword(config config.Config) (string, error) {
	password, err := generatePassword(config)
	if err != nil {
		return "", err
	}

	for !isPasswordValid(password, config) {
		password, err = generatePassword(config)
		if err != nil {
			return "", err
		}
	}

	return password, nil
}

func isPasswordValid(password string, config config.Config) bool {
	if len(password) != config.Length {
		return false
	}

	if config.Digits && !utils.ContainsAny(password, digitChars) {
		return false
	}
	if config.Symbols && !utils.ContainsAny(password, symbolChars) {
		return false
	}
	if config.Uppercase && !utils.ContainsAny(password, uppercaseChars) {
		return false
	}

	return true
}

func generatePassword(config config.Config) (string, error) {
	if config.Length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	charset := lowercaseChars

	if config.Digits {
		charset += digitChars
	}
	if config.Symbols {
		charset += symbolChars
	}
	if config.Uppercase {
		charset += uppercaseChars
	}

	if len(charset) == 0 {
		return "", errors.New("no character sets selected")
	}

	var password strings.Builder

	for i := 0; i < config.Length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password.WriteByte(charset[idx.Int64()])
	}

	return password.String(), nil
}
