# Go Bandit 

Go Bandit is a simple fruit machine game written using Go and SDL2. It is designed to be multilevel with objectives and
awards. The user collects gold coins, to reach a given goal and is rewarded with a new level of play.

# Architecture Synopsis

GoBandit consists of a Scene, Control, Score, Board and Tile. These are the implementation components each component implements an interface.

Ensure you take a look at GoBandit App Architecture.mm ( This is a Freeplane file)

## Scene
Scene is the visualisation for the current state of the game board. It expects to be passed a "Board" interface and a pointer to an SDL Surface.

## Control
Provides the user interface with controls such as "Play", "Nudge" and "Hold"
Control is an interface

## Scorer 
Provides a total count of points accumulated during the game, and provides
add, sub, get and reset methods

## Board
This is where the game logic is performed, as the board must be responsible
for configuring the pay area (i.e number of columns and rows) adjusting the tile
placeholder size accordingly. It must respond to Control calls, and Scene calls

# Tile
A Tile is a single set of faces and values which are used to populate a tileset, which
inturn is utilised by the Board 

# Sound and Audio Effects
The game should provide background music, as some form of loop along with a variety of
special effects. For example spining the reels, and pressing the Go button.
Audio should also be themed, (see Issue #.. Themes) such that the background loops and
sound effects can be different for each game pack.
## Design idea
My idea is to take a foray into Go Routines, enabling us to call the Sounds for playback
whilst still continuing on with game play