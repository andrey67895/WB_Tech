package grep

import "testing"

func TestSearch_BasicMatch(t *testing.T) {
	lines := []string{"a", "foo", "b"}
	m, _ := NewMatcher("foo", false, true, false)

	results := Search(lines, m, Options{})
	if len(results) != 1 {
		t.Errorf("ожидался 1 результат, получили %d", len(results))
	}
	if results[0].Line != "foo" {
		t.Errorf("ожидалась строка foo, получили %s", results[0].Line)
	}
}

func TestSearch_WithContext(t *testing.T) {
	lines := []string{"a", "foo", "b", "c"}
	m, _ := NewMatcher("foo", false, true, false)

	results := Search(lines, m, Options{Before: 1, After: 2})
	if len(results) != 4 {
		t.Errorf("ожидалось 4 строки (включая контекст), получили %d", len(results))
	}
}

func TestSearch_CountOnly(t *testing.T) {
	lines := []string{"foo", "bar", "foo"}
	m, _ := NewMatcher("foo", false, true, false)

	results := Search(lines, m, Options{Count: true})
	if len(results) != 2 {
		t.Errorf("ожидалось 2 совпадения, получили %d", len(results))
	}
}
