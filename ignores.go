package ignores

import (
	"bufio"
	"os"
	"path/filepath"
)

var (
	paths []string
)

type Ignore struct {
	paths []string
}

/*
Matches against a path to see if it should be ignored or not, returning true if it should, otherwise false.
This uses path/filepath.Match to to the matching so the only case where it would return error is when the pattern is malformed.
*/

func (ig *Ignore) Match(path string) (bool, error) {
	for _, igPath := range ig.paths {
		match, err := filepath.Match(igPath, path)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

/*
Creates a new instance of Ignore. It reads from 'path' line-by-line. The ignore pattern is the same as it is in path/filepath.Match
*/
func New(path string) (*Ignore, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ig := &Ignore{
		paths: make([]string, 0, 10),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ig.paths = append(ig.paths, scanner.Text())
	}
	return ig, nil
}
