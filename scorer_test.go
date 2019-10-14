package main

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestScorer(t *testing.T) {
	// <setup code>
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	b := new(Board)
	f := new(Tile)
	s := new(Score)
	f.init()
	f.load()
	// here we initialise the board which will generate some data
	b.Init(4, 4, f)
	// Pass a pointer for the board to the Scorer
	s.init(b)

	// I don't have the answer to this yet, I was thinking about injecting a function with embedded
	// logic. The receiving method executes the logic on the board tile slice.
	// At this stage the function injection works
	// FIXME Work out the logic to iterate through the board tileset and retrieve a Score value. Will need to fix the Mock data above too

	t.Run("Score Evaluation", func(t *testing.T) {
		rule := func(b *Board) int {
			// For this test evaluation we know we have 12 tiles, with a minimum value of 10 per tile
			// our evaluation will be a simple additional of all the tile values on the board
			// Remember that the each tile on the board is a pointer reference to a tile instance
			// each tile instance has multiple faces, and values and that we need to get an index value for
			// the current face, to then be able to look up its value from the possible face/values in the set
			sum := 0
			for r := 0; r != b.Rows; r++ {
				for c := 0; c != b.Cols; c++ {
					i := b.Tiles[r][c].currface
					v := b.Tiles[r][c].value[i]
					sum = sum + v
				}
			}
			return sum
		}
		v, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		if v == 0 {
			t.Errorf("Expected zero, got %v", v)
		}
		fmt.Printf("Board value in this test was %v", v)
	})

}
