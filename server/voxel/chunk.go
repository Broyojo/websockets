package voxel

import (
	"math/rand"
)

// ChunkSize - length of chunk on 3 axes
const ChunkSize = 16

// Chunk - holds block and position data for a chunk
type Chunk struct {
	X, Y, Z int
	Blocks  [ChunkSize][ChunkSize][ChunkSize]Block
}

// NewChunk - makes a new chunk
func NewChunk(x, y, z int) *Chunk {
	var c Chunk
	c.X, c.Y, c.Z = x, y, z
	return &c
}

// Map - maps a function over the chunk
func (c *Chunk) Map(f func(x, y, z int) Block) {
	for x := 0; x < ChunkSize; x++ {
		for y := 0; y < ChunkSize; y++ {
			for z := 0; z < ChunkSize; z++ {
				c.Blocks[x][y][z] = f(x, y, z)
			}
		}
	}
}

// Randomize - randomizes a chunk
func (c *Chunk) Randomize() {
	c.Map(func(x, y, z int) Block {
		return Block(rand.Intn(NumBlocks))
	})
}
