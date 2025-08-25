package cli

import "testing"

func TestNormalizeDelimiter(t *testing.T) {
	if r, _ := normalizeDelimiter("\t"); r != '\t' {
		t.Fatalf("ожидался символ TAB")
	}
	if r, _ := normalizeDelimiter("\\t"); r != '\t' {
		t.Fatalf("ожидался символ TAB при вводе '\\t'")
	}
	if _, err := normalizeDelimiter(""); err == nil {
		t.Fatalf("ожидалась ошибка для пустого разделителя")
	}
	if _, err := normalizeDelimiter("::"); err == nil {
		t.Fatalf("ожидалась ошибка для разделителя из нескольких символов")
	}
}

func TestParseOptions_Ok(t *testing.T) {
	opts, err := ParseOptions(OptionsInput{FieldsSpec: "1,3-5", Delimiter: ","})
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}
	if !opts.Selector.Want(1) || !opts.Selector.Want(3) || !opts.Selector.Want(5) {
		t.Fatalf("селектор не содержит ожидаемые поля")
	}
}

func TestParseOptions_BadFields(t *testing.T) {
	_, err := ParseOptions(OptionsInput{FieldsSpec: "a-b"})
	if err == nil {
		t.Fatalf("ожидалась ошибка для неверного выражения полей")
	}
}
