package sortlogic

import (
	"strconv"
	"strings"
)

// множители для человекочитаемых чисел
var humanMultipliers = map[string]float64{
	"":  1,
	"B": 1,
	"K": 1 << 10,
	"M": 1 << 20,
	"G": 1 << 30,
	"T": 1 << 40,
	"P": 1 << 50,
	"E": 1 << 60,
}

// parseHuman парсит строку с суффиксом (например 1K, 10M)
func parseHuman(s string) (float64, bool) {
	if s == "" {
		return 0, false
	}
	base := s
	suffix := ""
	if last := s[len(s)-1]; (last >= 'A' && last <= 'Z') || (last >= 'a' && last <= 'z') {
		base = s[:len(s)-1]
		suffix = strings.ToUpper(string(last))
	}

	val, err := strconv.ParseFloat(base, 64)
	if err != nil {
		return 0, false
	}

	mult, ok := humanMultipliers[suffix]
	if !ok {
		return 0, false
	}
	return val * mult, true
}
