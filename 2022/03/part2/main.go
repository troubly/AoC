package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
	data, err := os.ReadFile(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.ReadFile")
	}

	lines := strings.Split(string(data), "\n")

	var prioritiesSum int

	for i := 0; i < len(lines); i += elfGroupSize {
		low, high := i, i+elfGroupSize
		if high > len(lines) {
			high = len(lines)
		}

		groupItems := lines[low:high]

		itemTypeCount := make(map[rune]int)

		for _, rucksack := range groupItems {
			if rucksack == "" {
				continue
			}

			if !lettersRegexp.MatchString(rucksack) {
				return fmt.Errorf("line %d: unexpected format", i+1)
			}

			items := make(map[rune]struct{})

			rawRs := []rune(rucksack)

			for _, item := range rawRs {
				items[item] = struct{}{}
			}

			for item := range items {
				itemTypeCount[item]++
			}
		}

		for item, count := range itemTypeCount {
			if count < 3 {
				continue
			}

			prioritiesSum += computePriority(item)
		}
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
