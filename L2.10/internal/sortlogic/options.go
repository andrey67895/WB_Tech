package sortlogic

// Options описывает все поддерживаемые флаги сортировки.
type Options struct {
	Column   int  // -k N (колонка, 1-based; 0 = вся строка)
	Numeric  bool // -n (сравнивать как числа)
	Reverse  bool // -r (обратный порядок)
	Unique   bool // -u (только уникальные строки)
	Month    bool // -M (сортировка по названию месяца)
	TrimTail bool // -b (игнорировать хвостовые пробелы)
	Check    bool // -c (проверить, отсортированы ли данные)
	Human    bool // -h (сравнение человекочитаемых чисел)
}
