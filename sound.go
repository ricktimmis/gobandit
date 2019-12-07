package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

// Initialise Audio and sound effects
/*
	Use these constants to enable format support/tmp
	https://www.libsdl.org/projects/SDL_mixer/docs/SDL_mixer_9.html
	This cheatsheet will likely help to
	https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
*/
//const (
//	MIX_INIT_FLAC 	= 1
//	MIX_INIT_MOD	= 1
//	MIX_INIT_MP3	= 1
//	MIX_INIT_OGG	= 1
//)
//flags := mix.INIT_MOD
//if err := mix.Init(flags); err != nil{
//	fmt.Printf("Audio initialisation failed with error: %v \n", err)
//}
/* Opening the audio for output.
To Grok the Output Format I had to do a little research to understand it(ish)
To follow this you'll need the following link.
https://github.com/emscripten-ports/SDL2/blob/master/include/SDL_audio.h
http://soundfile.sapp.org/doc/WaveFormat/
Do read the GoDocs, and SDL reference before the links above
*/

/*
Trying to use a singleton pattern here, so that we only pass one sound instance around
FIXME Needs documenting
*/
type sound struct {
	initialised bool
	spinfx      *mix.Chunk
	colstop     *mix.Chunk
}

func (s *sound) Init() (instance *sound, err error) {

	if !s.initialised {
		if err := mix.OpenAudio(44100, sdl.AUDIO_U8, 2, 4096); err != nil {
			//if err := mix.OpenAudio(22050,16,2,4096); err != nil{
			err = fmt.Errorf("Audio failed to open for Output with error: %v \n", err)
			return nil, err
		}
		s.spinfx, _ = mix.LoadWAV(viper.GetString("spinsoundfx"))
		s.colstop, _ = mix.LoadWAV(viper.GetString("columnstopfx"))
		s.initialised = true
	}

	return instance, nil
}

func (s *sound) BackgroundMusic() (err error) {

	backgrndmusic, err := mix.LoadMUS(viper.GetString("backgrndmusic"))
	if err != nil {
		fmt.Printf("Audio failed loading BackgroundMusic music with error: %v \n", err)
		return err
	}
	// Let music play
	backgrndmusic.Play(-1)
	return nil
}

func (s *sound) Playspinsoundfx() (chnl int, err error) {
	chnl, err = s.spinfx.Play(-1, 0)
	if err != nil {
		fmt.Errorf("Play failed with error: %v", err)
		return chnl, err
	}
	return chnl, err

}
func (s *sound) Playcolumnstopfx() (chnl int, err error) {
	chnl, err = s.colstop.Play(-1, 0)
	if err != nil {
		fmt.Errorf("Play failed with error: %v", err)
		return chnl, err
	}
	return chnl, err
}
