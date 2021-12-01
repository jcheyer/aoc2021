package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadString reads a file into a slice of lines
func ReadInputLines(filename string) ([]string, error) {
	maxCapacity := 65536

	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result, nil
}

// ParseLines2Int64 takes a slice of strings and returns a slice of int64
func ParseLines2Int(lines []string) ([]int64, error) {
	res := make([]int64, len(lines))
	for i, line := range lines {
		convValue, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		res[i] = convValue
	}
	return res, nil
}
