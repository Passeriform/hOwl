package tokenizer

const eof = rune(-1)

// Token represents a lexical token.
type Token int

const (
	// Illegal
	ILLEGAL Token = iota

	// Whitespace
	EOF
	NL
	TAB

	// Matches
	PAREN_OPEN
	PAREN_CLOSE
	BRACE_OPEN
	BRACE_CLOSE
	BRACKET_OPEN
	BRACKET_CLOSE

	// Arrows
	FAT_ARROW
	SLIM_ARROW

	// Comparison
	LT
	GT
	LTE
	GTE
	EQ
	EQEQ
	UNEQ

	// Hyphen
	HYPHEN

	// Arithmetic
	PLUS
	MINUS
	MUL
	DIV

	// Context
	PERIOD_MARK
	QUESTION_MARK
	EXCLAIM_MARK

	// Literals
	IDENT
)
