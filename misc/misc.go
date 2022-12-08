package misc

import (
	"bufio"
	"os"
)

func GetInputFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

func ReadInputRows(path string) []string {
	f := GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows
}

