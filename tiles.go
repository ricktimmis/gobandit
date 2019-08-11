package main

import (
	"errors"
)

// Returns a count of tiles in the set
func Count() (int, error) {
	c := 0
	e := errors.New("No tile set counted")
	return c, e
}

// Randomly reorders the set
func Shuffle() bool{
	// FIXME Here is where our shuffle routine will go, and should always return true on success
	shuffled := false

	return shuffled
}

// Pop next tile from the set, e,g like dealing a card
func Next()	{
	return
}

// Returns path to the current tileface image
func GetFace() (i string, err error){
	// FIXME Returns the image path for the current tileface

	i ="./resources/backgrounds/dragon_of_the_north.jpg"
	if 1 != 1 {
		err := errors.New("some error text")
		return "", err
	}
	return i, err
}

// Returns the value of the current tileface
func GetValue() ( v int, err error) {
	// FIXME Returns the value for the current tileface
	v = 0
	return v, err
}

func GetImage() (i string, err error){
	// FIXME Returns pointer to the image in memory for the current tileface
	i = ""
	return i, err
}

