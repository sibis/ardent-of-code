package main

import (
	"fmt"

	"github.com/sibis/ardent-of-code/utils"
)

func calculateTrees(lines []string, x, y, pos, currLine, currTrees int) int {
	content := lines[currLine]
	if string(content[pos]) == "#" {
		currTrees++
	}
	if currLine < len(lines)-1 {
		pos = pos + x
		if pos >= len(content) {
			pos = pos - len(content)
		}
		return calculateTrees(lines, x, y, pos, currLine+y, currTrees)
	}
	return currTrees
}

func main() {
	pwd := utils.GetPath()
	fileContent, err := utils.ReadLines(pwd + "/day-03/input1.txt")
	if err != nil {
		panic(err)
	}

	res := calculateTrees(fileContent, 3, 1, 0, 0, 0)
	fmt.Println("Trees found so far : ", res)

	xData := []int{1, 3, 5, 7, 1}
	yData := []int{1, 1, 1, 1, 2}

	var mux = 1

	for i := 0; i < len(xData); i++ {
		mux *= calculateTrees(fileContent, xData[i], yData[i], 0, 0, 0)
	}
	fmt.Println("Trees found after multipying as per task 2 : ", mux)
}
