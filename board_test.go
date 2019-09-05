package main

import (
	"github.com/spf13/viper"
	"testing"
)

// Basic board method tests
// Initialise a 3 x 3 board, expects 3 columns and 3 rows
func TestBoard(t *testing.T) {
	// Setup code
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	b := new(Board)
	ti := new(Tile)
	t.Run("Tile Initialisation", func(t *testing.T) {
		b.Init(3, 3, ti)
		if b.Cols != 3 {
			t.Errorf("Expected board to have 3 columns, got %v", b.Cols)
		}
		if b.Rows != 3 {
			t.Errorf("Expected board to have 3 colums, got %v ", b.Rows)
		}
	})

	// Sets up a 4 x 4 board and expects a slice of 16 tiles to be returned
	t.Run("Test Rendering", func(t *testing.T) {
		//b := new(Board)
		//ti := new(Tile)
		b.Init(4, 4, ti)
		for r := b.Rows; r < b.Rows; r++ {
			for c := b.Cols; c < b.Cols; c++ {
				//test for Face and Value
				if b.Tiles[r][c].face == nil {
					t.Errorf("Expected tile faces")
				}
				if b.Tiles[r][c].value == nil {
					t.Errorf("Expected tile faces")
				}
			}
		}
	})

	t.Run("Test Play", func(t *testing.T) {
		//b := new(Board)
		//ti := new(Tile)
		b.Init(4, 4, ti)

	})

	t.Run("Test Score", func(t *testing.T){
		b.ScoreAdd(100)
		if b.GetScore() != 100{
			t.Errorf("Expected score of 100, got %v", b.GetScore())
		}
		b.ScoreAdd(50)
		if b.GetScore() !=150{
			t.Errorf("Expected score increase to 150, got %v", b.GetScore())
		}
	})
}

// FIXME Work out whether the board holds a winning combo
// func TestBoard_Eval(t *testing.T){}


