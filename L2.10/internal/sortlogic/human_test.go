package sortlogic

import "testing"

func TestParseHuman(t *testing.T) {
	tests := []struct {
		in   string
		want float64
		ok   bool
	}{
		{"1K", 1024, true},
		{"2M", 2 << 20, true},
		{"1.5K", 1536, true},
		{"5G", 5 << 30, true},
		{"", 0, false},
		{"ABC", 0, false},
	}

	for _, tt := range tests {
		got, ok := parseHuman(tt.in)
		if ok != tt.ok {
			t.Errorf("parseHuman(%q) ok=%v, ожидалось %v", tt.in, ok, tt.ok)
		}
		if ok && got != tt.want {
			t.Errorf("parseHuman(%q)=%f, ожидалось %f", tt.in, got, tt.want)
		}
	}
}
