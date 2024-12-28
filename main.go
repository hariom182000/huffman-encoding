package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFrequencyTable(fileName string) map[rune]int {
	frequencyMap := map[rune]int{}
	file, err := os.Open(fileName)
	if err != nil {
		return frequencyMap
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		frequencyMap[c]++
	}
	return frequencyMap
}

func encodeFile(fileName string, runeToPrefixTable map[rune]string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()
	encodedFileName := "encoded-" + fileName
	encodedFile, err := os.Create(encodedFileName)
	if err != nil {
		return ""
	}
	writer := bufio.NewWriter(encodedFile)
	reader := bufio.NewReader(file)
	defer encodedFile.Close()
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		writer.WriteString(runeToPrefixTable[c])
	}
	writer.Flush()
	return encodedFileName

}

func decodeFile(fileName string, prefixToRuneMap map[string]rune) {

	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	var charcter strings.Builder
	reader := bufio.NewReader(file)
	for {
		c, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		charcter.WriteRune(c)
		if r, ok := prefixToRuneMap[charcter.String()]; ok {
			fmt.Printf("%c", r)
			charcter.Reset()
		}

	}

}

func main() {

	fileName := os.Args[1]
	frequencyMap := getFrequencyTable(fileName)
	tree := CreateTree(frequencyMap)
	runeToPrefixTable := map[rune]string{}
	prefixToRuneMap := map[string]rune{}
	GetPrefixTable(tree, runeToPrefixTable, "")
	for r, s := range runeToPrefixTable {
		prefixToRuneMap[s] = r
		fmt.Printf("%c  %s\n", r, s)
	}
	encodedFileName := encodeFile(fileName, runeToPrefixTable)
	decodeFile(encodedFileName, prefixToRuneMap)

}
