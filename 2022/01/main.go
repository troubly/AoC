package main

import (
	"bufio"
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
	f, err := os.Open(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.Open")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		caloriesByElf []int64
		calories      int64
	)

	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())

		if l == "" {
			caloriesByElf = append(caloriesByElf, calories)
			calories = 0

			continue
		}

		c, err := strconv.ParseInt(l, 10, 64)
		if err != nil {
			return errors.Wrap(err, "strconv.ParseInt")
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
