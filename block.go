package main

import (
	"sort"
	"time"
)

type Block struct {
	cache *Cache
	index int
	lines []*Line
}

func NewBlock(c *Cache, index int) *Block {
	return &Block{
		cache: c,
		index: index,
		lines: make([]*Line, c.linesPerBlock),
	}
}

func (b *Block) access(a *Address) bool {
	for _, l := range b.lines {
		if l == nil {
			continue
		}
		if l.ContainsAddress(a) {
			l.lastAccessTime = time.Now()
			sort.Slice(b.lines, func(i, j int) bool {
				lineI := b.lines[i]
				if lineI == nil {
					return false
				}
				lineJ := b.lines[j]
				if lineJ == nil {
					return false
				}
				return lineI.lastAccessTime.Before(lineJ.lastAccessTime)
			})
			return true
		}
	}
	b.lines = append(b.lines[1:], NewLine(b, a))
	return false
}
