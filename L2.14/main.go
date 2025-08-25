package main

import (
	"fmt"
	"reflect"
	"time"
)

// При большом кол-ве каналов начнутся проблемы
// or объединяет несколько done-каналов в один.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func or2(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		// подготовим список случаев для select
		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			}
		}

		// как только хоть один канал закроется -> завершаем
		reflect.Select(cases)
	}()
	return orDone
}

// вспомогательная функция: вернёт канал, который закроется через after
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or2(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
