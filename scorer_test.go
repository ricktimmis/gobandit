package main

import (
	"github.com/spf13/viper"
	"testing"
)

func TestScorer(t *testing.T) {
	// <setup code>
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	b := new(Board)
	f := new(Tile)
	s := new(score)
	f.init()
	f.load()
	// here we load the board with some mock date
	b.Init(4, 4, f)

	t.Run("Score Evaluation", func(t *testing.T) {
		v := s.Evaluate(b)
		if v == 0 {
			t.Errorf("Expected non zero, got %v", v)
		}
	})
	// I don't have the answer to this yet, I was thinking about injecting a function with embedded
	// logic. The receiving method executes the logic on the board tile slice.
	// At this stage the function injection works
	// FIXME Work out the logic to iterate through the board tileset and retrieve a score value. Will need to fix the Mock data above too
	t.Run("Score Load Ruleset", func(t *testing.T) {
		rule := func(b *Board) int {
			return (1+1)
		}
		v := s.LoadRules(rule)
		if v != 0 {
			t.Errorf("Expected zero, got %v", v)
		}
	})
	// Tests

}