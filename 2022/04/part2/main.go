package main

import (
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
	data, err := os.ReadFile(inputFilename)
	if err != nil {
		return errors.Wrap(err, "os.ReadFile")
	}

	lines := strings.Split(string(data), "\n")

	var nbOverlaps int

	for i, pairs := range lines {
		if pairs == "" {
			continue
		}

		if !pairsRegexp.MatchString(pairs) {
			return fmt.Errorf("line %d: unexpected format", i+1)
		}

		p := strings.Split(pairs, ",")

		intvl1, err := intervalFromPairString(p[0])
		if err != nil {
			return errors.Wrapf(err, "line %d: intervalFromPairString", i+1)
		}

		intvl2, err := intervalFromPairString(p[1])
		if err != nil {
			return errors.Wrapf(err, "line %d: intervalFromPairString", i+1)
		}

		if intvl1.overlaps(intvl2) {
			nbOverlaps++
		}
	}

	fmt.Printf("nbOverlaps: %d\n", nbOverlaps)

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

func (i1 *interval) overlaps(i2 *interval) bool {
	switch {
	case i1.begin == i2.begin:
		return true
	case i1.begin < i2.begin:
		return i1.end >= i2.begin
	case i1.begin > i2.begin: // default
		return i2.end >= i1.begin
	}

	return false // unreachable
}
