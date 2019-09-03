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
	// FIXME Hardcoded window width x height
	colwidth := int32((900 / (b.Cols+2)))
	rowheight := int32((600 / (b.Rows+2)))

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
