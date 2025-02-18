package helpers

// Checks if a number is an extreme outlier
func IsExtreme(num int, numList []int) bool {
	if len(numList) < 2 {
		return false
	}

	mean := GetAverage(numList)
	stdDev := GetStandardDeviation(numList)

	lowerThreshold := mean - float64(3)*stdDev
	upperThreshold := mean + float64(3)*stdDev

	return float64(num) < lowerThreshold || float64(num) > upperThreshold
}

// Removes a specific outlier from the list
func RemoveOutlier(numList []int, outlier int) []int {
	for i, n := range numList {
		if n == outlier {
			return append(numList[:i], numList[i+1:]...)
		}
	}
	return numList
}
