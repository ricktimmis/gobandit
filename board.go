package main


type Board struct {
	Cols	int
	Rows	int
	Tiles   [8][8]*Tile // For now we'll limit the board to an 8 x 8 grid max
	score	uint32 // We don't export this, but use Getter Setter instead

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
// Setter for adding int values to the score
func (b *Board) ScoreAdd(v int) error{
	err := error(nil)
	return err
}
// Getter for retrieving current score
func (b *Board) GetScore() uint32{
	return 0
}
