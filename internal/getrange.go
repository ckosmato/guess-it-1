package internal

import (
	"bufio"
	"fmt"
	"guessit/internal/helpers"
	"os"
	"strconv"
)

func GetRange() {

	var numList []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("invalid input")
			break
		}

		numList = append(numList, num)

		if helpers.IsExtreme(num, numList) {
			numList = helpers.RemoveOutlier(numList, num)
		}

		if len(numList) == 0 {
			fmt.Printf("0 0\n")
		}

		low, up := helpers.GetMedian(numList)-int(helpers.GetStandardDeviation(numList)), helpers.GetMedian(numList)+int(helpers.GetStandardDeviation(numList))
		fmt.Printf("%d %d\n", low, up)
	}
}
