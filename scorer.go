package main

import "fmt"

type score struct {
	bd *Board

}

type scorer interface {

}

func (s *score) Evaluate (b *Board) int {

	// logic goes here
	return 0
}
func (s *score) LoadRules (f func(b *Board) int) int{
	result := f(s.bd)
	return result
}

func anotherFunction(f func(string) string) {
	result := f("David")
	fmt.Println(result) // Prints "Hiya, David"
}