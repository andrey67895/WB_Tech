package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrey67895/WB_Tech/L2.10/internal/sortlogic"
)

// main — точка входа
func main() {
	Execute()
}

// Execute разбирает аргументы и запускает утилиту
func Execute() {
	opts, files := parseFlags()
	if err := sortlogic.Run(opts, files); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

// parseFlags разбирает аргументы командной строки
func parseFlags() (sortlogic.Options, []string) {
	var opts sortlogic.Options
	flag.IntVar(&opts.Column, "k", 0, "сортировать по N-й колонке (1-based, TAB разделитель)")
	flag.BoolVar(&opts.Numeric, "n", false, "сравнивать как числа")
	flag.BoolVar(&opts.Reverse, "r", false, "обратный порядок")
	flag.BoolVar(&opts.Unique, "u", false, "только уникальные строки")
	flag.BoolVar(&opts.Month, "M", false, "сортировка по названию месяца (Jan..Dec)")
	flag.BoolVar(&opts.TrimTail, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&opts.Check, "c", false, "проверить, отсортированы ли данные")
	flag.BoolVar(&opts.Human, "h", false, "сравнение человекочитаемых чисел (1K, 5M и т.п.)")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Использование: %s [флаги] [файл...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	return opts, flag.Args()
}
