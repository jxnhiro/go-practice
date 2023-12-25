package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (value float64, err error){
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000.0, errors.New("failed to find file")
	}

	valueText := string(data)
	value, _ = strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000.0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}