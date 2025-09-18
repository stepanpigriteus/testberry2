package pkg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FileCreator(path string) {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Не удалось создать директорию:", err)
		return
	}
	newPath := path
	for i := 1; ; i++ {
		file, err := os.OpenFile(newPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0666)
		if err != nil {
			if errors.Is(err, os.ErrPermission) {
				fmt.Println("Нет прав на создание файла")
				return
			}
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("Родительский путь отсутствует")
				return
			}
			if errors.Is(err, os.ErrExist) {
				newPath = filepath.Join(dir, fmt.Sprintf("%s(%d)%s", name, i, ext))
				continue
			}
			fmt.Println("Ошибка:", err)
			return
		}

		defer file.Close()
		fmt.Println("Файл создан:", newPath)
		return
	}
}
