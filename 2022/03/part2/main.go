package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode"

	"github.com/pkg/errors"
)

const (
	inputFilename      = "../input"
	lowerCaseAPriority = 1
	upperCaseAPriority = 27
	elfGroupSize       = 3
)

var (
	lettersRegexp = regexp.MustCompile("^[a-zA-Z]+$")
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
		prioritiesSum int
		groupItems    = make([]string, 0, elfGroupSize)
	)

	for scanner.Scan() {
		groupItems = append(groupItems, scanner.Text())

		if len(groupItems) < elfGroupSize {
			continue
		}

		itemTypeCount := make(map[rune]int)

		for _, rucksack := range groupItems {
			if !lettersRegexp.MatchString(rucksack) {
				return errors.New("unexpected format")
			}

			items := make(map[rune]struct{})

			for _, item := range rucksack {
				items[item] = struct{}{}
			}

			for item := range items {
				itemTypeCount[item]++
			}
		}

		for item, count := range itemTypeCount {
			if count < elfGroupSize {
				continue
			}

			prioritiesSum += computePriority(item)
		}

		// reset
		groupItems = make([]string, 0, elfGroupSize)
	}

	fmt.Printf("prioritiesSum: %d\n", prioritiesSum)

	return nil
}

func computePriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item - 'a' + lowerCaseAPriority)
	} else {
		return int(item - 'A' + upperCaseAPriority)
	}
}
