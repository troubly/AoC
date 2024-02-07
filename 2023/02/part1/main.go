package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	inputFilename = "../input"

	nbBlueMax  = 14
	nbRedMax   = 12
	nbGreenMax = 13
)

var (
	gameIDCatcher = regexp.MustCompile(`Game ([0-9]+)`)
	blueCatcher   = regexp.MustCompile(`([0-9]+) blue`)
	redCatcher    = regexp.MustCompile(`([0-9]+) red`)
	greenCatcher  = regexp.MustCompile(`([0-9]+) green`)
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

	var sum int64

	for scanner.Scan() {
		l := scanner.Text()

		parts := strings.Split(l, ":")

		matches := gameIDCatcher.FindAllStringSubmatch(parts[0], 1)
		if len(matches) != 1 {
			return fmt.Errorf("unexpected input: %s", l)
		}

		gameID, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			return errors.Wrap(err, "strconv.ParseInt")
		}

		sets := strings.Split(parts[1], ";")

		var valid bool

		for _, set := range sets {
			valid, err = isSetValid(set)
			if err != nil {
				return errors.Wrap(err, "isSetValid")
			}

			if !valid {
				break
			}
		}

		if valid {
			sum += gameID
		}
	}

	fmt.Println(sum)

	return nil
}

func isSetValid(set string) (bool, error) {
	matches := blueCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbBlue, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			return false, errors.Wrap(err, "strconv.ParseInt")
		}

		if nbBlue > nbBlueMax {
			return false, nil
		}
	}

	matches = redCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbRed, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			return false, errors.Wrap(err, "strconv.ParseInt")
		}

		if nbRed > nbRedMax {
			return false, nil
		}
	}

	matches = greenCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbGreen, err := strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			return false, errors.Wrap(err, "strconv.ParseInt")
		}

		if nbGreen > nbGreenMax {
			return false, nil
		}
	}

	return true, nil
}
