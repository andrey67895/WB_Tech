package cut

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// Options — параметры работы утилиты.
type Options struct {
	Delimiter     rune
	SeparatedOnly bool
	Selector      FieldSelector
}

// FieldSelector — хранит множество нужных полей (1-базная индексация).
// Для эффективности — список закрытых диапазонов [start,end].
type FieldSelector struct {
	ranges [][2]int
}

func (s *FieldSelector) Add(v int)         { s.ranges = append(s.ranges, [2]int{v, v}) }
func (s *FieldSelector) AddRange(a, b int) { s.ranges = append(s.ranges, [2]int{a, b}) }
func (s FieldSelector) Want(idx int) bool {
	for _, r := range s.ranges {
		if idx >= r[0] && idx <= r[1] {
			return true
		}
	}
	return false
}

// Run — основной цикл: читает из r, пишет в w.
func Run(r io.Reader, w io.Writer, opts Options) error {
	if len(opts.Selector.ranges) == 0 {
		return fmt.Errorf("не указаны поля для вывода")
	}

	br := bufio.NewReader(r)
	for {
		line, err := readLongLine(br)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		trimmed, nl := trimNewline(line)

		// -s: пропускаем строки без разделителя.
		if opts.SeparatedOnly && !bytesContainsRune(trimmed, opts.Delimiter) {
			continue
		}

		out := selectFields(trimmed, opts.Delimiter, opts.Selector)
		if len(out) > 0 {
			if _, err := w.Write(out); err != nil {
				return err
			}
		}
		if nl != nil {
			if _, err := w.Write(nl); err != nil {
				return err
			}
		}
	}
	return nil
}

// readLongLine — читает строку любой длины (включая завершающий '\n', если есть).
func readLongLine(r *bufio.Reader) ([]byte, error) {
	var buf []byte
	for {
		chunk, err := r.ReadSlice('\n')
		if err == bufio.ErrBufferFull {
			buf = append(buf, chunk...)
			continue
		}
		buf = append(buf, chunk...)
		return buf, err
	}
}

// trimNewline убирает финальный перевод строки и возвращает (строка без \r?\n, что убрать, чем закончить).
func trimNewline(b []byte) (line []byte, nl []byte) {
	if len(b) == 0 {
		return b, nil
	}
	if b[len(b)-1] == '\n' {
		if len(b) >= 2 && b[len(b)-2] == '\r' {
			return b[:len(b)-2], []byte("\r\n")
		}
		return b[:len(b)-1], []byte("\n")
	}
	return b, nil
}

func bytesContainsRune(b []byte, r rune) bool {
	if r <= 0x7F {
		return bytes.IndexByte(b, byte(r)) >= 0
	}
	return bytes.Contains(b, []byte(string(r)))
}

// selectFields — одно-проходное выделение нужных полей.
// Для ASCII-разделителя работает по байтам; для юникодного — падение на Split.
func selectFields(line []byte, delim rune, selector FieldSelector) []byte {
	var out bytes.Buffer
	fieldIdx := 1
	start := 0

	// Быстрый путь: ASCII разделитель.
	if delim <= 0x7F {
		db := byte(delim)
		for i := 0; i <= len(line); i++ {
			if i == len(line) || line[i] == db {
				if selector.Want(fieldIdx) {
					if out.Len() > 0 {
						out.WriteByte(db)
					}
					out.Write(line[start:i])
				}
				fieldIdx++
				start = i + 1
			}
		}
		return out.Bytes()
	}

	// Общий случай: юникодный разделитель (редко нужен).
	parts := bytes.Split(line, []byte(string(delim)))
	for i, p := range parts {
		idx := i + 1
		if selector.Want(idx) {
			if out.Len() > 0 {
				out.WriteString(string(delim))
			}
			out.Write(p)
		}
	}
	return out.Bytes()
}
