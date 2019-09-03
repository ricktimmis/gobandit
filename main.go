package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"os"
	"time"
)

func main() {

	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")

	// CONFIGURATION CONSTANTS
	// Window dimensions
	var width = int32(900)
	var height = int32(600)
	// Pointer initialisation
	var window *sdl.Window
	var renderer *sdl.Renderer
	var surface *sdl.Surface
	var event sdl.Event

	var err error
	var board = new(Board)

	/* BACKGROUND IMAGE
	FIXME This should be replaced with a folder traversing function that can load multiple backgrounds
	*/
	bgimage := "./resources/backgrounds/dragon_of_the_north.jpg"
	// Check file exists
	_, err = os.Open(bgimage)
	if err != nil {
		panic(err)
	}

	/* FIXME I am very tempted to refactor all of this initialisation into a Game struct
	Thus a Game would expect a Board, and a Board would expect a Tileset
	*/

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

	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	err = renderer.Clear()
	if err != nil {
		panic(err)
	}
	err = drawBackground(renderer, bgimage)
	if err != nil {
		panic(err)
	}
	// Initialise the Board
	ti := new(Tile)  // FIXME This is coupling to tiles.go, and it must not
	board.Init(3, 4, ti)
	running := true
	iterate := false
	doiteration := false
	iterations := 0
	start := 0
	for running {
		window.UpdateSurface() //<- I think we need to do Clear, and Present instead

		// FIXME This logic is messy and confusing
		// Why? - to avoid constantly calling into the SDL library to update the surface, which
		// caused the Window renderer to stop performing updates. I have used conditional logic
		// to define which calls are made when.
		if iterate == true && doiteration == false {
			rand.Seed(time.Now().UnixNano())
			min := 5
			max := 20
			iterations = (rand.Intn(max-min+1) + min)
			start = 0
			doiteration = true
		}
		if doiteration == true {
			playNext(renderer, board)
			drawBoard(renderer, board)
			sdl.Delay(10)
			start++
			if start > iterations {
				doiteration = false
				iterate = false
			}
		}
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				//fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
				//	event., event.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				fmt.Printf("KEY PRESSED\n")
				err = drawBoard(renderer, board)
				if err != nil {
					fmt.Errorf("Hmm something went wrong : %v", err)
				}
			case *sdl.MouseButtonEvent:
				fmt.Printf("MOUSE CLICK\n")
				iterate = true

			}
		}
	}
}
