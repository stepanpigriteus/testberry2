package parseflags

import "flag"

type Flagos struct {
	cd   bool
	pwd  bool
	echo bool
	kill bool
	ps   bool
}

func FlagParser(cfg *Flagos) []string {
	flag.BoolVar(&cfg.cd, "cd", false, "смена текущей директории")
	flag.BoolVar(&cfg.pwd, "pwd", false, "вывод текущей директории")
	flag.BoolVar(&cfg.echo, "echo", false, "вывод аргументов")
	flag.BoolVar(&cfg.kill, "kill", false, "послать сигнал завершения процессу с заданным PID")
	flag.BoolVar(&cfg.ps, "ps", false, "вывести список запущенных процессов")
	flag.Parse()

	args := flag.Args()
	return args
}
