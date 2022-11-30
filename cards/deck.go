package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	s := []string(d) //returns a slice of strings
	return strings.Join(s, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //bs == byte slice
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //string(bs)is type conversion byte slice to string
	return deck(s)
}
