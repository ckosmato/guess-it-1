package helpers

import (
	"math"
	"sort"
)

// Calculates the average of a list of integers
func GetAverage(numList []int) float64 {
	if len(numList) == 0 {
		return 0
	}

	sum := 0
	for _, n := range numList {
		sum += n
	}
	return float64(sum) / float64(len(numList))
}

// Calculates the median of a list of integers
func GetMedian(numList []int) int {
	tempList := append([]int{}, numList...)
	sort.Ints(tempList)
	if len(tempList)%2 == 0 {
		return int(math.Round((float64(tempList[len(tempList)/2-1]) + float64(tempList[len(tempList)/2])) / 2))
	}
	return tempList[len(tempList)/2]
}

// Calculates the standard deviation of a list of integers
func GetStandardDeviation(numList []int) float64 {
	variance := GetVariance(numList) // Calculate variance
	return math.Sqrt(variance)
}

// Calculates the variance of a list of integers
func GetVariance(numList []int) float64 {
	average := GetAverage(numList)
	var varianceSum float64
	for _, n := range numList {
		diff := float64(n) - average
		varianceSum += diff * diff
	}
	return varianceSum / float64(len(numList))
}
