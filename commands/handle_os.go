package commands

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func getFileLines(filename string) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func createFile(filename string, lines []string) error {
	fmt.Printf("output file name: %s\n", filename)
	fp, err := os.Create(filename)
	defer fp.Close()
	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err := fp.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func readCfgCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var lines [][]string

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		lines = append(lines, row)
	}
	return lines, nil
}
