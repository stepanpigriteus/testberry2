package main

import (
	"bufio"
	"fmt"
	"minishell/parseflags"
	"minishell/processing"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	var cfg parseflags.Flagos
	args := parseflags.FlagParser(&cfg)
	fmt.Println(cfg, args)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for sig := range sigs {
			if sig == syscall.SIGINT {
				fmt.Println("\nПрерывание команды (Ctrl+C)")
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

		tokens := strings.Fields(input)
		cmd := tokens[0]

		switch cmd {
		case "pwd":
			processing.Pwd()
		case "cd":
			processing.Cd(tokens)
		case "exit":
			fmt.Println("Выход.")
			return
		case "echo":
			if len(tokens) > 1 {
				fmt.Println(strings.Join(tokens[1:], " "))
			} else {
				fmt.Println()
			}
		case "ps":
			cmd := exec.Command("ps", "-e", "-o", "pid,ppid,comm")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды ps:", err)
			}
		case "kill":
			if len(tokens) < 2 {
				fmt.Println("kill: отсутствует PID")
				break
			}

			pid, err := strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Println("kill: некорректный PID:", tokens[1])
				break
			}

			err = syscall.Kill(pid, syscall.SIGTERM)
			if err != nil {
				fmt.Println("kill: ошибка:", err)
			} else {
				fmt.Println("Процесс", pid, "завершён (SIGTERM)")
			}
		default:
			fmt.Println("Неизвестная команда:", cmd)
		}
	}
}
