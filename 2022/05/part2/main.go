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
	procedure = regexp.MustCompile("(?s)^move ([0-9]+) from ([0-9]+) to ([0-9]+)$")
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
		rawStacks [][]string
		stacks    []stack[string]
	)

	for scanner.Scan() {
		l := scanner.Text()

		// read crate line
		if strings.Contains(l, "[") {
			var rawStack []string

			rawLine := []rune(l)

			for i := 0; i < len(rawLine); i += 4 {
				if rawLine[i] == '[' {
					rawStack = append(rawStack, string(rawLine[i+1]))
				} else {
					rawStack = append(rawStack, "")
				}
			}

			rawStacks = append(rawStacks, rawStack)

			continue
		}

		// build stacks
		if stacks == nil {
			stacks = make([]stack[string], len(rawStacks[len(rawStacks)-1]))

			for i := len(rawStacks) - 1; i >= 0; i-- {
				for j, v := range rawStacks[i] {
					if v != "" {
						stacks[j].push(v)
					}
				}
			}
		}

		if !procedure.MatchString(l) {
			continue
		}

		// run procedure
		res := procedure.FindStringSubmatch(l)

		nbPop, _ := strconv.ParseInt(res[1], 10, 64)
		from, _ := strconv.ParseInt(res[2], 10, 64)
		to, _ := strconv.ParseInt(res[3], 10, 64)

		// could also use another tmp stack x)
		cratesToMove := make([]string, nbPop)

		for i := 0; i < int(nbPop); i++ {
			v, err := stacks[from-1].pop()
			if err != nil {
				return errors.Wrap(err, "stack.pop")
			}

			cratesToMove[i] = v
		}

		for i := len(cratesToMove) - 1; i >= 0; i-- {
			stacks[to-1].push(cratesToMove[i])
		}
	}

	var output string

	for _, s := range stacks {
		v, err := s.pop()
		if err != nil {
			return errors.Wrap(err, "s.pop")
		}

		output += v
	}

	fmt.Printf("output: %s\n", output)

	return nil
}
