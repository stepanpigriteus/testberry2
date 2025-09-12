package processing

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func Cmd(token, tokenNext []string) {
	cmd := token[0]
	if tokenNext != nil && len(tokenNext) > 0 {
		RunPipe(token, tokenNext)
		return
	}

	switch cmd {
	case "pwd":
		Pwd()
	case "cd":
		Cd(token)
	case "exit":
		fmt.Println("Выход.")
		os.Exit(0)
	case "echo":
		if len(token) > 1 {
			fmt.Println(strings.Join(token[1:], " "))
		} else {
			fmt.Println()
		}
	case "ps":
		ps := exec.Command("ps")
		ps.Stdout = os.Stdout
		ps.Stderr = os.Stderr
		if err := ps.Run(); err != nil {
			fmt.Println("Ошибка выполнения команды ps:", err)
		}
	case "kill":
		if len(token) < 2 {
			fmt.Println("kill: отсутствует PID")
			break
		}
		pid, err := strconv.Atoi(token[1])
		if err != nil {
			fmt.Println("kill: некорректный PID:", token[1])
			break
		}
		if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
			fmt.Println("kill: ошибка:", err)
		} else {
			fmt.Println("Процесс", pid, "завершён (SIGTERM)")
		}
	default:
		external := exec.Command(cmd, token[1:]...)
		external.Stdin = os.Stdin
		external.Stdout = os.Stdout
		external.Stderr = os.Stderr
		if err := external.Run(); err != nil {
			fmt.Println("Ошибка запуска команды:", cmd, "-", err)
		}
	}
}
