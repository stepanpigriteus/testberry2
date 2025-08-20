package main

import "testing"

func TestUnpackStringEdgeCases(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
		hasError bool
	}{
		"empty":          {"", "", false},
		"single_char":    {"x", "x", false},
		"zero_repeat":    {"a0", "", false},
		"one_repeat":     {"a1", "a", false},
		"max_digit":      {"a9", "aaaaaaaaa", false},
		"unicode":        {"я3", "яяя", false},
		"mixed_unicode":  {"a2я3", "aaяяя", false},
		"escape_unicode": {`\я2`, "яя", false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := unpackString(tt.input)

			if tt.hasError && err == nil {
				t.Errorf("Ожидалась ошибка")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Неожиданная ошибка: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Ожидалось %q, получено %q", tt.expected, result)
			}
		})
	}
}
