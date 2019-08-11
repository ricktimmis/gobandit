package main

import (
	"github.com/veandco/go-sdl2/sdl"

)
func main() {

	// CONFIGURATION CONSTANTS
	// Window dimensions
	width  := int32(900)
	height := int32(600)

	/* BACKGROUND IMAGE
	FIXME This should be replaced with a folder traversing function that can load multiple backgrounds
	*/
	bgimage :="./resources/backgrounds/dragon_of_the_north.jpg"


	// BASE INTERFACES
	// Board
	/* A Board provides an interface to a play board type
	more here later
	 */

	/* Tileset
	A Tileset provides the faces and values that will be displayed in the rows and columns of a Board.
	This abstraction enables tilesets to be created as plugable components to the GoBandit game. Thus
	a Fruit machine game might hold a tileset of 6 fruits, or a playing card game would hold a
	tileset of 52 playing cards.
	 */
	type tileset interface {
		Count() 	// Returns a count of tiles in the set
		Shuffle() 	// Randomly reorders the set
		Next()		// Pop next tile from the set, e,g like dealing a card
		GetFace()	// Returns the tile face description as a string
		GetValue()
		GetImage()
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	/*
	Remember Window->Surface->Renderer
	SDL Creates a Window and within in it is a Surface. We get a handle to the Surface using GetSurface
	Create a renderer for that surface. Use that Renderer to build the scene, and then update the Surface
	simples right ?
	 */

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	// Redundant from example code
	// surface.FillRect(nil, 0)

	//rect := sdl.Rect{0, 0, 200, 200}

	//surface.FillRect(&rect, 0xffff0000)
	//window.UpdateSurface()

	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	drawBackground(renderer, bgimage)
	running := true
	for running {
		window.UpdateSurface() //<- I think we need to do Clear, and Present instead
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}

