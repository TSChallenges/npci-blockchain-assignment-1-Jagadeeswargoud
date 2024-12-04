package main

import (
	"errors"
	"fmt"
	"slices"
)

// FilterAndSort filters numbers greater than or equal to the threshold and sorts them.
func FilterAndSort(nums []int, threshold int) []int {
	// TODO: Implement this function
	// Implemented the function as required by ranging thru slice of items and getting the numbers above the threshold
	newSlice := make([]int, 0)
	for _, num := range nums {
		if num >= threshold {
			newSlice = append(newSlice, num)
		}
	}
	slices.Sort(newSlice)
	fmt.Println(newSlice)
	return newSlice // Placeholder return
}

// FindMostFrequent finds the most frequent word in a slice of strings.
// Returns an error if the slice is empty.
func FindMostFrequent(words []string) (string, error) {
	// TODO: Implement this function
	// Handled the edge case where the slice is empty
	if len(words) == 0 {
		return "", errors.New("slice is empty")
	}
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	fmt.Println(wordMap)
	maxCount := 0
	maxWord := ""
	for word, count := range wordMap {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}
	return maxWord, nil // Placeholder return
}

func main() {
	// Test FilterAndSort
	fmt.Println("Testing FilterAndSort:")
	nums := []int{3, 10, 1, 7, 8, 2}
	threshold := 5
	fmt.Printf("Input: %v, Threshold: %d, Output: %v\n", nums, threshold, FilterAndSort(nums, threshold))

	// Test FindMostFrequent
	fmt.Println("\nTesting FindMostFrequent:")
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result, err := FindMostFrequent(words)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", words, result)
	}

	// Edge case: empty slice
	emptyWords := []string{}
	result, err = FindMostFrequent(emptyWords)
	if err != nil {
		fmt.Printf("Error: %v (empty input case handled)\n", err)
	} else {
		fmt.Printf("Input: %v, Most Frequent: %s\n", emptyWords, result)
	}
}
