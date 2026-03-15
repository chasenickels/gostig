package platform

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ReadFile(filename string) (osrelease map[string]string, err error) {
	osrelease = make(map[string]string)
	lines, err := parseFile(filename)
	if err != nil {
		return
	}

	for _, v := range lines {
		key, value, err := parseLine(v)
		if err == nil {
			osrelease[key] = value
		}
	}
	return
}

func parseFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseLine(line string) (key string, value string, err error) {
	err = nil
	if len(line) == 0 {
		err = errors.New("Skipping: line is of zero length.")
		return
	}

	if line[0] == '#' {
		err = errors.New("Skipping: Comment.")
		return
	}

	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		err = errors.New("Could not extract key:value pair.")
		return
	}

	key = strings.Trim(splitString[0], " ")
	value = strings.Trim(splitString[1], " ")
	if strings.ContainsAny(value, `"`) {
		first := string(value[0:1])
		last := string(value[len(value)-1:])
		if first == last && strings.ContainsAny(first, `"'`) {
			value = strings.ReplaceAll(value, `"`, ``)
			value = strings.ReplaceAll(value, `'`, ``)
		}
	}

	return
}
