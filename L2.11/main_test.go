package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Пример из условия",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "Одиночные слова игнорируются",
			input:    []string{"дом", "кот", "собака"},
			expected: map[string][]string{
				// пусто
			},
		},
		{
			name:  "С дубликатами",
			input: []string{"Пятак", "пятак", "тяпка", "ПЯТКА"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "Пустой ввод",
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindAnagrams(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("не совпало для %q\nввод: %#v\nожидал: %#v\nполучил: %#v",
					tt.name, tt.input, tt.expected, got)
			}
		})
	}
}
