package unpack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		in   string
		want string
		err  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"", "", false},
		{"45", "", true},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"aaa0b", "aab", false},
	}

	for _, tt := range tests {
		got, err := Unpack(tt.in)
		if tt.err {
			require.Error(t, err, tt.in)
		} else {
			require.NoError(t, err, tt.in)
			require.Equal(t, tt.want, got, tt.in)
		}
	}
}

func TestAppendRune(t *testing.T) {
	var res []rune
	appendRune(&res, 'a')
	appendRune(&res, 'б')
	require.Equal(t, []rune{'a', 'б'}, res)
}

func TestProcessDigit(t *testing.T) {
	tests := []struct {
		name    string
		inRune  rune
		pos     int
		prev    rune
		start   []rune
		want    []rune
		wantErr bool
	}{
		{"repeat 3", '3', 1, 'a', []rune{'a'}, []rune{'a', 'a', 'a'}, false},
		{"zero remove", '0', 1, 'a', []rune{'a'}, []rune{}, false},
		{"digit first", '2', 0, 0, []rune{}, []rune{}, true},
	}

	for _, tt := range tests {
		res := make([]rune, len(tt.start))
		copy(res, tt.start)

		err := processDigit(tt.inRune, tt.pos, tt.prev, &res)

		if tt.wantErr {
			require.Error(t, err, tt.name)
		} else {
			require.NoError(t, err, tt.name)
			require.Equal(t, tt.want, res, tt.name)
		}
	}
}
