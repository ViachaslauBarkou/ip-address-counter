package counter

import (
	"bufio"
	"io"
	"net"
)

func ProcessReader(reader io.Reader, bitset BitSet) (int, error) {
	scanner := bufio.NewScanner(reader)
	uniqueCount := 0

	for scanner.Scan() {
		ip := net.ParseIP(scanner.Text()).To4()
		if ip == nil {
			continue
		}

		index := ipToIndex(ip)
		if !bitset.IsBitSet(index) {
			uniqueCount++
			bitset.SetBit(index)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return uniqueCount, nil
}

func ipToIndex(ip net.IP) int {
	return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3])
}
