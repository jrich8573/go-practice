package main

import "fmt"

func main() {

	// identical assignments
	//var card string = "Ace of Spades"
	card := newCard()

	fmt.Println(card)
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
*/
