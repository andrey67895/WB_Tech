package sortlogic

import "strings"

// monthMap хранит сокращения месяцев
var monthMap = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

// parseMonth возвращает номер месяца по строке
func parseMonth(s string) (int, bool) {
	if len(s) < 3 {
		return 0, false
	}
	val, ok := monthMap[strings.ToLower(s[:3])]
	return val, ok
}
