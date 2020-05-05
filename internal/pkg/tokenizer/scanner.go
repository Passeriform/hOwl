package tokenizer

import (
	"bytes"
	"io"

	buffruneio "github.com/pelletier/go-buffruneio"
)

// Scanner represents a non-discriminatory reader.
type Scanner struct {
	reader *buffruneio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{reader: buffruneio.NewReader(reader)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (scanner *Scanner) read() rune {
	ch, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (scanner *Scanner) peek() rune {
	ch := scanner.reader.PeekRunes(1)[0]

	return ch
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (scanner *Scanner) peekMany(p_ind int) []rune {
	ch_slice := scanner.reader.PeekRunes(p_ind)

	return ch_slice
}

// unread places the previously read rune back on the reader.
func (scanner *Scanner) unread() { _ = scanner.reader.UnreadRune() }

// Scan returns the next token and literal value.
func (scanner *Scanner) Scan() (tok Token, lit string) {
	if isEOF(scanner) {
		return EOF, string(eof)
	} else if isWhitespace(scanner) {
		return scanner.scanWhitespace()
	} else if isMatches(scanner) {
		return scanner.scanMatches()
	} else if isSymbol(scanner) {
		return scanner.scanSymbol()
	} else if isContextual(scanner) {
		return scanner.scanContext()
	} else {
		return scanner.scanIdent()
	}
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (scanner *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	if isNewline(scanner) {
		for {
			if !isNewline(scanner) {
				break
			}
			buf.WriteRune(scanner.read())
		}
		return NL, buf.String()
	} else if isTab(scanner) {
		buf.WriteRune(scanner.read())
		return TAB, buf.String()
	}
	buf.WriteRune(scanner.read())
	return ILLEGAL, buf.String()
}

func (scanner *Scanner) scanMatches() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	if isParentheses(scanner) {
		ch := scanner.read()
		buf.WriteRune(ch)

		if ch == '(' {
			return PAREN_OPEN, buf.String()
		} else if ch == ')' {
			return PAREN_CLOSE, buf.String()
		}
	} else if isBraces(scanner) {
		ch := scanner.read()
		buf.WriteRune(ch)

		if ch == '{' {
			return BRACE_OPEN, buf.String()
		} else if ch == '}' {
			return BRACE_CLOSE, buf.String()
		}
	} else if isBrackets(scanner) {
		ch := scanner.read()
		buf.WriteRune(ch)

		if ch == '[' {
			return BRACKET_OPEN, buf.String()
		} else if ch == ']' {
			return BRACKET_CLOSE, buf.String()
		}
	}
	buf.WriteRune(scanner.read())
	return ILLEGAL, buf.String()
}

// scanSymbol consumes the current rune and all contiguous symbol.
func (scanner *Scanner) scanSymbol() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// Comparison precedence is crucial. Must be coded like so
	if isArrow(scanner) {
		tail := scanner.read() // Arrow tail
		head := scanner.read() // Arrow head

		buf.WriteRune(tail)
		buf.WriteRune(head)

		if tail == '=' {
			return FAT_ARROW, buf.String()
		} else if tail == '-' {
			return SLIM_ARROW, buf.String()
		}
	} else if isComparison(scanner) {
		f_eq := scanner.read()
		buf.WriteRune(f_eq)

		if f_eq == '<' {
			if isComparison(scanner) {
				s_eq := scanner.read()
				buf.WriteRune(s_eq)

				if s_eq == '=' {
					return LTE, buf.String()
				}
				if s_eq == '>' {
					return UNEQ, buf.String()
				}
				return ILLEGAL, buf.String()
			}
			return LT, buf.String()
		} else if f_eq == '>' {
			if isComparison(scanner) {
				s_eq := scanner.read()
				buf.WriteRune(s_eq)

				if s_eq == '=' {
					return GTE, buf.String()
				}
				return ILLEGAL, buf.String()
			}
			return GT, buf.String()
		} else if f_eq == '=' {
			if isComparison(scanner) {
				s_eq := scanner.read()
				buf.WriteRune(s_eq)

				if s_eq == '=' {
					return EQEQ, buf.String()
				}
				return ILLEGAL, buf.String()
			}
			return EQ, buf.String()
		}
	} else if isHyphen(scanner) {
		buf.WriteRune(scanner.read())
		return HYPHEN, buf.String()
	} else if isArithmetic(scanner) {
		ch := scanner.read()
		buf.WriteRune(ch)

		if ch == '+' {
			return PLUS, buf.String()
		} else if ch == '-' {
			return MINUS, buf.String()
		} else if ch == '*' {
			return MUL, buf.String()
		} else if ch == '/' {
			return DIV, buf.String()
		}
	}
	return ILLEGAL, buf.String()
}

func (scanner *Scanner) scanContext() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	ch := scanner.read()
	buf.WriteRune(ch)

	if ch == '.' {
		return PERIOD_MARK, buf.String()
	} else if ch == '?' {
		return QUESTION_MARK, buf.String()
	} else if ch == '!' {
		return EXCLAIM_MARK, buf.String()
	}
	return ILLEGAL, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (scanner *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if isReserved(scanner) {
			break
		}
		ch := scanner.read()
		buf.WriteRune(ch)
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}
