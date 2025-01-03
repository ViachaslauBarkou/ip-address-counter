package counter

type BitSet interface {
	SetBit(index int)
	IsBitSet(index int) bool
}
