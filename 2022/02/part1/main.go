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
	f, err := os.Open(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.Open")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var score int

	for scanner.Scan() {
		l := scanner.Text()

		choices := strings.Split(l, " ")
		if len(choices) != 2 {
			return errors.New("unexpected input format")
		}

		s, err := computeRoundScore(choices[1], choices[0])
		if err != nil {
			return errors.Wrap(err, "computeRoundScore")
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
