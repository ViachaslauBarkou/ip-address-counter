package counter

import (
	"bufio"
	"io"
	"net"
	"sync"
	"sync/atomic"
)

const chunkSize = 1000000 // Number of lines for each goroutine

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

func ProcessReaderWithConcurrency(reader io.Reader, bitset BitSet, numWorkers int) (int, error) {
	scanner := bufio.NewScanner(reader)
	lines := make(chan []string, numWorkers)

	var wg sync.WaitGroup
	uniqueCount := int32(0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		chunk := []string{}
		for scanner.Scan() {
			chunk = append(chunk, scanner.Text())
			if len(chunk) >= chunkSize {
				lines <- chunk
				chunk = []string{}
			}
		}
		if len(chunk) > 0 {
			lines <- chunk
		}
		close(lines)
	}()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chunk := range lines {
				localCount := processChunkWithAtomic(chunk, bitset)
				atomic.AddInt32(&uniqueCount, int32(localCount))
			}
		}()
	}

	wg.Wait()
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return int(uniqueCount), nil
}

func processChunkWithAtomic(chunk []string, bitset BitSet) int {
	localUnique := 0
	for _, line := range chunk {
		ip := net.ParseIP(line).To4()
		if ip == nil {
			continue
		}
		index := ipToIndex(ip)

		if !bitset.AtomicSetBit(index) {
			localUnique++
		}
	}
	return localUnique
}

func ipToIndex(ip net.IP) int {
	return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3])
}
