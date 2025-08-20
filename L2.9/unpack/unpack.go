package unpack

import (
	"errors"
	"unicode"
)

// Unpack выполняет распаковку строки с поддержкой экранирования.
func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	runes := []rune(input)
	var result []rune
	var prev rune
	escaped := false
	hasLetter := false

	for i, r := range runes {
		if escaped {
			appendRune(&result, r)
			prev = r
			escaped = false
			hasLetter = true
			continue
		}

		switch {
		case r == '\\':
			escaped = true

		case unicode.IsDigit(r):
			if err := processDigit(r, i, prev, &result); err != nil {
				return "", err
			}
		default:
			appendRune(&result, r)
			prev = r
			hasLetter = true
		}
	}

	if !hasLetter {
		return "", errors.New("некорректная строка: нет символов")
	}
	if escaped {
		return "", errors.New("некорректная строка: висячий escape")
	}

	return string(result), nil
}

// appendRune добавляет руну в результат.
func appendRune(dst *[]rune, r rune) {
	*dst = append(*dst, r)
}

// processDigit обрабатывает случай, когда встречается цифра.
func processDigit(r rune, pos int, prev rune, dst *[]rune) error {
	if pos == 0 || prev == 0 {
		return errors.New("некорректная строка: цифра без символа")
	}

	count := int(r - '0')
	if count == 0 {
		*dst = (*dst)[:len(*dst)-1]
		return nil
	}
	for j := 1; j < count; j++ {
		*dst = append(*dst, prev)
	}
	return nil
}
