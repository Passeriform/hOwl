package tokenizer

func isEOF(scanner *Scanner) bool {
	return scanner.peek() == eof
}

func isNewline(scanner *Scanner) bool {
	return scanner.peek() == '\n'
}

func isTab(scanner *Scanner) bool {
	return scanner.peek() == '\t'
}

func isParentheses(scanner *Scanner) bool {
	ch := scanner.peek()
	return ch == '(' || ch == ')'
}

func isBraces(scanner *Scanner) bool {
	ch := scanner.peek()
	return ch == '{' || ch == '}'
}

func isBrackets(scanner *Scanner) bool {
	ch := scanner.peek()
	return ch == '[' || ch == ']'
}

func isArrow(scanner *Scanner) bool {
	ch_slice := scanner.peekMany(2)

	if ch_slice[0] == '=' || ch_slice[0] == '-' {
		return ch_slice[1] == '>'
	}
	return false
}

func isComparison(scanner *Scanner) bool {
	ch := scanner.peek()
	result := ch == '<' || ch == '>' || ch == '='
	return result
}

func isHyphen(scanner *Scanner) bool {
	ch_slice := scanner.peekMany(2)
	if ch_slice[0] == '-' && ch_slice[1] == '-' {
		return true
	}
	return false
}

func isArithmetic(scanner *Scanner) bool {
	ch := scanner.peek()
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isContextual(scanner *Scanner) bool {
	ch := scanner.peek()
	return ch == '.' || ch == '?' || ch == '!'
}

func isWhitespace(scanner *Scanner) bool {
	return isNewline(scanner) || isTab(scanner)
}

func isMatches(scanner *Scanner) bool {
	return isParentheses(scanner) || isBraces(scanner) || isBrackets(scanner)
}

func isSymbol(scanner *Scanner) bool {
	return isArrow(scanner) || isComparison(scanner) || isHyphen(scanner)
}

func isReserved(scanner *Scanner) bool {
	return isEOF(scanner) || isWhitespace(scanner) || isSymbol(scanner) || isContextual(scanner)
}
