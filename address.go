package main

import (
	"strconv"
	"strings"
)

type Address struct {
	raw    string
	tag    uint64
	index  uint64
	offset uint64
}

func ParseAddress(raw string, c *Cache) (*Address, error) {
	decimal, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return nil, err
	}
	binary := strconv.FormatUint(uint64(decimal), 2)
	if err != nil {
		return nil, err
	}
	binaryPadded := strings.Repeat("0", c.bitsPerWord-len(binary)) + binary

	tagEnd := c.bitsPerWord - (c.blockOffset + c.lineOffset)
	tag, err := strconv.ParseUint(binaryPadded[:tagEnd], 2, 64)
	if err != nil {
		return nil, err
	}
	indexEnd := tagEnd + c.blockOffset
	var index uint64
	if c.blockOffset != 0 {
		index, err = strconv.ParseUint(binaryPadded[tagEnd:indexEnd], 2, 64)
		if err != nil {
			return nil, err
		}
	}
	var offset uint64
	offsetEnd := indexEnd + c.lineOffset
	if c.lineOffset != 0 {
		offset, err = strconv.ParseUint(binaryPadded[indexEnd:offsetEnd], 2, 64)
		if err != nil {
			return nil, err
		}
	}
	return &Address{
		raw:    raw,
		tag:    tag,
		index:  index,
		offset: offset,
	}, nil
}
