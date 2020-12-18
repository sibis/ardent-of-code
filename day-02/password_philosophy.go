package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/sibis/ardent-of-code/utils"
)

// part 1 
func checkPasswordValidity1(word, letter string, min, max int) bool {
	curr := 0
	for _, ch := range word {
		if string(ch) == letter {
			curr++
			if curr > max { return false } 
		}
	}
	if curr < min {
		return false
	}
	return true
}

// part 2
func checkPasswordValidity2(word, letter string, first, second int) bool {
	if letter == string(word[first]) && letter == string(word[second]) {
		return false
	} else if letter == string(word[first]) || letter == string(word[second]) {
		return true
	}
	
	return false
}

func main() {
	var result int 
	var result2 int 
	pwd := utils.GetPath()
	fileContent, err := utils.ReadLines(pwd+"/day-02/input.txt")
	if err != nil {
		panic(err)
	} 
	for _, line := range fileContent {
		left := strings.Split(line, ":")
		// fmt.Println(left[0])
		words := strings.Split(left[0], " ")
		limit := strings.Split(words[0], "-")
		min, _ := strconv.Atoi(limit[0])
		max, _ := strconv.Atoi(limit[1])

		isValid := checkPasswordValidity1(left[1][1:], words[1], min, max)
		if isValid {
			result++
		}

		isValid2 := checkPasswordValidity2(left[1][1:], words[1], min-1, max-1)
		if isValid2 {
			result2++
		}
	}
	fmt.Println("Number of valid passwords - part1! : ", result)
	fmt.Println("Number of valid passwords - part2! : ", result2)
}