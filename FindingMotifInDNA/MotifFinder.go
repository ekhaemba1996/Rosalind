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

func substringCount(vals []string) []int {
	s := vals[0]
	t := vals[1]
	s_length := len(s)
	t_length := len(t)

	indices := []int{}
	for index := 0; index < s_length - t_length; index++ {
		substring := s[index:index + t_length]
		if substring == t {
			indices = append(indices, index + 1)
		}
	}
	return indices
}

func main() {
	result, err := readStrings("rosalind_subs.txt")
	if err != nil {
        fmt.Println("File reading error", err)
	}
	printStringArray(result)
	fmt.Println("Length", len(result))
	fmt.Println(substringCount(result))
}