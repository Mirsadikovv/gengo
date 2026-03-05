package parser

import "unicode"

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input), pos: 0}
}

func (l *Lexer) Next() (rune, bool) {
	if l.pos >= len(l.input) {
		return 0, false
	}
	r := l.input[l.pos]
	l.pos++
	return r, true
}

func (l *Lexer) Peek() (rune, bool) {
	if l.pos >= len(l.input) {
		return 0, false
	}
	return l.input[l.pos], true
}

func (l *Lexer) NextPos() {
	l.pos++
}

func (l *Lexer) NextPosN(n int) {
	l.pos += n
}

func (l *Lexer) PrevPos() {
	l.pos--
}

func (l *Lexer) PrevPosN(n int) {
	l.pos -= n
}

func (l *Lexer) EscapeSpace() {
	for {
		r, ok := l.Peek()
		{
			if !ok {
				break
			}
		}

		if !unicode.IsSpace(r) {
			break
		}

		l.NextPos()
	}
}

func (l *Lexer) Read(open, close rune, tokenType TokenType) (Token, bool) {

	l.EscapeSpace()

	var start int = l.pos
	{
		if l.pos > 0 {
			start = l.pos
		}
	}

	for {
		r, ok := l.Next()
		{
			if !ok {
				return Token{Type: EOF, Start: start, End: l.pos}, false
			}
		}

		if r == close {
			break
		}
	}

	end := l.pos

	token := Token{
		Type:  tokenType,
		Value: string(l.input[start:end]),
		Start: start,
		End:   end,
	}

	return token, true
}