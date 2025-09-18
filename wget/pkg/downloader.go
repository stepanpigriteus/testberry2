package pkg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Downloader(rawURL, dest, dir string) ([]byte, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("не удалось создать директорию %w", err)
	}
	path := filepath.Join(dir, dest)
	out, err := os.Create(path)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return nil, fmt.Errorf("нет прав на запись файла: %w", err)
		}
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("каталог назначения не существует: %w", err)
		}
		return nil, err
	}
	defer out.Close()

	resp, err := http.Get(rawURL)
	if err != nil {
		var uerr *url.Error
		if errors.As(err, &uerr) {
			return nil, fmt.Errorf("ошибка сети: %w", uerr.Err)
		}
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка загрузки: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %w", err)
	}
	if _, err := out.Write(data); err != nil {
		return nil, fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return data, nil
}
