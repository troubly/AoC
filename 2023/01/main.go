package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const (
	inputFilename = "input"
)

var (
	catcherForwards  = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9]{1})")
	catcherBackwards = regexp.MustCompile("(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[0-9]{1})")
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

		matches := catcherForwards.FindAllStringSubmatch(l, 1)

		if len(matches) < 1 {
			return fmt.Errorf("failed to find matches in '%s'", l)
		}

		first, err := extractDigit(matches[0][0])
		if err != nil {
			return errors.Wrap(err, "extractDigit")
		}

		// Go does't have overlapping matches
		matches = catcherBackwards.FindAllStringSubmatch(revertString(l), 1)

		last, err := extractDigit(revertString(matches[0][0]))
		if err != nil {
			return errors.Wrap(err, "extractDigit")
		}

		sum += first*10 + last
	}

	fmt.Printf("sum: %d\n", sum)

	return nil
}

func extractDigit(s string) (int64, error) {
	switch s {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "strconv.ParseInt")
		}

		return val, nil
	}
}

func revertString(s string) string {
	rawString := []rune(s)

	for i := 0; i < len(s)/2; i++ {
		rawString[i], rawString[len(s)-1-i] = rawString[len(s)-1-i], rawString[i]
	}

	return string(rawString)
}
