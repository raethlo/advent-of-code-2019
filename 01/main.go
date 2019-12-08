package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func CalculateFuel(weight float64) float64 {
	return math.Floor(weight/3) - 2
}

func main() {
	filepath := "/01/input.txt"

	file, _ := os.Open("." + filepath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fuel float64
	//module_weights := make
	for scanner.Scan() {
		current, _ := strconv.ParseFloat(scanner.Text(), 64)
		log.Printf("Total=%.2f Row=%.2f", fuel, current)
		fuel += CalculateFuel(current)
	}

	log.Printf("Fuel needed: %.2f", fuel)
}
