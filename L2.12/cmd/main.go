package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrey67895/WB_Tech/L2.12/internal/grep"
	"github.com/andrey67895/WB_Tech/L2.12/internal/io"
)

func main() {
	// Флаги
	after := flag.Int("A", 0, "Показать N строк после совпадения")
	before := flag.Int("B", 0, "Показать N строк до совпадения")
	context := flag.Int("C", 0, "Показать N строк вокруг совпадения")
	count := flag.Bool("c", false, "Показать только количество совпадений")
	ignore := flag.Bool("i", false, "Игнорировать регистр")
	invert := flag.Bool("v", false, "Инвертировать результат")
	fixed := flag.Bool("F", false, "Фиксированная строка вместо regexp")
	linenum := flag.Bool("n", false, "Показывать номера строк")

	flag.Parse()

	// -C перекрывает -A и -B
	if *context > 0 {
		*after, *before = *context, *context
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Использование: grep [OPTIONS] PATTERN [FILE]")
		os.Exit(1)
	}

	pattern := args[0]
	filename := ""
	if len(args) > 1 {
		filename = args[1]
	}

	// читаем строки (STDIN или файл)
	lines, err := io.ReadLines(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения:", err)
		os.Exit(1)
	}

	// создаём матчёр
	m, err := grep.NewMatcher(pattern, *ignore, *fixed, *invert)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка компиляции паттерна:", err)
		os.Exit(1)
	}

	// запускаем поиск
	results := grep.Search(lines, m, grep.Options{
		Before:  *before,
		After:   *after,
		Count:   *count,
		LineNum: *linenum,
	})

	// если -c, печатаем только количество
	if *count {
		fmt.Println(len(results))
		return
	}

	// вывод строк
	for _, r := range results {
		if *linenum {
			fmt.Printf("%d:%s\n", r.LineNum, r.Line)
		} else {
			fmt.Println(r.Line)
		}
	}
}
