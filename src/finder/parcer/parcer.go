package parcer

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"io/ioutil"
)


func GetBlockByAnchors(markdown string, anchor string) []string {
	
	block := make([]string, 0)
	headFound := false

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// Scans line by line
	for scanner.Scan() {
		if targetAnchorFound(scanner.Text(), anchor) == true {
			block = append(block, scanner.Text() + "\n")
			headFound = true

			continue

		}

		if headFound == true {

			if startsWithHash(scanner.Text()) == true {
				break
			}

			block = append(block, scanner.Text() + "\n")

		}

	}


	if len(block) == 0 {
		log.Print("EMPTY")
		block = append(block, "Nothing found, sorry")
	}
	return block
}


func GetAllAnchors(markdown string) []string {
	// Holds all the anchors in slice
	
	s := make([]string, 0)


	scanner := bufio.NewScanner(strings.NewReader(markdown))
	// Scans line by line
	for scanner.Scan() {

		if startsWithHash(scanner.Text()) == true {
			s = append(s, scanner.Text() + "\n")
		}
	}
	fmt.Println(s)
	return s
}

func targetAnchorFound(line string, anchor string) bool {
	result := false
	re, _ := regexp.Compile(anchor)
	matches := re.FindStringSubmatch(line)
	if startsWithHash(line) == true {
		if matches != nil {
			result = true
		}

	}
	return result
}

func startsWithHash(line string) bool {
	return strings.HasPrefix(line, "#")
}

func ParseMarkdownFile(fileName string) (string, error) {
	file, err := fileToString(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file, nil
}

func fileToString(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	s := string(bytes)
	return s, nil
}

