package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Flagos struct {
	cd   bool
	pwd  bool
	echo bool
	kill bool
	ps   bool
}

func main() {
	var cfg Flagos
	args := parseFlag(&cfg)
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

	for {
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка ввода или EOF. Завершение.")
			break
		}
		fmt.Println("Введено:", input)
	}
}

func parseFlag(cfg *Flagos) []string {
	flag.BoolVar(&cfg.cd, "cd", false, "смена текущей директории")
	flag.BoolVar(&cfg.pwd, "pwd", false, "вывод текущей директории")
	flag.BoolVar(&cfg.echo, "echo", false, "вывод аргументов")
	flag.BoolVar(&cfg.kill, "kill", false, "послать сигнал завершения процессу с заданным PID")
	flag.BoolVar(&cfg.ps, "ps", false, "вывести список запущенных процессов")
	flag.Parse()

	args := flag.Args()
	return args
}

func catchSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	interrupt := make(chan bool, 1)

	<-sigs
	interrupt <- true
	fmt.Println("\nПрерывание команды (Ctrl+C)")
}
