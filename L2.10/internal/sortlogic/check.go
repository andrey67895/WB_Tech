package sortlogic

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// checkSorted проверяет, отсортированы ли данные.
func checkSorted(opts Options, files []string) error {
	var readers []io.Reader
	if len(files) == 0 {
		readers = []io.Reader{os.Stdin}
	} else {
		for _, name := range files {
			f, err := os.Open(name)
			if err != nil {
				return fmt.Errorf("не удалось открыть файл %s: %w", name, err)
			}
			defer f.Close()
			readers = append(readers, f)
		}
	}
	cmp := makeComparator(opts)
	lineNum := 0
	var prev string
	for _, r := range readers {
		s := bufio.NewScanner(r)
		s.Buffer(make([]byte, 1024), 10*1024*1024)
		for s.Scan() {
			curr := s.Text()
			lineNum++
			if lineNum > 1 {
				// Если текущая строка "меньше" предыдущей → порядок нарушен
				if cmp(curr, prev) {
					return fmt.Errorf("строка %d нарушает порядок сортировки", lineNum)
				}
			}
			prev = curr
		}
		if err := s.Err(); err != nil {
			return fmt.Errorf("ошибка чтения: %w", err)
		}
	}
	return nil
}
