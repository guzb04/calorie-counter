package main

import (
	"calorie-counter/controller"
	"calorie-counter/data"
	"calorie-counter/days"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	protein := addCmd.Int("protein", 0, "grams of protein")
	calories := addCmd.Int("calories", 0, "amount of calories")
	daysAverage := getCmd.Int("days", 0, "number of days")

	if len(os.Args) < 2 {
		fmt.Printf("Usage: calorie-counter <add|get> [options]")
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		jsonDays, err := data.GetData()

		if err != nil {

			newDay := data.Day{
				Calories: *calories,
				Protein:  *protein,
				Day:      time.Now().Day(),
				Month:    int(time.Now().Month()),
				Year:     time.Now().Year(),
			}
			var newDayArray [1]*data.Day
			newDayArray[0] = &newDay

			err1 := data.WriteData(newDayArray[:])

			if err1 != nil {
				fmt.Printf("%v", err1)
				os.Exit(1)
			}

		}

		modifiedDays, err := days.Add(*calories, *protein, jsonDays)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		data.WriteData(modifiedDays)

	case "get":
    getCmd.Parse(os.Args[2:])

		newAverage, err := controller.GetAverage(*daysAverage)

		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		fmt.Printf(" Calories: %v \n Protein: %v \n Days for average: %v", newAverage.Calories, newAverage.Protein, newAverage.Days)
	}
}
