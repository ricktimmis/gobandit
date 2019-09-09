package main

import (
	"fmt"
)

/* Score provides an evaluation component for calculating rules
   against the board. The simplest usage is to initialise a variable with
   a function, and then call the evaluate method injecting the function via the variable.
 */
type score struct {
	bd *Board

}

type scorer interface {
	init()
	evaluate()
}

func (s *score) init (p *Board) error {
	if p == nil {
		return fmt.Errorf("Failed passing board pointer")
	}
	s.bd = p
	return nil
}

// FIXME Issue #3 - Recursive Scoring evaluation - It would be great to pass in a set of rules to evaluate, rather than
//                  making multiple calls.
func (s *score) evaluate (f func(b *Board)int) (int, error){
	var err = error(nil)
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("Evaluation failed with error: %v",err)
		}
	}()
	result := f(s.bd)
	return result, err
}

