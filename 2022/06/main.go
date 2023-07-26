package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	inputFilename     = "input"
	nbCharPacketStart = 14 // 4 for part1
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("run failed: %v\n", err)
	}
}

func run() error {
	data, err := os.ReadFile(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.ReadFile")
	}

	bufferPos := 0
	processedChars := 0

LOOP:
	for bufferPos-nbCharPacketStart < len(data) {
		m := make(map[byte]int, nbCharPacketStart)

		for i := bufferPos; i < bufferPos+nbCharPacketStart; i++ {
			idx, ok := m[data[i]]
			if !ok {
				m[data[i]] = i
			} else {
				bufferPos = idx + 1

				continue LOOP
			}
		}

		processedChars = bufferPos + nbCharPacketStart

		break
	}

	if processedChars == 0 {
		return errors.New("failed to find start of packet marker")
	}

	fmt.Printf("processedChars: %d\n", processedChars)

	return nil
}
