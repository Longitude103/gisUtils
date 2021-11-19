package gisUtils

import (
	"errors"
	"math"
)

// InverseDW is a function that takes in a slice of values and returns the weighted
// 		distribution of each value in a slice of the same length.
//		This is also protected with a division by zero or negative with and error return
// This function uses the "Shepard's" method to determine the values and is a simple appoach. It uses the
// exponent value of p = 2 which is the defacto standard of this equation. Reference: https://en.wikipedia.org/wiki/Inverse_distance_weighting
func InverseDW(dist []float64) ([]float64, error) {
	totalDist := 0.0
	for _, v := range dist {
		totalDist += math.Pow(v, -2)
	}

	var weightResult []float64
	if totalDist > 0 {
		for _, v := range dist {
			weightResult = append(weightResult, math.Pow(v, -2)/totalDist)
		}
	} else {
		return nil, errors.New("division by zero or negative")
	}

	finalResult := normalize(weightResult)
	return finalResult, nil
}

// normalize function reduces the decimals to four and then normalizes the total value to 1
func normalize(v []float64) []float64 {
	// round first

	var interVals []float64
	var interSum float64
	for _, val := range v {
		value := math.Round(val*1000) / 1000
		interVals = append(interVals, value)
		interSum += value
	}

	normalizer := 1 / interSum

	var finalVals []float64
	for _, val := range interVals {
		calc := val * normalizer
		finalVals = append(finalVals, calc)
	}

	return finalVals
}

// findMaxAndMin function returns the min and max of an slice of float64s
func findMaxAndMin(data []float64) (min float64, max float64) {
	min = data[0]
	max = data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
