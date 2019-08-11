package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)
// Requires a file path to the background image
func drawBackground(r *sdl.Renderer, i string) error  {
	t, err := img.LoadTexture(r, "./resources/backgrounds/dragon_of_the_north.jpg")
	if err != nil {
		return fmt.Errorf("Could not load bakground image : %v", err)
	}
	r.Copy(t, nil,nil)
	return nil
}

//func drawScene(r *sdl.Renderer) error {

//}