package grep

import "testing"

func TestMatcher_Fixed(t *testing.T) {
	m, err := NewMatcher("foo", false, true, false)
	if err != nil {
		t.Fatal(err)
	}

	if !m.Match("foobar") {
		t.Error("ожидалось совпадение (fixed)")
	}
	if m.Match("barbaz") {
		t.Error("не ожидалось совпадение (fixed)")
	}
}

func TestMatcher_IgnoreCase(t *testing.T) {
	m, err := NewMatcher("foo", true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	if !m.Match("FOO bar") {
		t.Error("ожидалось совпадение при ignoreCase")
	}
}

func TestMatcher_Regexp(t *testing.T) {
	m, err := NewMatcher("^foo$", false, false, false)
	if err != nil {
		t.Fatal(err)
	}
	if !m.Match("foo") {
		t.Error("regexp не сработал")
	}
	if m.Match("foobar") {
		t.Error("regexp совпал неправильно")
	}
}

func TestMatcher_Invert(t *testing.T) {
	m, err := NewMatcher("foo", false, true, true)
	if err != nil {
		t.Fatal(err)
	}
	if m.Match("foo bar") {
		t.Error("invert должен исключить совпадение")
	}
	if !m.Match("bar") {
		t.Error("invert должен оставить строку")
	}
}
