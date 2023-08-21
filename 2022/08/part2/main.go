package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var (
	lsRegexp      = regexp.MustCompile(`^\$ ls$`)
	cdRegexp      = regexp.MustCompile(`(?s)^\$ cd (/|[a-zA-Z]+|\.\.)$`)
	contentRegexp = regexp.MustCompile(`(?s)^(dir|[0-9]+) (.+)$`)
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

	var grid [][]int

	for scanner.Scan() {
		l := scanner.Text()

		trees := make([]int, len(l))

		for i, height := range l {
			h := int(height - '0')
			if h > 10 || h < 0 {
				return errors.New("invalid input")
			}

			trees[i] = h
		}

		grid = append(grid, trees)
	}

	maxScore := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			s := computeScenicScore(grid, i, j)
			if s > maxScore {
				maxScore = s
			}
		}
	}

	fmt.Printf("maxScore: %d\n", maxScore)

	return nil
}

func computeScenicScore(grid [][]int, x, y int) int {
	h := grid[x][y]

	// north
	var sn int
	for j := y - 1; j >= 0; j-- {
		sn++

		if grid[x][j] >= h {
			break
		}
	}

	// south
	var ss int
	for j := y + 1; j < len(grid[x]); j++ {
		ss++

		if grid[x][j] >= h {
			break
		}
	}

	// east
	var se int
	for i := x - 1; i >= 0; i-- {
		se++

		if grid[i][y] >= h {
			break
		}
	}

	// west
	var sw int
	for i := x + 1; i < len(grid); i++ {
		sw++

		if grid[i][y] >= h {
			break
		}
	}

	return sn * ss * se * sw
}
