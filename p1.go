// 1. Реализуйте программу, которая создает массив случайных строк и сортирует их по длине.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	strings := generateRandomStrings(10, 5, 15)
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})

	fmt.Println("\nОтсортированный массив:")
	printStrings(strings)
}

func generateRandomStrings(count, minLen, maxLen int) []string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]string, count)

	for i := 0; i < count; i++ {
		length := rand.Intn(maxLen-minLen+1) + minLen
		runes := make([]rune, length)
		for j := range runes {
			runes[j] = letters[rand.Intn(len(letters))]
		}
		result[i] = string(runes)
	}

	return result
}

func printStrings(strings []string) {
	for i, s := range strings {
		fmt.Printf("%d. %s (длина: %d)\n", i+1, s, len(s))
	}
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
