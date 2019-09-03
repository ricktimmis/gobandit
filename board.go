package main

/*
Tileset
A Tileset provides the faces and values that will be displayed in the rows and columns of a Board.
This abstraction enables tilesets to be created as plugable components to the GoBandit game. Thus
a Fruit machine game might hold a tileset of 6 fruits, or a playing card game would hold a
tileset of 52 playing cards.
*/
type TileSet interface {
	//Init()		bool 	// True for success
	//Load()		bool 	// True for success
	Count()		int  	// Number of tiles in the set
	Shuffle()
	Next()
	GetFace()	string	// Returns the face name of the current tile
	GetValue()	int		// Returns the value of the current tile
	GetImage()	string	// Returns a filename for the tile image
	GetTile()	*Tile	// Returns a pointer to the tile instance
}

type Board struct {
	Cols	int
	Rows	int
	Tiles   [8][8]*Tile // For now we'll limit the board to an 8 x 8 grid max

}
func (b * Board) Init(c int, r int, t TileSet) error{

	// Set Row, Col values on the Board for reference
	b.Cols = c
	b.Rows = r

	/* Creating an array element for each grid column, enables us to load each column with a TileSet
	   Enabling each column to iterate through it's TileSet. Shuffle each set
	 */
	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Cols; col++ {
			b.Tiles[row][col] = t.GetTile()
			//b.Tiles[row][col].Init()
			//b.Tiles[row][col].Load()
			b.Tiles[row][col].Shuffle()
		}
	}
	err := error(nil)
	return err
	
}
