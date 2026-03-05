package parser

import (
	"strings"
	"unicode"
	"unsafe"
)

func CamelCaseOptimized(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	capitalizeNext := false
	for i, r := range input {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			capitalizeNext = true
			continue
		}

		if capitalizeNext || i == 0 {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	res := []rune(result.String())
	res[0] = unicode.ToLower(res[0])
	return string(res)
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func CamelCaseUnsafe(input string) string {

	inBytes := *(*[]byte)(unsafe.Pointer(&input))

	outBytes := unsafe.Slice(unsafe.StringData(input), 0)[:]

	capitalizeNext := true
	for _, b := range inBytes {
		if b == '_' || b == '-' || b == ' ' || unicode.IsPunct(rune(b)) {
			capitalizeNext = true
			continue
		}
		if capitalizeNext {
			if b >= 'a' && b <= 'z' {
				b -= 'a' - 'A'
			}
			capitalizeNext = false
		} else {
			if b >= 'A' && b <= 'Z' {
				b += 'a' - 'A'
			}
		}
		outBytes = append(outBytes, b)
	}

	return *(*string)(unsafe.Pointer(&outBytes))
	//return ""
}

func CamelCaseUnsafe2(input string) string {
	// Получаем указатель на исходные байты строки
	inBytes := unsafe.Slice(unsafe.StringData(input), len(input))

	// Создаем новый срез для результата
	outBytes := make([]byte, 0, len(input))

	capitalizeNext := true
	for _, b := range inBytes {
		if b == '_' || b == '-' || b == ' ' || unicode.IsPunct(rune(b)) {
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			// Преобразуем в заглавный символ, если это буква
			if b >= 'a' && b <= 'z' {
				b -= 'a' - 'A'
			}
			capitalizeNext = false
		} else {
			// Преобразуем в строчный символ, если это буква
			if b >= 'A' && b <= 'Z' {
				b += 'a' - 'A'
			}
		}

		// Добавляем байт в результат
		outBytes = append(outBytes, b)
	}

	// Преобразуем результат обратно в строку через unsafe
	return *(*string)(unsafe.Pointer(&outBytes))
}

func CamelCaseUnsafe3(input string) string {
	// Получаем указатель на исходные байты строки
	inBytes := unsafe.Slice(unsafe.StringData(input), len(input))

	// Выделяем память под результат с использованием unsafe
	outBytes := make([]byte, len(input))
	outPtr := unsafe.Pointer(&outBytes[0])

	// Инициализируем счетчик для записи в новый срез
	writeIndex := 0
	capitalizeNext := true

	for _, b := range inBytes {
		if b == '_' || b == '-' || b == ' ' || unicode.IsPunct(rune(b)) {
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			// Преобразуем в заглавный символ, если это буква
			if b >= 'a' && b <= 'z' {
				b -= 'a' - 'A'
			}
			capitalizeNext = false
		} else {
			// Преобразуем в строчный символ, если это буква
			if b >= 'A' && b <= 'Z' {
				b += 'a' - 'A'
			}
		}

		// Записываем байт в выделенную память
		*(*byte)(unsafe.Add(outPtr, writeIndex)) = b
		writeIndex++
	}

	// Преобразуем результат обратно в строку
	return *(*string)(unsafe.Pointer(&outBytes))
}

func CamelCaseUnsafeTokens(input string) []Token {
	// Получаем указатель на исходные байты строки
	// inBytes := unsafe.Slice(unsafe.SliceData(input), len(input))

	// Создаем новый срез для результата
	// outBytes := make([]byte, 0, len(input))

	inRunes := []rune(input)

	var result []Token

	capitalizeNext := true
	for _, b := range inRunes {

		if b == '_' || b == '-' || b == ' ' || unicode.IsPunct(rune(b)) {
			capitalizeNext = true
			continue
		}

		// if capitalizeNext {
		// 	// Преобразуем в заглавный символ, если это буква
		// 	if b >= 'a' && b <= 'z' {
		// 		b -= 'a' - 'A'
		// 	}
		// 	capitalizeNext = false
		// } else {
		// 	// Преобразуем в строчный символ, если это буква
		// 	if b >= 'A' && b <= 'Z' {
		// 		b += 'a' - 'A'
		// 	}
		// }

		if capitalizeNext {
			if b >= 'a' && b <= 'z' {
				// b -= 'a' - 'A'
			}
			capitalizeNext = false

		} else {
			if b >= 'A' && b <= 'Z' {
				// b += 'a' - 'A'
			}
		}

		// result = append(result, Token{
		// 	Type:  IDENT,
		// 	Value: string(b),
		// 	Start: 0,
		// 	End:   0,
		// })
	}

	// Преобразуем результат обратно в строку через unsafe
	return result
}

func CamelCase(s string) string {
	p := NewParser(s)

	var result = make([]byte, 0, len(s))

	// tokens := Filter(p.Parse(), func(t Token) bool {
	// 	return t.Type == IDENT
	// })

	for i, token := range p.Parse() {
		if token.Type == IDENT {
			res := token.Join(token.SplitUpper, func(s string) string {
				if len(s) == 0 {
					return s
				}

				if i == 0 {
					q := []rune(s)
					if unicode.IsLower(q[0]) {
						return s
					}
					q[0] = unicode.ToLower(q[0])
					return string(q)
				}

				q := []rune(s)

				if unicode.IsUpper(q[0]) {
					return s
				}

				q[0] = unicode.ToUpper(q[0])
				return string(q)
			})

			result = append(result, res...)
		}
	}

	return string(result)
}

// func PascalCase(s string) string {
// 	p := NewParser(s)

// 	tokens := p.Parse()

// 	for _, token := range tokens {
// 		token.UpperFirst()
// 	}

// 	return strings.Join(tokens, "")
// }

func SnakeCase(s string) string {
	p := NewParser(s)

	tokens := Filter(p.Parse(), func(t Token) bool {
		return t.Type == IDENT
	})

	var result []byte

	for i, token := range tokens {
		if token.Type == IDENT {

			res := token.Join(token.SplitUpper, func(s string) string {
				if i == 0 {
					return strings.ToLower(s)
				}
				return "_" + strings.ToLower(s)
			})

			result = append(result, res...)
		}
	}

	return string(result)
}