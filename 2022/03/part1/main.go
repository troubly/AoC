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

	for i, rucksack := range lines {
		if rucksack == "" {
			continue
		}

		rsl := len(rucksack)

		if rsl == 0 || rsl%2 != 0 {
			return fmt.Errorf("line %d: unexpected len %d", i+1, rsl)
		}

		if !lettersRegexp.MatchString(rucksack) {
			return fmt.Errorf("line %d: unexpected format", i+1)
		}

		rawRs := []rune(rucksack)

		compt1, compt2 := rawRs[0:rsl/2], rawRs[rsl/2:rsl]

		items := make(map[rune]struct{}, rsl/2)
		for _, item := range compt1 {
			items[item] = struct{}{}
		}

		commonItems := make(map[rune]struct{})

		for _, item := range compt2 {
			if _, ok := items[item]; !ok {
				continue
			}

			commonItems[item] = struct{}{}
		}

		for item := range commonItems {
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
