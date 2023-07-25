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
	pairsRegexp = regexp.MustCompile("^[0-9]+-[0-9]+,[0-9]+-[0-9]+$")
)

type interval struct {
	begin int64
	end   int64
}

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

	var nbFullyContains int

	for scanner.Scan() {
		pairs := scanner.Text()

		if !pairsRegexp.MatchString(pairs) {
			return errors.New("unexpected format")
		}

		p := strings.Split(pairs, ",")

		intvl1, err := intervalFromPairString(p[0])
		if err != nil {
			return errors.Wrap(err, "intervalFromPairString")
		}

		intvl2, err := intervalFromPairString(p[1])
		if err != nil {
			return errors.Wrap(err, "intervalFromPairString")
		}

		if intvl1.fullyContains(intvl2) || intvl2.fullyContains(intvl1) {
			nbFullyContains++
		}
	}

	fmt.Printf("nbFullyContains: %d\n", nbFullyContains)

	return nil
}

func intervalFromPairString(p string) (*interval, error) {
	endpoints := strings.Split(p, "-")

	begin, err := strconv.ParseInt(endpoints[0], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "strconv.ParseInt")
	}

	end, err := strconv.ParseInt(endpoints[1], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "strconv.ParseInt")
	}

	if begin > end {
		return nil, errors.New("invalid interval endpoints")
	}

	res := &interval{
		begin: begin,
		end:   end,
	}

	return res, nil
}

func (i1 *interval) fullyContains(i2 *interval) bool {
	return i1.begin <= i2.begin && i1.end >= i2.end
}
