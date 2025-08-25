package grep

import "regexp"

// Options хранит настройки поиска (аналог флагов).
type Options struct {
	Before  int
	After   int
	Count   bool
	LineNum bool
}

// Result хранит результат поиска по строке.
type Result struct {
	Line    string
	LineNum int
	Matched bool
}

type Matcher struct {
	Pattern    string
	IgnoreCase bool
	Fixed      bool
	Invert     bool
	re         *regexp.Regexp
}
