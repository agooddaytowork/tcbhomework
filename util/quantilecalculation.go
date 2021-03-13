package util

import (
	"errors"
	"math"
	"sort"
)

// Return the CalulatedQuantile and error
// Pool must be sorted in increasing order
// Implementation followed: https://en.wikipedia.org/wiki/Percentile#The_nearest-rank_method
func CalculateQuantileNearestRank(pool []int32, percentile float64) (int32, error) {

	poolLength := len(pool)
	if poolLength == 0 {
		return 0, errors.New("Input Pool must not be empty")
	}

	if percentile > 100.0 || percentile < 0 {
		return 0, errors.New("Percentile must be between 0 - 100")
	}

	sort.Slice(pool, func(i, j int) bool {
		return pool[i] < pool[j]
	})

	if percentile == 100 {
		return pool[poolLength-1], nil
	}

	ith := (percentile / 100.0) * float64(poolLength)
	// Find ordinal ranking
	or := int(math.Ceil(ith))

	// Return the item that is in the place of the ordinal rank
	if or == 0 {
		return pool[0], nil
	}

	return pool[or-1], nil

}
