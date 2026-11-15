package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	t.Run("нужной длины и все числа положительные", func(t *testing.T) {
		data := generateRandomElements(100)
		if len(data) != 100 {
			t.Fatalf("длина %d, ожидалась 100", len(data))
		}
		for i, value := range data {
			if value <= 0 {
				t.Fatalf("элемент %d = %d, ожидалось положительное число", i, value)
			}
		}
	})

	t.Run("размер 0", func(t *testing.T) {
		if data := generateRandomElements(0); len(data) != 0 {
			t.Fatalf("длина %d, ожидалась 0", len(data))
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"максимум в начале", []int{9, 1, 2, 3}, 9},
		{"максимум в середине", []int{1, 9, 2}, 9},
		{"максимум в конце", []int{1, 2, 3, 9}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
