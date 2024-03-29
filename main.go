package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// handle file
	file, err := os.Open("data/weather_stations.csv")
	CheckError(err)
	defer file.Close()

	// create scanner to scan file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// map for storing stations
	stations := make(map[string][]float64)

	// loop through file
	for scanner.Scan() {

		// split line into key and value
		line := scanner.Text()
		contents := strings.Split(line, ";")
		key := contents[0]
		value, err := strconv.ParseFloat(contents[1], 64)
		CheckError(err)

		// add key and value to map
		if _, ok := stations[key]; !ok {
			stations[key] = []float64{value, 1, value, value} // [cum_temp, count, min, max]
		} else {
			stations[key] = append(stations[key], value)
			stations[key][1] += 1
			if stations[key][2] > value {
				stations[key][2] = value
			}
			if stations[key][3] < value {
				stations[key][3] = value
			}
		}
	}

	for key, value := range stations {
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", key, math.Round(value[2]*10)/10, (math.Round(value[0]*10)/10)/(math.Round(value[1]*10)/10), math.Round(value[3]*10)/10)
	}

}
