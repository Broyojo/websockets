package voxel

// Block type alias
type Block int

// Block Enum
const (
	Air Block = iota
	Grass
	Dirt
	Stone
	NumBlocks int = iota
)
