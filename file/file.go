package file

import (
	"bufio"
	"os"
)

// Read reads the file at path and returns the result
func Read(path string) (lines [][]rune, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, []rune(s.Text()))
	}
	return
}
