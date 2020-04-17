package main

import (
	"math"
)

type Cache struct {
	bitsPerWord   int
	size          int
	bytesPerBlock int
	linesPerBlock int
	totalBlocks   int
	blockOffset   int
	lineOffset    int
	blocks        []*Block
}

func NewCache(bitsPerWord int, size int, bytesPerBlock int, linesPerBlock int) *Cache {
	c := &Cache{
		bitsPerWord:   bitsPerWord,
		size:          size,
		bytesPerBlock: bytesPerBlock,
		linesPerBlock: linesPerBlock,
		totalBlocks:   (size / bytesPerBlock),
		blockOffset:   int(math.Log2(float64(size / bytesPerBlock))),
		lineOffset:    int(math.Log2(float64(bytesPerBlock / linesPerBlock))),
	}
	c.blocks = make([]*Block, c.totalBlocks)
	for i := 0; i < c.totalBlocks; i++ {
		c.blocks[i] = NewBlock(c, i)
	}
	return c
}

func (c *Cache) access(a *Address) bool {
	return c.blocks[a.index].access(a)
}
