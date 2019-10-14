package main

import (
	"github.com/spf13/viper"
	"strings"
	"testing"
)

// TestTile provides a single function wrapper for a suite of subtests that require a tile instance.
// approaching it in this way enables us to work with a single Tile instance, and provides the ability
// for us to create interactive tests for behaviour later on.
// See https://godoc.org/testing
func TestTile(t *testing.T) {
	// <setup code>
	viper.SetDefault("FilePath", "./resources/tile_images/Fruits/")
	ti := new(Tile)
	//tile.Init()
	tile := ti.GetTile()
	//if err != nil {
	//	t.Errorf("loading configuration failed with: %v", err)
	//}

	// Tests
	t.Run("Tile Load", func(t *testing.T) {
		if len(tile.face) < 9 {
			t.Errorf("could not parse 10 png files in %v", tile.imgpath)
		}
		for _, v := range tile.face {
			if !(strings.Contains(v, ".png")) {
				t.Errorf("non PNG file loaded %v", v)
			}
		}
	})
	t.Run("Tile Count", func(t *testing.T) {
		if tile.Count() > 1 {
			// Provided there is at least 1 tile in the set PASS
			return
		}
		t.Errorf("Count of %v is too low", tile.Count())
	})

	t.Run("Tile Shuffle", func(t *testing.T) {
		// FIXME A bit messy Rick, clean this up
		t1 := tile.GetFace()
		tile.Shuffle()
		t2 := tile.GetFace()
		if t1 == t2 {
			// Could be a coincidence, that the shuffle put these back in situ
			tile.Shuffle()
			t2 := tile.GetFace()
			if t1 == t2 {
				t.Fail()
			}
		}
		return
	})

	t.Run("Tile Next", func(t *testing.T) {
		// FIXME A bit messy Rick, clean this up
		t1 := tile.GetFace()
		tile.Next()
		t2 := tile.GetFace()
		if t1 == t2 {
			t.Fail()
		}
		return
	})
	t.Run("Tile GetImage", func(t *testing.T) {
		s := tile.GetImage()
		if strings.Contains(s, ".jpg") {
			return
		}
		if strings.Contains(s, ".png") {
			return
		}
		t.Fail()
		return
	})
	t.Run("Tile GetValue", func(t *testing.T) {
		i := tile.GetValue()
		// Expect a non-zero value
		if i > 0 {
			return
		}
		t.Fail()
	})
	// <tear-down code>
}
