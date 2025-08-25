package cut

import (
	"bytes"
	"strings"
	"testing"
)

func TestFieldSelector(t *testing.T) {
	s := FieldSelector{}
	s.Add(1)
	s.AddRange(3, 5)
	if !s.Want(1) || !s.Want(3) || !s.Want(5) || s.Want(2) || s.Want(6) {
		t.Fatalf("ошибка в логике селектора полей")
	}
}

func TestRun_Basic(t *testing.T) {
	in := "a\tb\tc\n1\t2\t3\n"
	var out bytes.Buffer
	opts := Options{Delimiter: '\t', Selector: FieldSelector{ranges: [][2]int{{1, 1}, {3, 5}}}}
	if err := Run(strings.NewReader(in), &out, opts); err != nil {
		t.Fatalf("неожиданная ошибка выполнения: %v", err)
	}
	got := out.String()
	want := "a\tc\n1\t3\n"
	if got != want {
		t.Fatalf("результат не совпадает: получили %q, ожидалось %q", got, want)
	}
}

// Берём ровно второй столбец при разделителе ','
func TestRun_DifferentDelimiter(t *testing.T) {
	in := "a,b,c\n1,2,3\n"
	var out bytes.Buffer
	// ВАЖНО: если нужен только второй столбец — диапазон 2..2
	opts := Options{Delimiter: ',', Selector: FieldSelector{ranges: [][2]int{{2, 2}}}}
	if err := Run(strings.NewReader(in), &out, opts); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if out.String() != "b\n2\n" {
		t.Fatalf("результат не совпадает: получили %q, ожидалось %q", out.String(), "b\n2\n")
	}
}

// Диапазон 2-4 на трёхколоночной строке даёт 2-3, т.к. выход за границы игнорируется
func TestRun_RangeOutOfBounds(t *testing.T) {
	in := "a,b,c\n1,2,3\n"
	var out bytes.Buffer
	opts := Options{Delimiter: ',', Selector: FieldSelector{ranges: [][2]int{{2, 4}}}}
	if err := Run(strings.NewReader(in), &out, opts); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	want := "b,c\n2,3\n"
	if out.String() != want {
		t.Fatalf("результат не совпадает: получили %q, ожидалось %q", out.String(), want)
	}
}

func TestRun_SeparatedOnly(t *testing.T) {
	in := "no_delim_line\nwith,delim,line\n"
	var out bytes.Buffer
	opts := Options{Delimiter: ',', SeparatedOnly: true, Selector: FieldSelector{ranges: [][2]int{{1, 2}}}}
	if err := Run(strings.NewReader(in), &out, opts); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if out.String() != "with,delim\n" {
		t.Fatalf("результат не совпадает: получили %q, ожидалось %q", out.String(), "with,delim\n")
	}
}

func TestRun_LongLine(t *testing.T) {
	bigField := strings.Repeat("x", 200000)
	in := bigField + ",y\n"
	var out bytes.Buffer
	opts := Options{Delimiter: ',', Selector: FieldSelector{ranges: [][2]int{{1, 1}}}}
	if err := Run(strings.NewReader(in), &out, opts); err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if out.String() != bigField+"\n" {
		t.Fatalf("результат имеет неверную длину: %d символов", len(out.String()))
	}
}
