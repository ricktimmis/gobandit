# Go Bandit 

Go Bandit is a simple fruit machine game written using Go and SDL2. It is designed to be multilevel with objectives and
awards. The user collects gold coins, to reach a given goal and is rewarded with a new level of play.

# Architecture Synopsis

GoBandit consists of a Scene, Control, Score, Board and Tile. 
## Scene
Scene is the visualisation for the current state of the game board. It expects to be passed a "Board" interface and a pointer to an SDL Surface.

## Control
Provides the user interface with controls such as "Play", "Nudge" and "Hold"
Control is an interface

## Score 
Provides a total count of points accumulated during the game, and provides
add, sub, get and reset methods

## Board
This is where the game logic is performed, as the board must be responsible
for configuring the pay area (i.e number of columns and rows) adjusting the tile
placeholder size accordingly. It must respond to Control calls, and Scene calls

# Tile
A Tile is a single set of faces and values which are used to populate a tileset, which
inturn is utilised by the Board 