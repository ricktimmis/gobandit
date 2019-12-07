package main

import "testing"

func TestSound_Init(t *testing.T) {
	var err error
	testsounder := sound{}
	t.Run("Initialise", func(t *testing.T) {
		if testsounder.Init(); err != nil {
			t.Errorf("Unexpected error when initialising sound %v", err)
		}
		if !testsounder.initialised {
			t.Errorf("Sound failed to initialise, sound.initialised = false")
		}
	})
	t.Run("Play Spin Sound FX", func(t *testing.T) {
		chnl, err := testsounder.Playspinsoundfx()
		if err != nil {
			t.Errorf("Play Spin Sound return an error: %v", err)
		}
		if chnl == 0 {
			t.Errorf("Play Spin Sound failed to allocate a valid play channel")
		}

	})
	t.Run("Play Column Stop FX", func(t *testing.T) {
		chnl, err := testsounder.Playcolumnstopfx()
		if err != nil {
			t.Errorf("Play Column Stop Sound return an error: %v", err)
		}
		if chnl == 0 {
			t.Errorf("Play Columns Stop Sound failed to allocate a valid play channel")
		}
	})
	t.Run("Play Background Music", func(t *testing.T) {
		err := testsounder.BackgroundMusic()
		if err != nil {
			t.Errorf("Play Background Music return an error: %v", err)
		}
	})
}
