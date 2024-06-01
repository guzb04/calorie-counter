package controller

import (
	"calorie-counter/data"
	"calorie-counter/days"
	"fmt"
)

func GetDays(n int) ([]*data.Day, error) {
	jsonDays, err := data.GetData()

	numberDays := min(n, len(jsonDays))

	if n > len(jsonDays) {
		fmt.Printf("number input %v larger than length of days in data.json (%v)\n", n, len(jsonDays))
	}

	if err != nil {
		return nil, err
	}

  if numberDays == 0 {
    var lastDay  [1]*data.Day

    lastDay[0] = jsonDays[0]

    return lastDay[:], nil


  }

	return jsonDays[:numberDays], nil
}

func GetAverage(n int) (*data.Average, error) {
  jsonDays, err := data.GetData()


	numberDays := min(n, len(jsonDays))

	if n > len(jsonDays) {
		fmt.Printf("number input %v larger than length of days in data.json (%v)\n", n, len(jsonDays))
	}

	if err != nil {
		return nil, err
	}

  average := days.Calculate(jsonDays[:numberDays])

  return average, nil

}
