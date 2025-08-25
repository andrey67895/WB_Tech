package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/andrey67895/WB_Tech/L2.13/internal/cut"
)

// OptionsInput — сырой ввод флагов от пользователя.
type OptionsInput struct {
	FieldsSpec string // пример: "1,3-5"
	Delimiter  string // ожидается 1 символ (можно "\t")
	Separated  bool
}

// ParseOptions парсит CLI флаги в структуру cut.Options.
func ParseOptions(in OptionsInput) (cut.Options, error) {
	if in.Delimiter == "" {
		in.Delimiter = "\t"
	}

	delim, err := normalizeDelimiter(in.Delimiter)
	if err != nil {
		return cut.Options{}, err
	}

	selector, err := parseFieldsSpec(in.FieldsSpec)
	if err != nil {
		return cut.Options{}, err
	}

	return cut.Options{
		Delimiter:     delim,
		SeparatedOnly: in.Separated,
		Selector:      selector,
	}, nil
}

func normalizeDelimiter(s string) (rune, error) {
	// Поддерживаем "\t" и обычные одиночные символы
	if s == "\\t" || s == "\t" {
		return '\t', nil
	}
	runes := []rune(s)
	if len(runes) != 1 {
		return 0, fmt.Errorf("разделитель должен быть ровно один символ, получено: %q", s)
	}
	return runes[0], nil
}

// parseFieldsSpec парсит выражение вида "1,3-5,7".
func parseFieldsSpec(spec string) (cut.FieldSelector, error) {
	spec = strings.TrimSpace(spec)
	if spec == "" {
		return cut.FieldSelector{}, errors.New("пустое выражение полей")
	}

	selector := cut.FieldSelector{}
	parts := strings.Split(spec, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			return cut.FieldSelector{}, fmt.Errorf("пустой элемент в выражении: %q", spec)
		}

		if strings.Contains(p, "-") {
			bounds := strings.Split(p, "-")
			if len(bounds) != 2 {
				return cut.FieldSelector{}, fmt.Errorf("неверный диапазон: %q", p)
			}
			start, err1 := parsePositive(bounds[0])
			end, err2 := parsePositive(bounds[1])
			if err1 != nil || err2 != nil {
				return cut.FieldSelector{}, fmt.Errorf("диапазон должен быть вида N-M, найдено: %q", p)
			}
			if end < start {
				return cut.FieldSelector{}, fmt.Errorf("обратный диапазон: %q", p)
			}
			selector.AddRange(start, end)
			continue
		}

		v, err := parsePositive(p)
		if err != nil {
			return cut.FieldSelector{}, fmt.Errorf("неверный номер поля: %q", p)
		}
		selector.Add(v)
	}
	return selector, nil
}

func parsePositive(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil || v <= 0 {
		return 0, fmt.Errorf("ожидалось положительное число, получено: %q", s)
	}
	return v, nil
}
