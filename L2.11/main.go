package main

import (
	"fmt"
	"sort"
	"strings"
)

// sortRunes функция для сортировки букв в слове
func sortRunes(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// FindAnagrams основная функция для поиска анаграмм
func FindAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)
	order := make(map[string]string) // хранит первое встреченное слово для ключа

	for _, w := range words {
		word := strings.ToLower(w)
		sorted := sortRunes(word)

		if _, exists := anagramGroups[sorted]; !exists {
			order[sorted] = word // запоминаем первое слово
		}
		anagramGroups[sorted] = append(anagramGroups[sorted], word)
	}

	result := make(map[string][]string)
	for key, group := range anagramGroups {
		if len(group) > 1 { // исключаем одиночные слова
			sort.Strings(group)
			result[order[key]] = group
		}
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	res := FindAnagrams(words)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}
}
