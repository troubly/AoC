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

type playerChoice struct {
	baseScore       int
	opponentChoices map[string]int // choice -> score
}

var (
	playerChoices = map[string]playerChoice{
		"X": {
			baseScore: 1,
			opponentChoices: map[string]int{
				"A": 3,
				"B": 0,
				"C": 6,
			},
		},
		"Y": {
			baseScore: 2,
			opponentChoices: map[string]int{
				"A": 6,
				"B": 3,
				"C": 0,
			},
		},
		"Z": {
			baseScore: 3,
			opponentChoices: map[string]int{
				"A": 0,
				"B": 6,
				"C": 3,
			},
		},
	}
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

		choices := strings.Split(l, " ")
		if len(choices) != 2 {
			return fmt.Errorf("line %d: unexpected input format", i+1)
		}

		s, err := computeRoundScore(choices[1], choices[0])
		if err != nil {
			return fmt.Errorf("line %d: computeRoundScore", i+1)
		}

		score += s
	}

	fmt.Printf("score: %d\n", score)

	return nil
}

func computeRoundScore(playerChoice, opponentChoice string) (int, error) {
	ch, ok := playerChoices[playerChoice]
	if !ok {
		return 0, fmt.Errorf("unexpected player choice '%s'", playerChoice)
	}

	score, ok := ch.opponentChoices[opponentChoice]
	if !ok {
		return 0, fmt.Errorf("unexpected opponent choice '%s'", opponentChoice)
	}

	return ch.baseScore + score, nil
}
