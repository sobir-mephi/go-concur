package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	SIZE        = 100_000_000
	CHUNKS      = 8
	MAX_ELEMENT = 1_000_000
)

// generateRandomElements создаёт слайс из size случайных положительных чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.IntN(MAX_ELEMENT) + 1
	}

	return data
}

// maximum возвращает максимальный элемент слайса.
// Не менял исходную сигнатуру, но лучше бы сделать (int, error)
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks параллельный поиск максимуа.
// Не менял исходную сигнатуру, но лучше бы сделать (int, error)
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		start := chunkIndex * chunkSize
		end := start + chunkSize
		if chunkIndex == CHUNKS-1 {
			end = len(data) // последний чанк
		}

		wg.Add(1)
		go func(chunk []int, resultIndex int) {
			defer wg.Done()
			chunkMaximums[resultIndex] = maximum(chunk)
		}(data[start:end], chunkIndex)
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startSingle := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startSingle)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSingle, elapsedSingle.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startChunks := time.Now()
	maxParallel := maxChunks(data)
	elapsedChunks := time.Since(startChunks)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxParallel, elapsedChunks.Microseconds())
}
