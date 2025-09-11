package processing

import (
	"fmt"
	"os"
)

func Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка при получении текущей директории:", err)
		return
	}

	fmt.Println("Текущая директория:", dir)
}
