package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

var (
	lsRegexp      = regexp.MustCompile(`^\$ ls$`)
	cdRegexp      = regexp.MustCompile(`(?s)^\$ cd (/|[a-zA-Z]+|\.\.)$`)
	contentRegexp = regexp.MustCompile(`(?s)^(dir|[0-9]+) (.+)$`)
)

const (
	inputFilename = "input"
)

type file struct {
	name string
	size int64
}

type dir struct {
	name   string
	parent *dir

	subdirs []*dir
	files   []*file

	size int64
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

	var (
		currentDir *dir
		listing    bool
		root       = &dir{
			name: "/",
		}
	)

	for scanner.Scan() {
		l := scanner.Text()

		// cd
		if res := cdRegexp.FindStringSubmatch(l); res != nil {
			switch res[1] {
			case "/":
				currentDir = root
			case "..":
				currentDir = currentDir.parent
			default:
				var found bool

				for _, d := range currentDir.subdirs {
					if d.name == res[1] {
						currentDir = d
						found = true

						break
					}
				}

				if !found {
					return fmt.Errorf("cd '%s'", res[1])
				}
			}

			listing = false

			continue
		}

		// ls
		if lsRegexp.MatchString(l) {
			listing = true

			continue
		}

		// dir / file
		if !listing {
			return errors.New("unexpected input: not listing")
		}

		res := contentRegexp.FindStringSubmatch(l)
		if res == nil {
			return errors.New("unexpected input: unprocessable content format")
		}

		if res[1] == "dir" {
			currentDir.subdirs = append(currentDir.subdirs, &dir{
				name:   res[2],
				parent: currentDir,
			})
		} else {
			size, _ := strconv.ParseInt(res[1], 10, 64)

			currentDir.files = append(currentDir.files, &file{
				name: res[2],
				size: size,
			})
		}
	}

	// compute size of all dirs
	computeSize := func(d *dir) {
		var size int64

		for _, f := range d.files {
			size += f.size
		}

		for _, sd := range d.subdirs {
			size += sd.size
		}

		d.size = size
	}

	walkDir(root, computeSize)

	// part1
	var totalSizeSmallDirs int64

	sumSmallDirs := func(d *dir) {
		if d.size <= 100_000 {
			totalSizeSmallDirs += d.size
		}
	}

	walkDir(root, sumSmallDirs)

	fmt.Printf("totalSizeSmallDirs: %d\n", totalSizeSmallDirs)

	// part2
	var dirs []*dir

	getDirs := func(d *dir) {
		dirs = append(dirs, &dir{
			name: d.name,
			size: d.size,
		})
	}

	walkDir(root, getDirs)

	freeSpaceNeeded := 30_000_000 - (70_000_000 - root.size)

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	var dirToDelete *dir

	for _, d := range dirs {
		if d.size > freeSpaceNeeded {
			dirToDelete = d

			break
		}
	}

	if dirToDelete == nil {
		return errors.New("failed to find a dir to delete")
	}

	fmt.Printf("sizeOfDirToDelete: %d\n", dirToDelete.size)

	return nil
}

func walkDir(d *dir, fn func(*dir)) {
	for _, sd := range d.subdirs {
		walkDir(sd, fn)
	}

	fn(d)
}
