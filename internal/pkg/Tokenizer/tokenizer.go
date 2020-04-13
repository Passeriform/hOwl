package tokenizer

import (
	"strings"
)

type TokenBuffer struct {
	tok Token
	lit string
}

func Tokenize(raw_str string) []TokenBuffer {
	scanner := NewScanner(strings.NewReader(raw_str))
	var buf []TokenBuffer

	for {
		tok, lit := scanner.Scan()
		buf = append(buf, TokenBuffer{tok: tok, lit: lit})

		if tok == EOF {
			break
		}
	}

	return buf
}
