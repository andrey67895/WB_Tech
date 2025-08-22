package sortlogic

import (
	"strconv"
	"strings"
)

// makeComparator строит функцию сравнения строк согласно заданным опциям.
func makeComparator(opts Options) func(a, b string) bool {
	return func(a, b string) bool {
		ka := extractKey(a, opts)
		kb := extractKey(b, opts)

		var less bool
		switch {
		case opts.Month:
			ma, ok1 := parseMonth(ka)
			mb, ok2 := parseMonth(kb)
			if ok1 && ok2 {
				less = ma < mb
			} else {
				less = ka < kb
			}
		case opts.Human:
			va, oka := parseHuman(ka)
			vb, okb := parseHuman(kb)
			if oka && okb {
				less = va < vb
			} else {
				less = ka < kb
			}
		case opts.Numeric:
			na, err1 := strconv.ParseFloat(ka, 64)
			nb, err2 := strconv.ParseFloat(kb, 64)
			if err1 == nil && err2 == nil {
				less = na < nb
			} else {
				less = ka < kb
			}
		default:
			less = ka < kb
		}

		if opts.Reverse {
			return !less
		}
		return less
	}
}

// extractKey извлекает ключ сортировки из строки.
func extractKey(s string, opts Options) string {
	if opts.TrimTail {
		s = strings.TrimRight(s, " \t")
	}
	if opts.Column > 0 {
		cols := strings.Split(s, "\t")
		if opts.Column-1 < len(cols) {
			return cols[opts.Column-1]
		}
		return ""
	}
	return s
}

// uniqueAdjacent удаляет соседние дубликаты.
func uniqueAdjacent(lines []string, equal func(a, b string) bool) []string {
	if len(lines) == 0 {
		return lines
	}
	out := []string{lines[0]}
	for i := 1; i < len(lines); i++ {
		if !equal(lines[i], lines[i-1]) {
			out = append(out, lines[i])
		}
	}
	return out
}
