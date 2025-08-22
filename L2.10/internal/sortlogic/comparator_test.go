package sortlogic

import "testing"

func TestNumericComparator(t *testing.T) {
	opts := Options{Numeric: true}
	cmp := makeComparator(opts)

	if !cmp("2", "10") {
		t.Error("2 должно быть меньше 10")
	}
}

func TestReverseComparator(t *testing.T) {
	opts := Options{Numeric: true, Reverse: true}
	cmp := makeComparator(opts)

	if cmp("2", "10") {
		t.Error("в обратном порядке 2 не должно быть меньше 10")
	}
}

func TestColumnComparator(t *testing.T) {
	opts := Options{Column: 2}
	cmp := makeComparator(opts)

	a := "foo\t10"
	b := "bar\t20"
	if !cmp(a, b) {
		t.Errorf("%q должно быть меньше %q по 2-й колонке", a, b)
	}
}

func TestUniqueAdjacent(t *testing.T) {
	lines := []string{"a", "a", "b", "b", "c"}
	out := uniqueAdjacent(lines, func(a, b string) bool { return a == b })

	want := []string{"a", "b", "c"}
	if len(out) != len(want) {
		t.Fatalf("ожидалось %v, получили %v", want, out)
	}
	for i := range want {
		if out[i] != want[i] {
			t.Errorf("позиция %d: ожидалось %q, получили %q", i, want[i], out[i])
		}
	}
}
