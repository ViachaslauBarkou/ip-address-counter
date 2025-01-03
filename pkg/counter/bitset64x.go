package counter

import (
	"fmt"
	"math/bits"
)

type BitSetN struct {
	array   []uint64
	bitSize int
	parts   int
}

func NewBitSetN(bitSize int) (*BitSetN, error) {
	if bitSize != 32 && (bitSize%64 != 0 || bitSize < 64) {
		return nil, fmt.Errorf("wrong bitSet size: %d. Must be a multiple of 64", bitSize)
	}
	parts := bitSize / 64
	arraySize := 1 << (32 - bits.TrailingZeros(uint(bitSize)))
	return &BitSetN{
		array:   make([]uint64, arraySize*parts),
		bitSize: bitSize,
		parts:   parts,
	}, nil
}

func (b *BitSetN) SetBit(index int) {
	wordIndex := index / b.bitSize
	bitIndex := index % b.bitSize
	partIndex := bitIndex / 64
	bitOffset := bitIndex % 64
	b.array[wordIndex*b.parts+partIndex] |= 1 << bitOffset
}

func (b *BitSetN) IsBitSet(index int) bool {
	wordIndex := index / b.bitSize
	bitIndex := index % b.bitSize
	partIndex := bitIndex / 64
	bitOffset := bitIndex % 64
	return (b.array[wordIndex*b.parts+partIndex] & (1 << bitOffset)) != 0
}
