package main

import (
	"flag"
	"fmt"
	"os"
	"password-generator/config"
)

func parseFlags(args []string) (config.Config, error) {
	fs := flag.NewFlagSet("password-generator", flag.ContinueOnError)

	length := fs.Int("length", 12, "Length of the password")
	digits := fs.Bool("digits", false, "Include digits (0-9)")
	symbols := fs.Bool("symbols", false, "Include symbols (e.g. !@#$%)")
	uppercase := fs.Bool("uppercase", false, "Include uppercase letters (A-Z)")

	err := fs.Parse(args)
	if err != nil {
		return config.Config{}, err
	}

	return config.Config{
		Length:    *length,
		Digits:    *digits,
		Symbols:   *symbols,
		Uppercase: *uppercase,
	}, nil
}

func main() {
	config, err := parseFlags(os.Args[1:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}
}
