package io

import (
	"os"
	"testing"
)

func TestReadLines_FromStdin(t *testing.T) {
	// временный файл с тестовыми строками
	content := "line1\nline2\nline3\n"
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tmpfile.Close()

	lines, err := ReadLines(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if len(lines) != 3 {
		t.Errorf("ожидалось 3 строки, получили %d", len(lines))
	}
	if lines[0] != "line1" || lines[2] != "line3" {
		t.Errorf("строки прочитаны неправильно: %+v", lines)
	}
}

func TestReadLines_FileNotFound(t *testing.T) {
	_, err := ReadLines("no_such_file.txt")
	if err == nil {
		t.Error("ожидалась ошибка при отсутствии файла")
	}
}
