# IP Address Counter

## Description
Go-based console application that allows to count unique IPs and generate a random number IPv4 addresses.

## Features
- **Count Unique IPs**: Count unique IP addresses in a file with support for customizable BitSet sizes for various system configurations.  
- **Generate IP Addresses**: Create a file with a specified number of random IPv4 addresses.  

## Prerequisites
- Go 1.22+

## Installation
1. Clone the repository:
```bash
git clone https://github.com/your-username/ip-address-counter.git
cd ip-address-counter
```
2. Build the project:
```bash
go build ./...
```

## Running Tests
```bash
go test ./...
```

## Configuration Parameters
### Flags for `counter`
- **`-test_file`**: Path to the test file (relative to the `data` folder). Default: `data/ip_addresses`.
- **`-bit_set`**: Choose the BitSet size: 32 or any multiple of 64 (64, 128, 256, ...). Default: 256.
### Flags for `generator`
- **`-test_file`**: Path to the output file (relative to the `data` folder). Default: `data/ip_addresses`.
- **`-count`**: Number of IP addresses to generate. Default: 10,000,000.

## Running the Application (examples)
### Option 1: Using `go run`
#### Run the IP Address Generator
```bash
go run cmd/generator/main.go -output ip_addresses -count 1000000
```
#### Run the IP Address Counter
```bash
go run cmd/counter/main.go -test_file ip_addresses -bit_set 256
```
### Option 2: Using Precompiled Binaries
#### Build the binaries
To create standalone binaries for the generator and counter:
1. Build the generator binary:
```bash
go build -o ip-address-generator cmd/generator/main.go
```
2. Build the counter binary:
```bash
go build -o ip-address-counter cmd/counter/main.go
```
#### Run the binaries
##### Run the IP Address Generator
```bash
./ip-address-generator -output ip_addresses -count 1000000
```
##### Run the IP Address Counter
```bash
./ip-address-counter -test_file ip_addresses -bit_set 256
```

## Execution Results
### System Configuration
The application was tested on the following system:
- **Processor**: 2.3 GHz 8-Core Intel Core i9.
- **Bit Operation Support**: Supports AVX2, enabling 256-bit SIMD instructions for optimal performance.
- **Operating System**: macOS.
### Results of running the program on the file with a size of ~114.5 GB:

| BitSet | Unique IPs | Execution Time | Memory Used |
|--------|------------| -------------- |-------------|
| 256    | 1000000000 | ~14m29s        | ~493.93 MB  |

_For all BitSet sizes in single-threaded mode, the application uses up to 512 MB of memory for the full IPv4 address space. This is calculated as (example): `Memory = (2^32 / 256) * 32 bytes = 512 MB`. For smaller datasets, memory usage is reduced as fewer bits are set in the BitSet.  
On described system, testing was conducted with the optimal configuration for a 256-bit BitSet._

## Author
Viachaslau Barkou

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
