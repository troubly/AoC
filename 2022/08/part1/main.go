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

	nbVisible := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if isVisible(grid, i, j) {
				nbVisible++
			}
		}
	}

	fmt.Printf("nbVisible: %d\n", nbVisible)

	return nil
}

// returns true if the tree at coords (`x`, `y`) in `grid` is visible
func isVisible(grid [][]int, x, y int) bool {
	// edge of the grid is visible
	if x == 0 || x == len(grid)-1 {
		return true
	}

	if y == 0 || y == len(grid[x])-1 {
		return true
	}

	// interior of the grid
	h := grid[x][y]

	visible := true
	for i := 0; i < x; i++ {
		if h <= grid[i][y] {
			visible = false

			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for i := x + 1; i < len(grid); i++ {
		if h <= grid[i][y] {
			visible = false

			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for j := 0; j < y; j++ {
		if h <= grid[x][j] {
			visible = false

			break
		}
	}
	if visible {
		return true
	}

	visible = true
	for j := y + 1; j < len(grid[x]); j++ {
		if h <= grid[x][j] {
			visible = false

			break
		}
	}
	if visible {
		return true
	}

	return false
}
