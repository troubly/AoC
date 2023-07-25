package main

import (
	"bufio"
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
	f, err := os.Open(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.Open")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var score int

	for scanner.Scan() {
		l := scanner.Text()

		tokens := strings.Split(l, " ")
		if len(tokens) != 2 {
			return errors.New("unexpected input format")
		}

		s, err := computeRoundScore(tokens[0], tokens[1])
		if err != nil {
			return errors.Wrap(err, "computeRoundScore")
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
