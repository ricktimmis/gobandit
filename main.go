package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// CONFIGURATION GLOBALS & CONSTANTS
// Window dimensions

var width = int32(900)
var height = int32(600)

func main() {

	// FIXME Implement loadable configuration from a local .cfg file
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	viper.SetDefault("FontFile", "./resources/fonts/open-sans/OpenSans-Regular.ttf")
	viper.SetDefault("backgrndmusic", "./resources/sound/music/halloween.wav")
	viper.SetDefault("spinsoundfx", "./resources/sound/effects/producerspot-sfx-11.wav")
	viper.SetDefault("columnstopfx", "./resources/sound/effects/Bounce.wav")

	// SDL Pointer initialisation
	var window *sdl.Window
	var renderer *sdl.Renderer
	var surface *sdl.Surface
	var event sdl.Event

	// FIXME - I think this is redundant, but we need to fix up error management, and returns errors up the
	//		   the chain of functional calls - Thanks Dave Cheney
	var err error

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

	window, err = sdl.CreateWindow("Go Bandit", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Initialise a ttf font renderer
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont((viper.GetString("FontFile")), 320)
	if err != nil {
		panic(err)
	}

	// Board is a model of the game, it holds a tileset for each column
	var board = new(Board)
	// Initialise Score, which saves us from having to pass the board pointer each time we call
	// the checkScore function - see scene.go
	var score = new(Score)
	score.init(board)

	// Initialise the Board
	ti := new(Tile) // FIXME This is coupling to tiles.go, and it must not
	//board.Init(3, 4, ti)
	board.Init(5, 4, ti)

	// SoundFX
	var soundfx = new(sound)
	soundfx.Init()

	// Control is where the action is.... more to come here
	var controller = control{
		0,
		false,
		board,
		score,
		soundfx,
		window,
		font,
		renderer,
		surface,
	}
	// Initialise the controller
	controller.init()
	// Game Main loop
	running := true
	for running {
		window.UpdateSurface() //<- I think we need to do Clear, and Present instead
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				controller.debounce++
				if controller.debounce == 1 {
					fmt.Printf("KEY PRESSED -- Current Score = %d\n", board.score)
					controller.debounce++
					err = controller.spin()
					if err != nil {
						fmt.Errorf("Hmm something went wrong : %v", err)
					}
				}
			case *sdl.MouseButtonEvent:
				controller.debounce++
				if controller.debounce == 1 {
					fmt.Printf("MOUSE CLICK -- Current Score = %d\n", board.score)
					controller.debounce++
					err = controller.spin()
					if err != nil {
						fmt.Errorf("Hmm something went wrong : %v", err)
					}
				}
			}
			if controller.debounce > 0 {
				controller.debounce--
			}
		}
	}
}
