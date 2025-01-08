package counter

// calculateIndexes calculates the wordIndex and bitIndex based on the provided index and bitSize.
func calculateIndexes(index, bitSize int) (wordIndex, bitIndex int) {
	wordIndex = index / bitSize
	bitIndex = index % bitSize
	return
}

// calculatePartIndexes calculates the partIndex and bitOffset based on the bitIndex and partSize.
func calculatePartIndexes(bitIndex, partSize int) (partIndex, bitOffset int) {
	partIndex = bitIndex / partSize
	bitOffset = bitIndex % partSize
	return
}
