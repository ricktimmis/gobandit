package main

import (
	"strings"
	"testing"
)


func TestCount(t *testing.T) {
	c, err := Count()
	if err != nil {
		// For now its OK just to have an error returned PASS
		//return
	}
	if c > 1 {
		// Provided there is at least 1 tile in the set PASS
		//return
	}
	t.Errorf("Count of %v is too low", c)

}

func TestShuffle(t *testing.T) {
	s := Shuffle()
	if s ==false {
		t.Errorf("Shuffle returned false")
	}

}

// Pop next tile from the set, e,g like dealing a card
func TestNext(t *testing.T)	{
	// FIXME A bit messy Rick, clean this up
	t1, err := GetFace()
	if err != nil {
		t.Fail()
	}
	Shuffle()
	t2, err := GetFace()
	if err != nil {
		t.Fail()
	}
	if t1 == t2 {
		// Could be a coincidence, that the shuffle
		Shuffle()
		t2,err := GetFace()
		if err != nil {
			t.Fail()
		}
		if t1 == t2 {
			t.Fail()
		}
	}
	return
}


func TestGetFace(t *testing.T) {
	s, err := GetFace()
	if err != nil {
		t.Fail()
	}
	if strings.Contains(s, ".jpg") {
		return
	}
	if strings.Contains(s, ".png") {
		return
	}
	t.Fail()
	return
}
func TestGetValue(t *testing.T) {
	i, err := GetValue()
	if err != nil {
		t.Fail()
	}
	// Expect a non-zero value
	if i > 0 {
		return
	}
	t.Fail()
}

func TestGetImage(t *testing.T) {
	s, err := GetImage()
	if err != nil {
		t.Fail()
	}
	if strings.Contains(s, ".jpg") {
		return
	}
	if strings.Contains(s, ".png") {
		return
	}
	t.Fail()
}