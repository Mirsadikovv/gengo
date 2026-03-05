package parser

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	EOF         TokenType = iota // End of file
	SPACE                        // Space
	INVALID                      // Invalid
	STRING                       // String
	NUMBER                       // Number
	BOOL                         // Boolean
	NULL                         // Null
	COMMENT                      // Comment
	TEXT                         // Text
	LPAREN                       // Left parenthesis (
	RPAREN                       // Right parenthesis )
	LBRACE                       // Left brace {
	RBRACE                       // Right brace }
	IDENT                        // Identifier
	ASSIGN                       // Assignment operator =
	COLON                        // Colon :
	COMMA                        // Comma ,
	SEMICOLON                    // Semicolon ;
	PIPE                         // Pipe |
	QUOTE                        // Quote "
	CHAR                         // '
	MINUS                        // -
	PLUS                         // +
	ASTERISK                     // *
	SLASH                        // /
	PERCENT                      // %
	CARET                        // ^
	AMPERSAND                    // &
	BAR                          // |
	UNDERSCORE                   // _
	AT                           // @
	EXCLAMATION                  // !
	QUESTION                     // ?
	PERIOD                       // .
)

var TokenStrings = [...]string{
	EOF:         "EOF",         // End of file
	SPACE:       "SPACE",       // Space
	INVALID:     "INVALID",     // Invalid
	STRING:      "STRING",      // String
	NUMBER:      "NUMBER",      // Number
	BOOL:        "BOOL",        // Boolean
	NULL:        "NULL",        // Null
	COMMENT:     "COMMENT",     // Comment
	TEXT:        "TEXT",        // Text
	LPAREN:      "LPAREN",      // Left parenthesis (
	RPAREN:      "RPAREN",      // Right parenthesis )
	LBRACE:      "LBRACE",      // Left brace {
	RBRACE:      "RBRACE",      // Right brace }
	IDENT:       "IDENT",       // Identifier
	ASSIGN:      "ASSIGN",      // Assignment operator =
	COLON:       "COLON",       // Colon :
	COMMA:       "COMMA",       // Comma ,
	SEMICOLON:   "SEMICOLON",   // Semicolon ;
	PIPE:        "PIPE",        // Pipe |
	QUOTE:       "QUOTE",       // Quote "
	CHAR:        "CHAR",        // '
	MINUS:       "MINUS",       // -
	PLUS:        "PLUS",        // +
	ASTERISK:    "ASTERISK",    // *
	SLASH:       "SLASH",       // /
	PERCENT:     "PERCENT",     // %
	CARET:       "CARET",       // ^
	AMPERSAND:   "AMPERSAND",   // &
	BAR:         "BAR",         // |
	UNDERSCORE:  "UNDERSCORE",  // _
	AT:          "AT",          // @
	EXCLAMATION: "EXCLAMATION", // !
	QUESTION:    "QUESTION",    // ?
	PERIOD:      "PERIOD",      // .
}

var tokenTypeNames = map[string]TokenType{
	"EOF":         EOF,         // End of file
	"SPACE":       SPACE,       // Space
	"INVALID":     INVALID,     // Invalid
	"STRING":      STRING,      // String
	"NUMBER":      NUMBER,      // Number
	"BOOL":        BOOL,        // Boolean
	"NULL":        NULL,        // Null
	"COMMENT":     COMMENT,     // Comment
	"TEXT":        TEXT,        // Text
	"LPAREN":      LPAREN,      // Left parenthesis (
	"RPAREN":      RPAREN,      // Right parenthesis )
	"LBRACE":      LBRACE,      // Left brace {
	"RBRACE":      RBRACE,      // Right brace }
	"IDENT":       IDENT,       // Identifier
	"ASSIGN":      ASSIGN,      // Assignment operator =
	"COLON":       COLON,       // Colon :
	"COMMA":       COMMA,       // Comma ,
	"SEMICOLON":   SEMICOLON,   // Semicolon ;
	"PIPE":        PIPE,        // Pipe |
	"QUOTE":       QUOTE,       // Quote "
	"CHAR":        CHAR,        // '
	"MINUS":       MINUS,       // -
	"PLUS":        PLUS,        // +
	"ASTERISK":    ASTERISK,    // *
	"SLASH":       SLASH,       // /
	"PERCENT":     PERCENT,     // %
	"CARET":       CARET,       // ^
	"AMPERSAND":   AMPERSAND,   // &
	"BAR":         BAR,         // |
	"UNDERSCORE":  UNDERSCORE,  // _
	"AT":          AT,          // @
	"EXCLAMATION": EXCLAMATION, // !
	"QUESTION":    QUESTION,    // ?
	"PERIOD":      PERIOD,      // .
}

type Token struct {
	Type   TokenType
	Value  string
	Start  int
	End    int
	IsLast bool
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %s [%d:%d]\n", TokenStrings[t.Type], t.Value, t.Start, t.End)
}

func (t Token) Trim() string {
	return strings.TrimSpace(t.Value)
}

func (t Token) IsUpper() bool {
	if len(t.Value) == 0 {
		return false
	}
	return unicode.IsUpper(rune(t.Value[0]))
}

func (t Token) IsLower() bool {
	if len(t.Value) == 0 {
		return false
	}
	return unicode.IsLower(rune(t.Value[0]))
}

func (t Token) IsDigit() bool {
	if len(t.Value) == 0 {
		return false
	}
	return unicode.IsDigit(rune(t.Value[0]))
}

func (t Token) IsLetter() bool {
	if len(t.Value) == 0 {
		return false
	}
	return unicode.IsLetter(rune(t.Value[0]))
}

func (t Token) Split(splitter func(rune) bool) []Token {

	var tokens []Token
	start := 0
	runes := []rune(t.Value)

	for i := 0; i < len(runes); i++ {
		if i > 0 && splitter(runes[i]) {
			tokens = append(tokens, Token{
				Type:  t.Type,
				Value: string(runes[start:i]),
				Start: start,
				End:   i,
			})
			start = i
		}
	}

	if start < len(runes) {
		tokens = append(tokens, Token{
			Type:  t.Type,
			Value: string(runes[start:]),
			Start: start,
			End:   len(runes),
		})
	}

	return tokens
}

func (t Token) SplitUpper() []Token {
	return t.Split(unicode.IsUpper)
}

func (t Token) SplitLower() []Token {
	return t.Split(unicode.IsLower)
}

func (t Token) SplitDigit() []Token {
	return t.Split(unicode.IsDigit)
}

func (t Token) SplitLetter() []Token {
	return t.Split(unicode.IsLetter)
}

func (t Token) SplitSpace() []Token {
	return t.Split(unicode.IsSpace)
}

func (t Token) SplitPunct() []Token {
	return t.Split(unicode.IsPunct)
}

func (t Token) Join(joiner func() []Token, transform func(string) string) string {
	var result []byte
	for _, v := range joiner() {
		result = append(result, transform(v.Value)...)
	}
	return string(result)
}

func (t Token) JoinUpper() string {
	return t.Join(t.SplitUpper, func(s string) string {
		return strings.ToUpper(s)
	})
}

func (t Token) JoinLower() string {
	return t.Join(t.SplitLower, func(s string) string {
		return strings.ToLower(s)
	})
}

func (t Token) JoinDigit() string {
	return t.Join(t.SplitDigit, func(s string) string {
		return s
	})
}

func (t Token) JoinLetter() string {
	return t.Join(t.SplitLetter, func(s string) string {
		return s
	})
}

func (t Token) UpperFirst() string {
	return t.Join(t.SplitUpper, func(s string) string {
		return strings.ToUpper(s[0:1]) + s[1:]
	})
}

func (t Token) LowerFirst() string {
	return t.Join(t.SplitLetter, func(s string) string {
		return strings.ToLower(s[0:1]) + s[1:]
	})
}
