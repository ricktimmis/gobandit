package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Requires a file path to the background image
func drawBackground(r *sdl.Renderer, i string) error  {
	// Initialise pointer ;-)
	var texture *sdl.Texture
	/* Fixes issue #2 Crashes with segev null pointer Kubuntu 18.04
	https://github.com/ricktimmis/gobandit/issues/2

	Turns out that although SDL2 provides default src:nil dst:nil in renderer.Copy
	on some systems the c.go lowlevel binding to C blows up. Probably because
	src and dst use the unsafe.Pointer package. Declaring a rectangle layer the same
	size as the window, and initialising it gives us a definite destination for the Copy.
	3 hours it took me to debug that!!
	 */
	var dst = sdl.Rect{0,0,900,600}

	texture, err := img.LoadTexture(r, i)
	if err != nil {
		return fmt.Errorf("could not load background image : %v", err)
	}
	err = r.Copy(texture, nil, &dst)
	if err != nil {
		return fmt.Errorf("could not render background image : %v", err)
	}
	return err
}

func drawBoard(r *sdl.Renderer, b *Board) error {
	// Work out positions sizes based upon Window Size and number
	// of Columns and Rows in the board, iterate through the board applying face
	// textures
	// FIXME Refactor this int/int32 recasting, which will require board Cols and Rows to be int32,
	colwidth := int32(int(width) / (b.Cols+2))
	rowheight := int32(int(height) / (b.Rows+2))

	x := colwidth
	y := rowheight
	for row := 0; row < (b.Rows); row++{
		for col := 0; col < (b.Cols); col++ {
			var dst= sdl.Rect{x, y, colwidth, rowheight}
			p := b.Tiles[row][col].GetImage()
			texture, err := img.LoadTexture(r, p)
			if err != nil {
				return fmt.Errorf("could not load background image : %v", err)
			}
			err = r.Copy(texture, nil, &dst)
			if err != nil {
				return fmt.Errorf("could not render background image : %v", err)
			}
			x = x + colwidth
		}
		// Reset back to first column
		x = colwidth
	y = y + rowheight
	}

	/* Draw Score Panel
	Avoiding a fixed score rectangle by calculating the score window as a percentage of the window.
	20% (1/5) of width and height
	 */
	wide := width / 3
	high := height / 10
	x = (width / 2) - (wide / 2)
	y = (height / 25)
	var scoreborder = sdl.Rect{x, y, wide, high}
	var scorepanel = sdl.Rect{(x + 5), (y + 5 ), (wide - 10), (high - 10)}
	r.SetDrawColor(255,230,15,255)
	r.FillRect(&scoreborder)
	r.SetDrawColor(58,58,58,255)
	r.FillRect(&scorepanel)

	return nil
}

func playNext(r *sdl.Renderer, b *Board) {

	for c := 0; c < b.Cols; c++{
		for r := 0; r < b.Rows; r++{
			b.Tiles[r][c].Next()
		}
	}
	return
}

func checkScore(s *score,) (int, error){


	// Define rules here to evaluate board and generate a score

	// Rule 1 - Check each row for adjacent matching tile
	rule1 := func(b *Board)int{
		sum := 0
		for r:=0; r != b.Rows; r++{
			for c:=0; c < (b.Cols - 1); c++{
				if (b.Tiles[r][c].GetFace()) == (b.Tiles[r][c+1].GetFace()){
					sum = sum + b.Tiles[r][c].GetValue()
					fmt.Printf("Row %d Columns %d and %d Matching %s \n", r, c, c+1, (b.Tiles[r][c].GetFace()))

					// FIXME Call or setup a display routine, to highlight the match
				}
			}
		}
		return sum
	}
	v, err := s.evaluate(rule1)
	return v, err
}