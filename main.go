package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	elements := make([]int, size)

	for i := range elements {
		elements[i] = rand.Intn(math.MaxInt) + 1
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxElement := data[0]
	for _, v := range data[1:] {
		if v > maxElement {
			maxElement = v
		}
	}

	return maxElement
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	parts := min(len(data), CHUNKS)
	chunkSize := len(data) / parts
	maxElements := make([]int, parts)

	for i := range parts {
		start := i * chunkSize
		end := start + chunkSize
		if i == parts-1 {
			end = len(data)
		}

		chunk := data[start:end]

		wg.Go(func() {
			maxElements[i] = maximum(chunk)
		})
	}
	wg.Wait()

	return maximum(maxElements)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxElement := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mсs\n", maxElement, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxElement = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mсs\n", maxElement, elapsed)
}
