package utils

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	
)

func GetPath() string {
	pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
	}
	return pwd
}

func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func ReadFile(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
    if err != nil {
		panic(err)
	}
	return string(dat)
}