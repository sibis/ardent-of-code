package main

import (
	"fmt"
	"strconv"

	"github.com/sibis/ardent-of-code/utils"
)

func calculateProduct1(data []string) int {
	for i, item := range data {
		for j := i+1; j < len(data); j++ {
			a, _ := strconv.Atoi(item)
			b, _ := strconv.Atoi(data[j])
			//fmt.Println(a+b)
			if a + b == 2020 {
				return a * b
			}
		} 
	}
	return 0
}

func calculateProduct2(data []string) int {
	for i := 0; i < len(data)-3 ; i++ {
		for j := i+1; j < len(data)-3; j++ {
			for k := j+1; j < len(data)-1; k++ {
				fmt.Println(len(data), " - ", k, " - ", j ," - ", i)
				a, _ := strconv.Atoi(data[i])
				b, _ := strconv.Atoi(data[j])
				c, _ := strconv.Atoi(data[k])
				//fmt.Println(a+b)
				if a + b + c == 2020 {
					return a * b * c
				}
			}
		} 
	}
	return 0
}

func main() {
	pwd := utils.GetPath()
	fileContent, err := utils.ReadLines(pwd+"/day-01/input.txt")
	if err != nil {
		panic(err)
	}

	productResult1 := calculateProduct1(fileContent)
	fmt.Println("Result that sums 2020 : ", productResult1)

	productResult2 := calculateProduct2(fileContent)
	fmt.Println("Result that sums 2020 : ", productResult2)
}