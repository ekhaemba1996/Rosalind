package main

import (  
    "fmt"
	"io/ioutil"
	"errors"
	"strings"
)

func readStrings(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("File reading error", err)
        return nil, errors.New("Unable to open file")
	}
	read_string := string(data)
	vals := strings.Split(read_string, "\n")
	return vals, nil
}

func printStringArray(result []string) {
	for _, element := range result {
		fmt.Println(element)
	}
}

func mutationCount(sequences []string) int {
	sequences[0] = sequences[0][:len(sequences[0]) - 1]
	muteCount := 0
	for index, _ := range sequences[0] {
		if sequences[0][index] != sequences[1][index] {
			muteCount++
		}
	}
	return muteCount
}

func main() {
	result, err := readStrings("rosalind_hamm.txt")
	if err != nil {
        fmt.Println("File reading error", err)
	}
	printStringArray(result)
	fmt.Println("Length", len(result))
	fmt.Println(mutationCount(result))
}