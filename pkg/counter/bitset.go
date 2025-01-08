package counter

type BitSet interface {
	SetBit(index int)
	AtomicSetBit(index int) bool
	IsBitSet(index int) bool
}
