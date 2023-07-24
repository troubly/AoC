package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

const (
	inputFilename = "../input"
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

	lines := strings.Split(string(data), "\n")

	var score int

	for i, l := range lines {
		if l == "" {
			continue
		}

		tokens := strings.Split(l, " ")
		if len(tokens) != 2 {
			return fmt.Errorf("line %d: unexpected input format", i+1)
		}

		s, err := computeRoundScore(tokens[0], tokens[1])
		if err != nil {
			return fmt.Errorf("line %d: computeRoundScore", i+1)
		}

		score += s
	}

	fmt.Printf("score: %d\n", score)

	return nil
}

func computeRoundScore(opponentChoice, playerStrategy string) (int, error) {
	var score int

	switch playerStrategy {
	case "X":
		score = 0

		switch opponentChoice {
		case "A":
			// playerMove = "C"
			score += 3
		case "B":
			// playerMove = "A"
			score += 1
		case "C":
			// playerMove = "B"
			score += 2
		default:
			return 0, fmt.Errorf("unexpected opponent choice '%s'", opponentChoice)
		}
	case "Y":
		score = 3

		switch opponentChoice {
		case "A":
			// playerMove = "A"
			score += 1
		case "B":
			// playerMove = "B"
			score += 2
		case "C":
			// playerMove = "C"
			score += 3
		default:
			return 0, fmt.Errorf("unexpected opponent choice '%s'", opponentChoice)
		}
	case "Z":
		score = 6

		switch opponentChoice {
		case "A":
			// playerMove = "B"
			score += 2
		case "B":
			// playerMove = "C"
			score += 3
		case "C":
			// playerMove = "A"
			score += 1
		default:
			return 0, fmt.Errorf("unexpected opponent choice '%s'", opponentChoice)
		}
	default:
		return 0, fmt.Errorf("unexpected player strategy '%s'", playerStrategy)
	}

	return score, nil
}
