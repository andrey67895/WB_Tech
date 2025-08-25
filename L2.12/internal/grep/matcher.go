package grep

import (
	"regexp"
	"strings"
)

func NewMatcher(pattern string, ignoreCase, fixed, invert bool) (*Matcher, error) {
	m := &Matcher{
		Pattern:    pattern,
		IgnoreCase: ignoreCase,
		Fixed:      fixed,
		Invert:     invert,
	}
	if !fixed {
		flags := ""
		if ignoreCase {
			flags = "(?i)"
		}
		re, err := regexp.Compile(flags + pattern)
		if err != nil {
			return nil, err
		}
		m.re = re
	}
	if fixed && ignoreCase {
		m.Pattern = strings.ToLower(m.Pattern)
	}
	return m, nil
}

func (m *Matcher) Match(line string) bool {
	var matched bool
	if m.Fixed {
		cmp := line
		if m.IgnoreCase {
			cmp = strings.ToLower(cmp)
		}
		matched = strings.Contains(cmp, m.Pattern)
	} else {
		matched = m.re.MatchString(line)
	}
	if m.Invert {
		return !matched
	}
	return matched
}
