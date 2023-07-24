package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	inputFilename = "input"
	nbTopElves    = 3
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

	var (
		caloriesByElf []int64
		calories      int64
	)

	for i, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" {
			caloriesByElf = append(caloriesByElf, calories)
			calories = 0

			continue
		}

		c, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			return errors.Wrapf(err, "line %d: strconv.ParseInt", i+1)
		}

		calories += c
	}

	if calories > 0 {
		caloriesByElf = append(caloriesByElf, calories)
	}

	sort.Slice(caloriesByElf, func(i, j int) bool {
		return caloriesByElf[i] < caloriesByElf[j]
	})

	nbElves := len(caloriesByElf)
	if nbElves < nbTopElves {
		return fmt.Errorf("invalid input: expected at least %d elves, got %d\n", nbTopElves, nbElves)
	}

	var totalCaloriesTopElves int64

	for i := 0; i < nbTopElves; i++ {
		totalCaloriesTopElves += caloriesByElf[nbElves-1-i]
	}

	fmt.Printf("totalCaloriesTopElves: %d\n", totalCaloriesTopElves)

	return nil
}
