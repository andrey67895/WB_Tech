package sortlogic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

// Run запускает основную логику: сортировку или проверку порядка.
func Run(opts Options, files []string) error {
	if opts.Check {
		return checkSorted(opts, files)
	}

	lines, err := readAllLines(files)
	if err != nil {
		return fmt.Errorf("не удалось прочитать входные данные: %w", err)
	}

	cmp := makeComparator(opts)
	sort.Slice(lines, func(i, j int) bool { return cmp(lines[i], lines[j]) })

	if opts.Unique {
		lines = uniqueAdjacent(lines, func(a, b string) bool {
			return !cmp(a, b) && !cmp(b, a)
		})
	}
	return writeLines(os.Stdout, lines)
}

// readAllLines читает все строки из файлов (или stdin, если файлов нет).
func readAllLines(files []string) ([]string, error) {
	var readers []io.Reader
	if len(files) == 0 {
		readers = []io.Reader{os.Stdin}
	} else {
		for _, name := range files {
			f, err := os.Open(name)
			if err != nil {
				return nil, fmt.Errorf("не удалось открыть файл %s: %w", name, err)
			}
			defer f.Close()
			readers = append(readers, f)
		}
	}

	var lines []string
	for _, r := range readers {
		s := bufio.NewScanner(r)
		s.Buffer(make([]byte, 1024), 10*1024*1024) // до 10МБ строка
		for s.Scan() {
			lines = append(lines, s.Text())
		}
		if err := s.Err(); err != nil {
			return nil, fmt.Errorf("ошибка чтения: %w", err)
		}
	}
	return lines, nil
}

// writeLines выводит строки на stdout.
func writeLines(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, line := range lines {
		if _, err := bw.WriteString(line); err != nil {
			return fmt.Errorf("ошибка записи: %w", err)
		}
		if err := bw.WriteByte('\n'); err != nil {
			return fmt.Errorf("ошибка записи: %w", err)
		}
	}
	return bw.Flush()
}
