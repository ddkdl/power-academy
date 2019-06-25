package main

import "fmt"

func main() {
	var word string
	var otherWord string

	fmt.Print("Input first word: ")
	fmt.Scan(&word)

	fmt.Print("Input second word: ")
	fmt.Scan(&otherWord)

	n_1 := len(word) - 3
	n_2 := len(otherWord) - 3

	if word[n_1:] == otherWord[n_2:] {
		fmt.Println("These words rhyme!")
	} else {
		fmt.Println("These words do not rhyme!")
	}
}
