package processing

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunPipe(first, second []string) {
	var cmd1 *exec.Cmd
	if first[0] == "ps" && len(first) == 1 {
		cmd1 = exec.Command("ps")
	} else {
		cmd1 = exec.Command(first[0], first[1:]...)
	}
	cmd2 := exec.Command(second[0], second[1:]...)
	pipe, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println("Ошибка при создании stdout пайпа:", err)
		return
	}
	cmd2.Stdin = pipe
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd1.Stderr = os.Stderr

	if err := cmd2.Start(); err != nil {
		fmt.Println("Ошибка запуска второй команды:", err)
		return
	}
	time.Sleep(10 * time.Millisecond)

	if err := cmd1.Start(); err != nil {
		fmt.Println("Ошибка запуска первой команды:", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println("Первая команда завершилась с ошибкой:", err)
	}
	if err := cmd2.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			fmt.Println("grep: совпадений не найдено")
		} else {
			fmt.Println("Ошибка выполнения второй команды:", err)
		}
	}
}

func RunMultiplePipes(commandStrings []string) {
	if len(commandStrings) < 2 {
		fmt.Println("Ошибка: недостаточно команд для пайпа")
		return
	}
	var commands []*exec.Cmd
	for i, cmdStr := range commandStrings {
		if cmdStr == "" {
			fmt.Printf("Ошибка: пустая команда в позиции %d\n", i+1)
			return
		}

		tokens := strings.Fields(cmdStr)
		if len(tokens) == 0 {
			fmt.Printf("Ошибка: пустая команда в позиции %d\n", i+1)
			return
		}

		cmd := createCommand(tokens)
		commands = append(commands, cmd)
	}

	for i := 0; i < len(commands)-1; i++ {
		pipe, err := commands[i].StdoutPipe()
		if err != nil {
			fmt.Printf("Ошибка создания пайпа между ком %d и %d: %v\n", i+1, i+2, err)
			return
		}
		commands[i+1].Stdin = pipe
	}
	commands[0].Stdin = os.Stdin
	commands[len(commands)-1].Stdout = os.Stdout
	commands[len(commands)-1].Stderr = os.Stderr
	for i := 0; i < len(commands)-1; i++ {
		commands[i].Stderr = os.Stderr
	}

	for i := len(commands) - 1; i >= 0; i-- {
		if err := commands[i].Start(); err != nil {
			fmt.Printf("Ошибка запуска команды %d (%s): %v\n", i+1, commandStrings[i], err)
			return
		}
		if i > 0 {
			time.Sleep(5 * time.Millisecond)
		}
	}
	for i, cmd := range commands {
		if err := cmd.Wait(); err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				if strings.Contains(commandStrings[i], "grep") && exitErr.ExitCode() == 1 {
					if i == len(commands)-1 {
						fmt.Println("grep: совпадений не найдено")
					}
				} else {
					fmt.Printf("Команда %d (%s) завершилась с ошибкой: %v\n", i+1, commandStrings[i], err)
				}
			} else {
				fmt.Printf("Команда %d (%s) завершилась с ошибкой: %v\n", i+1, commandStrings[i], err)
			}
		}
	}
}

func createCommand(tokens []string) *exec.Cmd {
	cmd := tokens[0]

	switch cmd {
	case "ps":
		return exec.Command("ps")
	case "echo":
		if len(tokens) > 1 {
			return exec.Command("echo", strings.Join(tokens[1:], " "))
		} else {
			return exec.Command("echo")
		}
	default:
		if len(tokens) > 1 {
			return exec.Command(tokens[0], tokens[1:]...)
		} else {
			return exec.Command(tokens[0])
		}
	}
}
