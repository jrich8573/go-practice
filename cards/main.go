package main

import "fmt"

func main() {

	// identical assignments
	// var card string = "Ace of Spades"
	// card := newCard()
	cards := []string{"Ace of Diamonds", newCard()} // new cards
	cards = append(cards, "Six of Spades")          // add a card to cards

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string { // string return type

	return "Five of Diamonds"
}

/*
Notes:
1. go is statically typed language similar to c++ and java
2. declared var is immutable
3. Basic data types:
	a. bool (true/false)
	b. string (" ")
	c. int (1, -1000, 999)
	d. float64 (10.0003, 0.09999, -100.3)
4. only use the := when assigning a new variable (when reassigning a variable use the = )
5. go allows you to print the return value of a function without assigning it to a var
6. go has 2 data structures for handling lists of data.
	a. arrary = fixed length list of data
	b. slice = an array that can grow and shrink
7. in a slice = every element must be of the same type
*/
