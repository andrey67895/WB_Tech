package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andrey67895/WB_Tech/L2.13/internal/cli"
	"github.com/andrey67895/WB_Tech/L2.13/internal/cut"
)

func main() {
	// Флаги CLI
	fieldsSpec := flag.String("f", "", "Номера полей/диапазоны (напр.: 1,3-5)")
	delimiter := flag.String("d", "\t", "Разделитель (один символ). По умолчанию: TAB")
	separated := flag.Bool("s", false, "Только строки, содержащие разделитель")
	flag.Parse()

	if *fieldsSpec == "" {
		fmt.Fprintln(os.Stderr, "ошибка: нужен флаг -f (напр.: -f 1,3-5)")
		flag.Usage()
		os.Exit(2)
	}

	opts, err := cli.ParseOptions(cli.OptionsInput{
		FieldsSpec: *fieldsSpec,
		Delimiter:  *delimiter,
		Separated:  *separated,
	})
	if err != nil {
		log.Fatalf("парсинг опций: %v", err)
	}

	if err := cut.Run(os.Stdin, os.Stdout, opts); err != nil {
		log.Fatalf("ошибка выполнения: %v", err)
	}
}
