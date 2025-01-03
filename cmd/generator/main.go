package main

import (
	"bufio"
	"fmt"
	"ip-address-counter/cmd/config"
	"math/rand"
	"os"
)

func main() {
	cfg := config.ParseFlags()

	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		fmt.Println("Error creating data directory:", err)
		return
	}

	file, err := os.Create(cfg.TestFile)
	if err != nil {
		fmt.Println("Error creating test file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Printf("Generating %d IP addresses in %s ...\n", cfg.NumAddresses, cfg.TestFile)
	for i := 0; i < cfg.NumAddresses; i++ {
		ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		writer.WriteString(ip + "\n")
		if i > 0 && i%10_000_000 == 0 {
			fmt.Printf("Generated %d addresses...\n", i)
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing to file:", err)
		return
	}

	fmt.Println("IP address generation completed.")
}
