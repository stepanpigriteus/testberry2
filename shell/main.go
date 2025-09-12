package main

import (
	"bufio"
	"fmt"
	"minishell/processing"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	go func() {
		for sig := range sigs {
			if sig == syscall.SIGINT {
				fmt.Println("\nОбрыв команды (Ctrl+C)")
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода или EOF. Завершение.")
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		commands := strings.Split(input, "|")
		for i := range commands {
			commands[i] = strings.TrimSpace(commands[i])
		}

		if len(commands) == 1 {
			token := strings.Fields(commands[0])
			if len(token) > 0 {
				processing.Cmd(token, nil)
			}
		} else {
			processing.RunMultiplePipes(commands)
		}
	}
}
