package main

import "fmt"

// create a new type of deck, which is a slice of strings
type deck []string

func newDeck() deck { // return a value of type deck
	// doesn't need a receiver
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// function to prints every card in type deck
func (d deck) print() { // (d deck) is called a receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // return two values of type deck
	return d[:handSize], d[handSize:]
}
