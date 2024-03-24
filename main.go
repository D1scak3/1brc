package main

import (
	"bufio"
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// open file
	file, err := os.Open("data/weather_stations.csv")
	CheckError(err)
	defer file.Close()

	// create scanner to scan file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// loop through file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
