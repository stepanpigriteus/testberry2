package processing

import (
	"fmt"
	"os"
)

func Cd(tokens []string) {
	if len(tokens) < 2 {
		fmt.Println("cd: отсутствует аргумент пути")
	} else {
		err := os.Chdir(tokens[1])
		if err != nil {
			fmt.Println("cd: ошибка:", err)
		}
	}
}
