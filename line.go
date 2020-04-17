package main

import (
	"time"
)

type Line struct {
	block          *Block
	addr           *Address
	lastAccessTime time.Time
}

func NewLine(b *Block, a *Address) *Line {
	return &Line{
		block:          b,
		addr:           a,
		lastAccessTime: time.Now(),
	}
}

func (l *Line) ContainsAddress(a *Address) bool {
	return l.addr.tag == a.tag
}
