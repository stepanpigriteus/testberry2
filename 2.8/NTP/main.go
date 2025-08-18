package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Fprintln(os.Stderr, "Ошибка: превышено время ожидания ответа от NTP-сервера")
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка получения времени:", err)
		}
		os.Exit(1)
	}
	fmt.Println(currentTime.Format(time.RFC1123))
}
