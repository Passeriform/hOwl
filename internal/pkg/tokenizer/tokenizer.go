package tokenizer

import (
	"strings"
)

type TokenType struct {
	Token Token
	Literal string
}

func Tokenize(raw_str string) []TokenType {
	scanner := NewScanner(strings.NewReader(raw_str))
	var buf []TokenType

	for {
		token, literal := scanner.Scan()
		buf = append(buf, TokenType{Token: token, Literal: literal})

		if token == EOF {
			break
		}
	}

	return buf
}
