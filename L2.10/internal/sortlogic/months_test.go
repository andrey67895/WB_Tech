package sortlogic

import "testing"

func TestParseMonth(t *testing.T) {
	tests := []struct {
		in   string
		want int
		ok   bool
	}{
		{"Jan", 1, true},
		{"feb", 2, true},
		{"MAR", 3, true},
		{"December", 12, true},
		{"abc", 0, false},
		{"", 0, false},
	}

	for _, tt := range tests {
		got, ok := parseMonth(tt.in)
		if got != tt.want || ok != tt.ok {
			t.Errorf("parseMonth(%q) = (%d,%v), ожидалось (%d,%v)",
				tt.in, got, ok, tt.want, tt.ok)
		}
	}
}
