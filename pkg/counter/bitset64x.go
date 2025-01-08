package counter

import (
	"fmt"
	"math/bits"
	"sync/atomic"
)

type BitSetN struct {
	array   []uint64
	bitSize int
	parts   int
}

func NewBitSetN(bitSize int) (*BitSetN, error) {
	if bitSize%64 != 0 || bitSize < 64 {
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
	wordIndex, bitIndex := calculateIndexes(index, b.bitSize)
	partIndex, bitOffset := calculatePartIndexes(bitIndex, 64)
	b.array[wordIndex*b.parts+partIndex] |= 1 << bitOffset
}

func (b *BitSetN) AtomicSetBit(index int) bool {
	wordIndex, bitIndex := calculateIndexes(index, b.bitSize)
	partIndex, bitOffset := calculatePartIndexes(bitIndex, 64)

	addr := &b.array[wordIndex*b.parts+partIndex]
	for {
		oldValue := atomic.LoadUint64(addr)
		if oldValue&uint64(1<<bitOffset) != 0 {
			return true
		}
		newValue := oldValue | uint64(1<<bitOffset)
		if atomic.CompareAndSwapUint64(addr, oldValue, newValue) {
			return false
		}
	}
}

func (b *BitSetN) IsBitSet(index int) bool {
	wordIndex, bitIndex := calculateIndexes(index, b.bitSize)
	partIndex, bitOffset := calculatePartIndexes(bitIndex, 64)
	return (b.array[wordIndex*b.parts+partIndex] & (1 << bitOffset)) != 0
}
