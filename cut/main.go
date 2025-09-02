package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Cutter - настройки утилиты cut
type Cutter struct {
	fields     string
	sep        string
	onlySep    bool
	filePath   string
	indexes    []int
	rangeStart int
	toTheEnd   bool
}

// Run запускает обработку входных
func Run(args []string) error {
	c := &Cutter{
		filePath: args[len(args)-1],
		indexes:  make([]int, 0, 8),
	}

	fs := flag.NewFlagSet("cut", flag.ContinueOnError)
	fs.StringVar(&c.fields, "f", "", "choose columns")
	fs.StringVar(&c.sep, "d", "\t", "separator")
	fs.BoolVar(&c.onlySep, "s", false, "skip lines without separator")

	if err := fs.Parse(args); err != nil {
		return err
	}

	file, err := os.Open(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := c.prepareFields(); err != nil {
		return err
	}
	return c.handle(file)
}

func (c *Cutter) prepareFields() error {
	// перечисление через запятую
	if strings.Contains(c.fields, ",") {
		for _, part := range strings.Split(c.fields, ",") {
			num, err := strconv.Atoi(part)
			if err != nil {
				return err
			}
			c.indexes = append(c.indexes, num-1)
		}
		return nil
	}

	if len(c.fields) == 3 && c.fields[1] == '-' {
		from, err := strconv.Atoi(string(c.fields[0]))
		if err != nil {
			return err
		}
		to, err := strconv.Atoi(string(c.fields[2]))
		if err != nil {
			return err
		}
		for i := from; i <= to; i++ {
			c.indexes = append(c.indexes, i-1)
		}
		return nil
	}
	if len(c.fields) == 2 {
		if c.fields[0] == '-' {
			n, err := strconv.Atoi(string(c.fields[1]))
			if err != nil {
				return err
			}
			for i := 1; i <= n; i++ {
				c.indexes = append(c.indexes, i-1)
			}
			return nil
		}
		if c.fields[1] == '-' {
			n, err := strconv.Atoi(string(c.fields[0]))
			if err != nil {
				return err
			}
			c.rangeStart = n - 1
			c.toTheEnd = true
			return nil
		}
	}

	n, err := strconv.Atoi(c.fields)
	if err != nil {
		return err
	}
	c.indexes = append(c.indexes, n-1)
	return nil
}

func (c *Cutter) handle(f *os.File) error {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, c.sep)
		if len(parts) == 1 && c.onlySep {
			continue
		}
		if len(parts) == 1 {
			fmt.Println(parts[0])
			continue
		}
		if c.toTheEnd {
			fmt.Println(strings.Join(parts[c.rangeStart:], c.sep))
			continue
		}

		var out []string
		for _, idx := range c.indexes {
			if idx < len(parts) {
				out = append(out, parts[idx])
			}
		}
		fmt.Println(strings.Join(out, c.sep))
	}

	return scanner.Err()
}

func main() {
	if err := Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
	}
}
