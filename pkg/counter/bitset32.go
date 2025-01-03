package counter

const arraySize32 = 1 << 27 // 2^32 / 32

type BitSet32 struct {
	array []uint32
}

func NewBitSet32() *BitSet32 {
	return &BitSet32{array: make([]uint32, arraySize32)}
}

func (b *BitSet32) SetBit(index int) {
	wordIndex := index / 32
	bitIndex := index % 32
	b.array[wordIndex] |= 1 << bitIndex
}

func (b *BitSet32) IsBitSet(index int) bool {
	wordIndex := index / 32
	bitIndex := index % 32
	return (b.array[wordIndex] & (1 << bitIndex)) != 0
}
