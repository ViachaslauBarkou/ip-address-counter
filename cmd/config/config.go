package config

import (
	"flag"
	"path/filepath"
	"runtime"
)

type Config struct {
	TestFile       string
	BitSetSize     int
	UseConcurrency bool
	Workers        int
	NumAddresses   int
}

// Default values
const (
	DefaultTestFile       = "ip_addresses"
	DefaultDataFolder     = "data"
	DefaultBitSetSize     = 32
	DefaultUseConcurrency = false
	DefaultNumAddresses   = 10_000_000
)

func ParseFlags() Config {
	testFile := flag.String("test_file", DefaultTestFile, "Path to the test file (relative to the data folder)")
	bitSetSize := flag.Int("bit_set", DefaultBitSetSize, "Choose BitSet bitSetSize: 32, 128, 256, 512")
	useConcurrency := flag.Bool("concurrent", DefaultUseConcurrency, "Enable multithreading mode")
	workers := flag.Int("workers", runtime.NumCPU(), "Number of concurrent workers (used in multithreading mode)") // Default is number of CPU cores
	numAddresses := flag.Int("count", DefaultNumAddresses, "Number of addresses to generate (for generator)")
	flag.Parse()

	return Config{
		TestFile:       filepath.Join(DefaultDataFolder, *testFile),
		BitSetSize:     *bitSetSize,
		UseConcurrency: *useConcurrency,
		Workers:        *workers,
		NumAddresses:   *numAddresses,
	}
}
