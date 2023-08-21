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

	var prioritiesSum int

	for scanner.Scan() {
		rucksack := scanner.Text()

		rsl := len(rucksack)

		if rsl == 0 || rsl%2 != 0 {
			return fmt.Errorf("unexpected len %d", rsl)
		}

		if !lettersRegexp.MatchString(rucksack) {
			return errors.New("unexpected format")
		}

		compt1, compt2 := rucksack[0:rsl/2], rucksack[rsl/2:rsl]

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
