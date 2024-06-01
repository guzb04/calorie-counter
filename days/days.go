package days

import (
	"calorie-counter/data"
	"time"
)



func Calculate(days []*data.Day) *data.Average {

	a := data.Average{
		Calories: 0,
		Protein:  0,
		Days:     0,
	}

	for i := 0; i < len(days); i++ {
		a.Calories += float32(days[i].Calories) / float32(len(days))
		a.Protein += float32(days[i].Protein) / float32(len(days))
		a.Days++
	}
	return &a
}

func Add(calories, protein int, days []*data.Day) ([]*data.Day, error) {
	currentTime := time.Now()

	if currentTime.Day() == days[0].Day && currentTime.Month() == time.Month(days[0].Month) && currentTime.Year() == days[0].Year {
		days[0].Calories += calories
		days[0].Protein += protein

		return days, nil 
	} 
		var newDay = data.Day{
			Year:     currentTime.Year(),
			Month:    int(currentTime.Month()),
			Day:      currentTime.Day(),
			Calories: calories,
	}

  newDays := prependDay(days, &newDay)

  return newDays, nil

  
}

func prependDay(jsonDays []*data.Day, newDay *data.Day) []*data.Day{

  newDays := make([]*data.Day, len(jsonDays)+1)

  copy(newDays[1:], jsonDays)
  newDays[0] = newDay

  return newDays


}
