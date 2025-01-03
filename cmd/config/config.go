package config

import (
	"flag"
	"path/filepath"
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
	DefaultTestFile     = "ip_addresses"
	DefaultDataFolder   = "data"
	DefaultBitSetSize   = 32
	DefaultNumAddresses = 10_000_000
)

func ParseFlags() Config {
	testFile := flag.String("test_file", DefaultTestFile, "Path to the test file (relative to the data folder)")
	bitSetSize := flag.Int("bit_set", DefaultBitSetSize, "Choose BitSet bitSetSize: 32, 128, 256, 512")
	numAddresses := flag.Int("count", DefaultNumAddresses, "Number of addresses to generate (for generator)")
	flag.Parse()

	return Config{
		TestFile:     filepath.Join(DefaultDataFolder, *testFile),
		BitSetSize:   *bitSetSize,
		NumAddresses: *numAddresses,
	}
}
