package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var arr []string = []string{"Пятак", "пятка", "тяпка", "атпкя", "cапог"}
	anagramm(arr)
}

func anagramm(s []string) map[string][]string {
	prev := make(map[string][]string)
	for _, r := range s {

		runes := []rune(strings.ToLower(r))
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		prev[string(runes)] = append(prev[string(runes)], r)
	}
	for key, value := range prev {
		if len(value) < 2 {
			delete(prev, key)
		}
		sort.Strings(prev[key])
	}

	fmt.Println(prev)
	return prev
}

// технически условие соблюдено - так кк вторичная сортировка не будет занимать больше О (n log n)
// ессли нужно O(n * m log m) - то просто убиираемм ссортировку во втором цикле

