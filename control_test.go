package main

import (
	"github.com/spf13/viper"
	"testing"
)

func TestController(t *testing.T) {
	// <setup code>
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	b := new(Board)
	f := new(Tile)
	s := new(Score)
	c := new(control)
	f.init()
	f.load()
	// here we initialise the board which will generate some data
	b.Init(4, 4, f)
	// Pass a pointer for the board to the Scorer
	s.init(b)

	t.Run("Control Spin", func(t *testing.T) {
		rule := func(b *Board) int {
			// For this test evaluation we know we have a number of tiles, with a minimum value of 10 per tile
			// our evaluation will be to sum the initial board.
			// Then spin the board, and recalculate the sum, and compare the two.
			// It is unlikely that the two sums would be equal after a board spin1
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
		v1, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		c.spin()
		v2, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		if v1 == v2 {
			t.Errorf("Spin() failed to change board values")
		}
	})
	t.Run(" Control Hold", func(t *testing.T) {
		// We can test behaviour again here. This time we Spin the board take a sum
		// Hold all columns and Spin again
		// Take a second sum and compare them. They should be the same if all columns were held.
		rule := func(b *Board) int {
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
		// First spin
		c.spin()
		v1, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		// Hold all columns
		c.hold()
		// Second spin
		c.spin()
		v2, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		if v1 != v2 {
			t.Errorf("Hold() failed to retain columns change board values")
		}
	})
	t.Run(" Control Nudge", func(t *testing.T) {
		// This time we Spin the board take a sum, then read ahead and sum the next face values in the tileset ;-)
		// Nudge the columns and Sum again
		// Compare with the read ahead sum. They should be the same if all columns were nudged 1 place.
		// FIXME Issue #5 Nudge Feature
		rule := func(b *Board) int {
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
		// First spin
		c.spin()
		v1, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		// Hold all columns
		c.hold()
		// Second spin
		c.spin()
		v2, err := s.evaluate(rule)
		if err != nil {
			t.Errorf("evaluation failed with error %v", err)
		}
		if v1 != v2 {
			t.Errorf("Hold() failed to retain columns change board values")
		}
	})
}
