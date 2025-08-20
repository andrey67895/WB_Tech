package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

var defaultServers = []string{
	"0.ru.pool.ntp.org",
	"1.ru.pool.ntp.org",
	"2.ru.pool.ntp.org",
	"3.ru.pool.ntp.org",
}

func main() {
	timeout := flag.Duration("timeout", 5*time.Second, "таймаут для запроса к NTP")
	flag.Parse()

	currentTime, err := getTimeFromPool(defaultServers, *timeout)
	if err != nil {
		log.Fatalf("ошибка получения времени через NTP: %s", err)
	}

	fmt.Println(currentTime.Format(time.RFC3339))
}

// getTimeFromPool опрашивает список NTP-серверов и возвращает время с первого доступного.
func getTimeFromPool(servers []string, timeout time.Duration) (time.Time, error) {
	opts := ntp.QueryOptions{Timeout: timeout}

	var lastErr error
	for _, server := range servers {
		resp, err := ntp.QueryWithOptions(server, opts)
		if err != nil {
			lastErr = fmt.Errorf("сервер %s недоступен: %v", server, err)
			continue
		}
		return time.Now().Add(resp.ClockOffset), nil
	}

	return time.Time{}, fmt.Errorf("все NTP-серверы недоступны: %v", lastErr)
}
