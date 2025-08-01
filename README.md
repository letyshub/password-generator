# password-generator

A simple and secure command-line tool written in Go for generating random passwords. Customizable, fast, and easy to use.

# Command-Line Flags

The password generator supports the following command-line flags to customize password output:

| Flag          | Type   | Default | Description                                                                |
| ------------- | ------ | ------- | -------------------------------------------------------------------------- |
| `--length`    | `int`  | `12`    | Length of the password to generate.                                        |
| `--digits`    | `bool` | `false` | Include numeric characters (`0–9`) in the password.                        |
| `--symbols`   | `bool` | `false` | Include special characters (e.g. `!@#$%^&*()_+-=[]{}<>?`) in the password. |
| `--uppercase` | `bool` | `false` | Include uppercase letters (`A–Z`) in the password.                         |

Example usage:

```
password-generator.exe --length 16 --digits --symbols --uppercase
```
