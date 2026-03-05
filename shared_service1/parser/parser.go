package parser

import (
	"sort"
	"strings"
	"unicode"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(input string) *Parser {
	return NewParserWithLexer(NewLexer(input))
}

func NewParserWithLexer(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() []Token {

	l := p.lexer

	var tokens []Token

	for {
		start := l.pos

		symbol, ok := l.Next()

		if !ok {
			tokens = append(tokens, Token{Type: EOF, Value: "EOF", Start: start, End: l.pos})
			break
		}

		var token Token

		switch {
		case unicode.IsSpace(symbol):
			token = Token{Type: SPACE, Value: string(symbol), Start: start, End: l.pos}
		case unicode.IsLetter(symbol):

			for {
				symbol, _ := l.Peek()
				if unicode.IsSpace(symbol) || !unicode.IsLetter(symbol) {
					break
				}
				l.NextPos()
			}

			token = Token{Type: IDENT, Value: string(p.lexer.input[start:l.pos]), Start: start, End: l.pos}
			tokens = append(tokens, token)
			continue
		case unicode.IsDigit(symbol):

			for {
				symbol, _ := l.Peek()
				if unicode.IsSpace(symbol) || !unicode.IsDigit(symbol) {
					break
				}
				l.NextPos()
			}

			token = Token{Type: NUMBER, Value: string(p.lexer.input[start:l.pos]), Start: start, End: l.pos}
			tokens = append(tokens, token)
			continue
		case symbol == '{':
			token = Token{Type: LBRACE, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '}':
			token = Token{Type: RBRACE, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '(':
			token = Token{Type: LPAREN, Value: string(symbol), Start: start, End: l.pos}
		case symbol == ')':
			token = Token{Type: RPAREN, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '=':
			token = Token{Type: ASSIGN, Value: string(symbol), Start: start, End: l.pos}
		case symbol == ':':
			token = Token{Type: COLON, Value: string(symbol), Start: start, End: l.pos}
		case symbol == ',':
			token = Token{Type: COMMA, Value: string(symbol), Start: start, End: l.pos}
		case symbol == ';':
			token = Token{Type: SEMICOLON, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '|':
			token = Token{Type: PIPE, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '"':
			token = Token{Type: QUOTE, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '\'':
			token = Token{Type: CHAR, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '-':
			token = Token{Type: MINUS, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '+':
			token = Token{Type: PLUS, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '*':
			token = Token{Type: ASTERISK, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '/':
			token = Token{Type: SLASH, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '%':
			token = Token{Type: PERCENT, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '^':
			token = Token{Type: CARET, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '&':
			token = Token{Type: AMPERSAND, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '|':
			token = Token{Type: BAR, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '_':
			token = Token{Type: UNDERSCORE, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '@':
			token = Token{Type: AT, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '!':
			token = Token{Type: EXCLAMATION, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '.':
			token = Token{Type: PERIOD, Value: string(symbol), Start: start, End: l.pos}
		case symbol == '?':
			token = Token{Type: QUESTION, Value: string(symbol), Start: start, End: l.pos}
		default:
			token = Token{Type: INVALID, Value: string(symbol), Start: start, End: l.pos}
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func (p *Parser) ParseText() Token {
	p.lexer.EscapeSpace()

	isLast := false
	for {
		ch, ok := p.lexer.Peek()
		{

			if !ok {
				isLast = true
				break
			}

			if !unicode.IsSymbol(ch) {
				break
			}

			p.lexer.NextPos()
		}
	}

	start := p.lexer.pos
	for {
		ch, ok := p.lexer.Peek()
		{

			if !ok {
				isLast = true
				break
			}

			if unicode.IsSymbol(ch) || unicode.IsSpace(ch) {
				break
			}

			p.lexer.NextPos()
		}
	}

	return Token{Type: TEXT, Value: string(p.lexer.input[start:p.lexer.pos]), Start: start, End: p.lexer.pos, IsLast: isLast}
}

func (p *Parser) ParseTexts() []Token {
	var tokens []Token

	for {
		token := p.ParseText()

		tokens = append(tokens, token)

		if token.IsLast {
			break
		}
	}

	return tokens
}

func (p *Parser) ParsePlaceholder() (Token, bool) {

	for {
		token, ok := p.lexer.Peek()
		{
			if !ok {
				return Token{Type: EOF, Value: "PLACEHOLDER not found"}, false
			}
		}

		if token == '{' {
			break
		}

		p.lexer.NextPos()
	}

	token, ok := p.lexer.Read('{', '}', IDENT)
	{
		if !ok {
			return Token{Type: EOF, Value: "PLACEHOLDER not found"}, false
		}
	}

	return token, true
}

func (p *Parser) ParsePlaceholders() []Token {

	var tokens []Token

	for {
		token, ok := p.ParsePlaceholder()
		if !ok {
			break
		}
		tokens = append(tokens, token)
	}

	return tokens
}

func Filter[T any](input []T, fn func(T) bool) []T {
	var output []T
	for _, v := range input {
		if fn(v) {
			output = append(output, v)
		}
	}
	return output
}

type Replacer interface {
	map[string]string | func(string) string
}

func ReplaceWithTokens[T Replacer](input string, tokens []Token, replacer T) string {

	tokens = Filter(tokens, func(token Token) bool {
		return token.Type == IDENT
	})

	if len(tokens) == 0 {
		return input
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].Start < tokens[j].Start
	})

	var repl func(string) string
	{
		switch v := any(replacer).(type) {
		case map[string]string:
			repl = func(key string) string {
				return v[key]
			}
		case func(string) string:
			repl = v
		}
	}

	var prevEnd int

	buff := make([]byte, 0, len(input))

	for _, token := range tokens {

		segment := input[prevEnd:token.Start]

		value := token.Trim()
		{
			if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
				value = value[1 : len(value)-1]
			}
		}

		combined := segment + repl(value)

		buff = append(buff, combined...)

		prevEnd = token.End
	}

	buff = append(buff, input[prevEnd:]...)

	return string(buff)
}