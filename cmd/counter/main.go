package main

import (
	"fmt"
	"ip-address-counter/cmd/config"
	"ip-address-counter/pkg/counter"
	"os"
	"runtime"
	"time"
)

func main() {
	cfg := config.ParseFlags()

	file, err := os.Open(cfg.TestFile)
	if err != nil {
		fmt.Println("Error opening test file:", err)
		return
	}
	defer file.Close()

	bitset, err := getBitSet(cfg.BitSetSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	start := time.Now()
	var memStatsStart runtime.MemStats
	runtime.ReadMemStats(&memStatsStart)

	fmt.Printf("Counting... file: %s | bitset: %d\n", cfg.TestFile, cfg.BitSetSize)
	uniqueCount, err := counter.ProcessReader(file, bitset)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return
	}

	var memStatsEnd runtime.MemStats
	runtime.ReadMemStats(&memStatsEnd)
	elapsed := time.Since(start)
	memoryUsed := memStatsEnd.Alloc - memStatsStart.Alloc

	fmt.Printf("Unique IPs: %d\n", uniqueCount)
	fmt.Printf("Execution time: %v\n", elapsed)
	fmt.Printf("Memory used: %d bytes\n", memoryUsed)
}

func getBitSet(size int) (counter.BitSet, error) {
	if size == 32 {
		return counter.NewBitSet32(), nil
	} else if size%64 == 0 && size >= 64 {
		bitSet, err := counter.NewBitSetN(size)
		if err != nil {
			return nil, err
		}
		return bitSet, nil
	}
	return nil, fmt.Errorf("wrong size: %d. Use 32 or any multiple of 64 (64, 128, 256, ...)", size)
}
