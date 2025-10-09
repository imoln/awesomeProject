package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var word1, word2 string

	fmt.Println("First word: ")
	fmt.Scanln(&word1)

	fmt.Println("Second word: ")
	fmt.Scanln(&word2)

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	arr1 := strings.Split(word1, " ")
	arr2 := strings.Split(word2, " ")

	sort.Strings(arr1)
	sort.Strings(arr2)

	if strings.Join(arr1, " ") == strings.Join(arr2, " ") {
		fmt.Println("Anagrams")
	} else {
		fmt.Println("Not Anagrams")
	}
}
