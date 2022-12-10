package misc

import (
	"bufio"
	"fmt"
	"io/fs"
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

func NewDay(day int) {
	var folder string
	if day < 10 {
		folder = fmt.Sprintf("Day0%d", day)
	} else {
		folder = fmt.Sprintf("Day%d", day)
	}
	if _, err := os.Stat(folder); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(folder, fs.ModePerm)
			if err != nil {
				fmt.Printf("Folder not created.")
				panic(err)
			}
			_, err = os.Create(folder + string(os.PathSeparator) + folder + ".go")
			_, err = os.Create(folder + string(os.PathSeparator) + "input.txt")

			// Doesn't work yet
			/*	cmd := exec.Command("cobra-cli", "add "+folder)
				if err := cmd.Run(); err != nil {
					log.Fatal("Tried running ", cmd.String(), " --> ", err)
				}*/
		} else {
			fmt.Printf("Folder %s already exists.\n", folder)
		}
	}
}
