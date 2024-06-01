package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Average struct {
	Calories float32
	Protein  float32
	Days     int
}

type Day struct {
	Year     int
	Month    int
	Day      int
	Calories int
	Protein  int
}

const JSONPATH string = "./data.json"

func GetData() ([]*Day, error) {

	jsonBlob, err := os.ReadFile(JSONPATH)
	if err != nil {
    fmt.Printf("error reading json file: %v", err)
		return nil, err
	}
	var data []*Day

	err = json.Unmarshal(jsonBlob, &data)
	if err != nil {
    fmt.Printf("error unmarshalling json file: %v", err)
		return nil, err
	}

	return data, nil
}

func WriteData(data []*Day) error {

	if data == nil {
    _, err := os.Create(JSONPATH)
		if err != nil {
      fmt.Printf("error writing file: %v\n", err)
			return err
		}
	}

	jsonString, err := json.Marshal(&data)
	if err != nil {
    fmt.Printf("error marshalling data: %v", err)
		return err
	}

	err = os.WriteFile(JSONPATH, jsonString, 0644)
	if err != nil {
    fmt.Printf("error writing file: %v", err)
		return err
	}

	return nil
}


