package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type control struct {
	debounce int    // We use debounce to filter out multiple queued SDL_events e.g Mouse Clicks, or KeyDown
	ready    bool   // Defines if GO button is lit or not
	board    *Board //FIXME - Switch this out for an interface
	score    *Score //FIXME - Switch this out for an interface
	soundfx  *sound //FIXME - Switch this out for an interface
	window   *sdl.Window
	font     *ttf.Font
	renderer *sdl.Renderer // Initialised here
	surface  *sdl.Surface  // Initialised here
}

type controller interface {
	spin()
	hold()
	nudge()
	init()
}

func (c *control) init() {
	var err error
	// Instantiate the other window elements
	c.surface, err = c.window.GetSurface()
	if err != nil {
		panic(err)
	}

	c.renderer, err = sdl.CreateSoftwareRenderer(c.surface)
	if err != nil {
		panic(err)
	}
	/* BACKGROUND IMAGE
	FIXME This should be replaced with a folder traversing function that can load multiple backgrounds
	      once complete, then drawbackground can be refactored into drawboard too.
	*/
	bgimage := "./resources/backgrounds/dragon_of_the_north.jpg"
	// Check file exists
	_, err = os.Open(bgimage)
	if err != nil {
		panic(err)
	}
	c.drawbackground(bgimage)
	c.ready = true
	c.drawboard()

	return
}

// All the interesting functions are here !
// Implementation functions that do the work and the rendering
// This is the stuff you want to mess with to get different game behaviour

func (c *control) spin() error {
	//err = renderer.Clear()
	//if err != nil {
	//	panic(err)
	//}
	if c.ready {
		c.spinimation()
	}
	return nil
}

func (c *control) hold() {
	return
}

func (c *control) nudge() {
	return
}

// FIXME - Requires a file path to the BackgroundMusic image
func (c *control) drawbackground(i string) error {
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
	var dst = sdl.Rect{0, 0, 900, 600}

	texture, err := img.LoadTexture(c.renderer, i)
	if err != nil {
		return err
	}
	err = c.renderer.Copy(texture, nil, &dst)
	if err != nil {
		return err
	}
	return err
}

func (c *control) drawboard() error {
	// Work out positions sizes based upon Window Size and number
	// of Columns and Rows in the board, iterate through the board applying face
	// textures
	// FIXME Refactor this int/int32 recasting, which will require board Cols and Rows to be int32,
	colwidth := int32(int(width) / (c.board.Cols + 2))
	rowheight := int32(int(height) / (c.board.Rows + 2))

	x := colwidth
	y := rowheight
	for row := 0; row < (c.board.Rows); row++ {
		for col := 0; col < (c.board.Cols); col++ {
			var dst = sdl.Rect{x, y, colwidth, rowheight}
			p := c.board.Tiles[row][col].GetImage()
			texture, err := img.LoadTexture(c.renderer, p)
			if err != nil {
				return fmt.Errorf("could not load BackgroundMusic image : %v", err)
			}
			err = c.renderer.Copy(texture, nil, &dst)
			if err != nil {
				return fmt.Errorf("could not render BackgroundMusic image : %v", err)
			}
			x = x + colwidth
		}
		// Reset back to first column
		x = colwidth
		y = y + rowheight
	}
	/* Draw control buttons
	 */
	for col := 0; col < (c.board.Cols); col++ {
		var dst = sdl.Rect{(x + 10), y, (colwidth - 20), rowheight}
		// FIXME - Implement a range of buttons, including hold, nudge etc.
		//         Need to load these through a function so that we can involve
		//         logic here to decide if we are offering hold, nudge, or are
		//         holding, or nudging
		button := "./resources/controller_assets/nudge_button.png"
		texture, err := img.LoadTexture(c.renderer, button)
		if err != nil {
			return fmt.Errorf("could not load BackgroundMusic image : %v", err)
		}
		err = c.renderer.Copy(texture, nil, &dst)
		if err != nil {
			return fmt.Errorf("could not render BackgroundMusic image : %v", err)
		}
		x = x + colwidth

	}
	/* Draw Go button
	   This checks the controllers ready state and draws one of two buttons
	   FIXME - This is a horrible, hardcoding shortcut to get the code into release, as it
	           was required for a time bound magazine article...
	*/
	var dst = sdl.Rect{(width - 125), y, 80, 80}
	var button = ""
	if c.ready {
		button = "./resources/controller_assets/go_button_rdy.png"
	} else {
		button = "./resources/controller_assets/go_button.png"
	}
	texture, err := img.LoadTexture(c.renderer, button)
	if err != nil {
		return fmt.Errorf("could not load BackgroundMusic image : %v", err)
	}
	err = c.renderer.Copy(texture, nil, &dst)
	if err != nil {
		return fmt.Errorf("could not render BackgroundMusic image : %v", err)
	}
	/* Draw Score Panel
	Avoiding a fixed Score rectangle by calculating the Score window as a percentage of the window.
	20% (1/5) of width and height
	*/
	wide := width / 3
	high := height / 10
	x = (width / 2) - (wide / 2)
	y = (height / 25)
	var scoreborder = sdl.Rect{x, y, wide, high}
	var scorepanel = sdl.Rect{(x + 5), (y + 5), (wide - 10), (high - 10)}
	c.renderer.SetDrawColor(255, 230, 15, 255)
	c.renderer.FillRect(&scoreborder)
	c.renderer.SetDrawColor(58, 58, 58, 255)
	c.renderer.FillRect(&scorepanel)

	// Score

	scorestring := strconv.Itoa(int(c.board.score))

	textcolor := sdl.Color{255, 230, 15, 20}
	scoretextsurface, err := c.font.RenderUTF8Solid(scorestring, textcolor)
	if err != nil {
		return err
	}
	scoretexture, err := c.renderer.CreateTextureFromSurface(scoretextsurface)
	if err != nil {
		return err
	}
	if c.renderer.Copy(scoretexture, nil, &scorepanel); err != nil {
		fmt.Errorf("Failed to render Score texture")
	}
	//if r.FillRect(&scorepanel); err !=nil {
	//	fmt.Errorf("Failed to fill Score panel")
	//}

	//RenderUTF8Solid(text string, color sdl.Color) (*sdl.Surface, error)

	// Draw the controls

	return nil
}

func (ctrl *control) playnext(col int) error {

	for c := col; c < ctrl.board.Cols; c++ {
		for r := 0; r < ctrl.board.Rows; r++ {
			ctrl.board.Tiles[r][c].Next()
		}
	}
	return nil
}

func (c *control) checkscore() (int, error) {

	// Define rules here to evaluate board and generate a Score

	// Rule 1 - Check each row for adjacent matching tile
	rule1 := func(b *Board) int {
		sum := 0
		for r := 0; r != b.Rows; r++ {
			for c := 0; c < (b.Cols - 1); c++ {
				if (b.Tiles[r][c].GetFace()) == (b.Tiles[r][c+1].GetFace()) {
					sum = sum + b.Tiles[r][c].GetValue()
					fmt.Printf("Row %d Columns %d and %d Matching %s \n", r, c, c+1, (b.Tiles[r][c].GetFace()))

					// FIXME Call or setup a display routine, to highlight the match
				}
			}
		}
		return sum
	}
	v, err := c.score.evaluate(rule1)
	return v, err
}

func (c *control) spinimation() error {

	// Start music - see sound.go
	//sounds := sound{}
	//sounds.Init()
	c.soundfx.Playspinsoundfx()

	rand.Seed(time.Now().UnixNano())
	min := 50
	max := 400
	iterations := (rand.Intn(max-min+1) + min)
	// I noticed that for some spins there appeared to be a value less than 100, perhaps due to behaviour of
	// rand. So this catches that, and fixes it.
	if iterations < 50 {
		iterations = 50
	}

	/*
		columnstops work by making each column from left to right appear to stop before the other
		by dividing the number of iterations for this spin by the columns on the board we calculate a value at which
		we increment the col parameter to be passed into play next().
	*/
	columnstops := iterations / c.board.Cols
	iterationspercolumn := columnstops
	columnlocked := 0

	for i := 0; i < iterations; i++ {
		if i > columnstops {
			columnlocked++
			columnstops = columnstops + iterationspercolumn
			c.soundfx.Playcolumnstopfx()
		}
		c.ready = false // Defines if GO button is lit or not
		c.playnext(columnlocked)
		sdl.Delay(10)
		c.drawboard()
		err := c.window.UpdateSurface()
		if err != nil {
			return err
		}
	}
	c.soundfx.Playcolumnstopfx()
	c.ready = true
	v, err := c.checkscore()
	if err != nil {
		return err
	}
	if v > 0 {
		fmt.Printf("Match found and scored %d \n", v)
	}
	c.board.ScoreAdd(v)
	c.drawboard()
	if c.window.UpdateSurface(); err != nil {
		return err
	}
	return nil
}
