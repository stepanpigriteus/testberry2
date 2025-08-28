package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type config struct {
	fields    []int
	delimiter string
	separated bool
}

func main() {
	var fieldsStr string
	var cfg config

	flag.StringVar(&fieldsStr, "f", "", "выбор полей (колонок, например, '1,3-5')")
	flag.StringVar(&cfg.delimiter, "d", "\t", "использовать указанный разделитель вместо табуляции")
	flag.BoolVar(&cfg.separated, "s", false, "выводить только строки, содержащие разделитель")
	flag.Parse()

	if fieldsStr == "" {
		fmt.Fprintln(os.Stderr, "Использование: cut -f список [-d разделитель] [-s] [файл...]")
		fmt.Fprintln(os.Stderr, "Ошибка: флаг -f обязателен")
		os.Exit(1)
	}

	var err error
	cfg.fields, err = parseFields(fieldsStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при разборе полей: %v\n", err)
		os.Exit(1)
	}

	// Обработка входных данных из STDIN или файлов
	args := flag.Args()
	var scanners []*bufio.Scanner
	if len(args) == 0 {
		scanners = append(scanners, bufio.NewScanner(os.Stdin))
	} else {
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка при открытии файла %s: %v\n", file, err)
				continue
			}
			scanners = append(scanners, bufio.NewScanner(f))
			defer f.Close()
		}
	}

	var sb strings.Builder
	for _, scanner := range scanners {
		if err := processInput(scanner, &cfg, &sb); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при чтении данных: %v\n", err)
			os.Exit(1)
		}
	}
}

func parseFields(s string) ([]int, error) {
	var fields []int
	seen := make(map[int]bool)
	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("неверный диапазон: %s", part)
			}
			startStr, endStr := strings.TrimSpace(rangeParts[0]), strings.TrimSpace(rangeParts[1])
			start, err := strconv.Atoi(startStr)
			if err != nil || start < 1 {
				return nil, fmt.Errorf("неверное начало диапазона: %s", startStr)
			}
			end, err := strconv.Atoi(endStr)
			if err != nil || end < 1 {
				return nil, fmt.Errorf("неверный конец диапазона: %s", endStr)
			}
			if start > end {
				return nil, fmt.Errorf("начало диапазона не может быть больше конца: %s", part)
			}
			for i := start; i <= end; i++ {
				if !seen[i] {
					fields = append(fields, i)
					seen[i] = true
				}
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 {
				return nil, fmt.Errorf("неверный номер поля: %s", part)
			}
			if !seen[num] {
				fields = append(fields, num)
				seen[num] = true
			}
		}
	}
	if len(fields) == 0 {
		return nil, fmt.Errorf("не указано ни одного валидного поля")
	}
	return fields, nil
}

func processInput(scanner *bufio.Scanner, cfg *config, sb *strings.Builder) error {
	for scanner.Scan() {
		line := scanner.Text()
		hasDelim := strings.Contains(line, cfg.delimiter)
		if cfg.separated && !hasDelim {
			continue
		}

		var lineFields []string
		if hasDelim {
			lineFields = strings.Split(line, cfg.delimiter)
		} else {
			lineFields = []string{line}
		}

		sb.Reset()
		for i, fieldNum := range cfg.fields {
			idx := fieldNum - 1
			if idx < len(lineFields) {
				sb.WriteString(lineFields[idx])
			}
			if i < len(cfg.fields)-1 {
				sb.WriteString(cfg.delimiter)
			}
		}
		fmt.Println(sb.String())
	}
	return scanner.Err()
}
