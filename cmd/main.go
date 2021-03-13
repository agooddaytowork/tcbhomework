package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {

	f1 := []float64{15, 20, 30, 40, 50}
	// f1 := []float64{}
	result, err := stats.PercentileNearestRank(f1, 102)

	fmt.Println(result, err)
}
