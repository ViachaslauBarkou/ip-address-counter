package counter_test

import (
	"bytes"
	"fmt"
	"ip-address-counter/pkg/counter"
	"testing"
)

func TestProcessReader(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		expected   int
		bitSetSize int
		shouldFail bool
	}{
		{
			name:       "Valid IPs with 32-bit BitSet",
			input:      "192.168.1.1\n192.168.1.2\n192.168.1.3\n192.168.1.4\n192.168.1.5\n192.168.1.6\n192.168.1.7\n192.168.1.8\n192.168.1.9\n192.168.1.10\n192.168.1.1\n192.168.1.2\n192.168.1.3\n192.168.1.4\n192.168.1.5\n",
			expected:   10,
			bitSetSize: 32,
			shouldFail: false,
		},
		{
			name:       "Valid IPs with 128-bit BitSet",
			input:      "10.0.0.1\n10.0.0.2\n10.0.0.3\n10.0.0.4\n10.0.0.5\n10.0.0.6\n10.0.0.7\n10.0.0.8\n10.0.0.9\n10.0.0.10\n10.0.0.1\n10.0.0.2\n10.0.0.3\n10.0.0.4\n10.0.0.5\n",
			expected:   10,
			bitSetSize: 128,
			shouldFail: false,
		},
		{
			name:       "Invalid IPs and duplicates",
			input:      "invalid\n10.0.0.1\n10.0.0.2\n10.0.0.3\n10.0.0.4\n10.0.0.5\n10.0.0.1\n10.0.0.2\n10.0.0.3\n10.0.0.4\n10.0.0.5\n",
			expected:   5,
			bitSetSize: 256,
			shouldFail: false,
		},
		{
			name:       "Empty input",
			input:      "",
			expected:   0,
			bitSetSize: 512,
			shouldFail: false,
		},
		{
			name:       "Invalid BitSet size",
			input:      "192.168.1.1\n",
			expected:   0,
			bitSetSize: 48,
			shouldFail: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var bitSet counter.BitSet
			var err error

			if test.bitSetSize == 32 {
				bitSet = counter.NewBitSet32()
			} else if test.bitSetSize%64 == 0 && test.bitSetSize >= 64 {
				bitSet, err = counter.NewBitSetN(test.bitSetSize)
			} else {
				err = fmt.Errorf("invalid bitSet size")
			}

			if test.shouldFail {
				if err == nil {
					t.Errorf("expected failure but got success")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			reader := bytes.NewReader([]byte(test.input))
			uniqueCount, err := counter.ProcessReader(reader, bitSet)
			if err != nil {
				t.Fatalf("unexpected error during processing: %v", err)
			}

			if uniqueCount != test.expected {
				t.Errorf("expected %d unique IPs, got %d", test.expected, uniqueCount)
			}
		})
	}
}
