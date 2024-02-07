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

		sets := strings.Split(parts[1], ";")

		var nbB, nbR, nbG int64

		for _, set := range sets {
			b, r, g, err := cubesInSet(set)
			if err != nil {
				return errors.Wrap(err, "cubesInSet")
			}

			if b > nbB {
				nbB = b
			}

			if r > nbR {
				nbR = r
			}

			if g > nbG {
				nbG = g
			}
		}

		power := nbB * nbR * nbG

		sum += power
	}

	fmt.Println(sum)

	return nil
}

func cubesInSet(set string) (nbBlue, nbRed, nbGreen int64, err error) {
	matches := blueCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbBlue, err = strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			err = errors.Wrap(err, "strconv.ParseInt")
			return
		}
	}

	matches = redCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbRed, err = strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			err = errors.Wrap(err, "strconv.ParseInt")

			return
		}
	}

	matches = greenCatcher.FindAllStringSubmatch(set, 1)
	if len(matches) == 1 {
		nbGreen, err = strconv.ParseInt(matches[0][1], 10, 64)
		if err != nil {
			err = errors.Wrap(err, "strconv.ParseInt")

			return
		}
	}

	return
}
