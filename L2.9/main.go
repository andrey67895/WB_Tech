package main

import (
	"fmt"
	"log"

	"github.com/andrey67895/WB_Tech/L2.9/unpack"
)

func main() {
	examples := []string{
		"a4bc2d5e",
		"abcd",
		"",
		"45",
		"qwe\\4\\5",
		"qwe\\45",
		"aaa0b",
	}

	for _, in := range examples {
		out, err := unpack.Unpack(in)
		if err != nil {
			log.Printf("Вход: %s → ошибка: %s\n", in, err.Error())
		} else {
			fmt.Printf("Вход: %s → Выход: %s\n", in, out)
		}
	}
}
