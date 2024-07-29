package tiles

// A TileID uniquely identifies a Tile on a factory floor.
type TileID int

const (
	// An Empty Tile has no contents, and may be moved into.
	Empty TileID = 0
	// A Barrier Tile is immobile and has no properties.
	Barrier TileID = 1

	Iron TileID = 2

	
)
