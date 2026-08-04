package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find file") // 1000 means the balance
	}

	valueText := string(data)
	balance, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}

	return balance, nil
}

func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
