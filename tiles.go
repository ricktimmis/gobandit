package main

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

//Tile provides a data structure to hold image sets and values, and meets the requirements
//for the tileset interface. Tile loads a data model meta information from the configuration, and expects
//image assets to be provided as png files.
type Tile struct {
	imgpath  string         //path to png image assets
	face     map[int]string // png image files
	value    map[int]int    // value of each face
	currface int
	lastface int
	nextface int
	settotal int
	index    map[int]int // map pointer to keep faces and values in sync

}

// Constructor / Initialiser
func (t *Tile) init() {
	t.imgpath = ""
	t.face = make(map[int]string)
	t.value = make(map[int]int)
	t.currface = 1 // Keeping initial values simple
	t.lastface = 0 // so we can use a simle inc to
	t.nextface = 2 // to move indexes
	t.settotal = 0
	t.index = make(map[int]int, 10)
	return
}

// Loads data into the tile struct from passed config key value map
func (t *Tile) Load() error {
	t.imgpath = viper.GetString("FilePath")
	files, err := ioutil.ReadDir(t.imgpath)
	if err != nil {
		return err
	}
	//Load all PNG image files from the directory
	count := 0
	for _, f := range files {
		if strings.Contains(f.Name(), ".png") {
			t.face[count] = f.Name()
			count++
			//filenames should be name_value.png
			s := strings.Split(f.Name(), "_")
			v := strings.Split(s[1], ".")
			t.value[count], err = strconv.Atoi(v[0])
			if err != nil {
				return err
			}
		}

	}
	t.settotal = count
	return nil
}

// Returns a count of tiles in the set
func (t *Tile) Count() int {
	return t.settotal
}

// Randomly reorders the set
func (t *Tile) Shuffle() {

	rand.Shuffle(len(t.face), func(i, j int) {
		t.face[i], t.face[j] = t.face[j], t.face[i]
		t.value[i], t.value[j] = t.value[j], t.value[i]
	})
	return
}

// Pop next tile from the set, e,g like dealing a card
func (t *Tile) Next() {
	if t.lastface != t.settotal {
		t.lastface++
	}else {
		t.lastface = 0
	}
	t.currface++
	t.nextface++
	if t.currface == t.settotal {
		t.nextface = 0
	}
	return
}

// Returns the name of the current tile image
func (t *Tile) GetFace() string {
	//filenames should be name_value.png
	n := strings.Split(t.face[t.currface], "_")
	return n[0]
}

// Returns the value of the current tileface
func (t *Tile) GetValue() int {
	return t.value[t.currface]
}
// GetImage returns path to the image on disc
func (t *Tile) GetImage() string {
	// FIXME Return pointer to the image as a []byte in memory for the current tileface
	// Doing the above will remove the need for disk I/O
	return t.imgpath+t.face[t.currface]
}
