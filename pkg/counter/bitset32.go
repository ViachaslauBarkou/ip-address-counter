package counter

import (
	"sync/atomic"
)

const arraySize32 = 1 << 27 // 2^32 / 32

type BitSet32 struct {
	array []uint32
}

func NewBitSet32() *BitSet32 {
	return &BitSet32{array: make([]uint32, arraySize32)}
}

func (b *BitSet32) SetBit(index int) {
	wordIndex, bitIndex := calculateIndexes(index, 32)
	b.array[wordIndex] |= 1 << bitIndex
}

func (b *BitSet32) AtomicSetBit(index int) bool {
	wordIndex, bitIndex := calculateIndexes(index, 32)
	bitOffset := uint32(1 << bitIndex)

	addr := &b.array[wordIndex]
	for {
		oldValue := atomic.LoadUint32(addr)
		if oldValue&bitOffset != 0 {
			return true
		}
		newValue := oldValue | bitOffset
		if atomic.CompareAndSwapUint32(addr, oldValue, newValue) {
			return false
		}
	}
}

func (b *BitSet32) IsBitSet(index int) bool {
	wordIndex, bitIndex := calculateIndexes(index, 32)
	return (b.array[wordIndex] & (1 << bitIndex)) != 0
}
