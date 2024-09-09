package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := executionSelectedMethod(getUserInput())
	fmt.Println(result)
}

func getUserInput() ([]float64, string) {
	var operationName string
	var userInputNumbers string
	var arrayOfNumbers []float64

	for {
		fmt.Println("Введите операцию (AVG, SUM, MED): ")
		fmt.Scan(&operationName)
		if operationName != "avg" && operationName != "sum" && operationName != "med" {
			fmt.Println("Вы не ввели нужную операцию, введите заново: ")
			continue
		}
		break
	}

	fmt.Println("Введите числа через запятую: ")
	fmt.Scan(&userInputNumbers)
	userInputNumberWithoutComma := strings.Split(userInputNumbers, ",")
	for _, numberStr := range userInputNumberWithoutComma {
		numberStr = strings.TrimSpace(numberStr)
		num, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			fmt.Println("Ошибка: неверный формат числа", numberStr)
		}
		arrayOfNumbers = append(arrayOfNumbers, num)
	}
	return arrayOfNumbers, operationName
}

func executionSelectedMethod(arrayOfNumbers []float64, operationName string) float64 {
	var result float64
	if operationName == "avg" {
		for _, number := range arrayOfNumbers {
			result += number
		}
		result /= float64(len(arrayOfNumbers))
	}
	if operationName == "sum" {
		for _, number := range arrayOfNumbers {
			result += number
		}

	}
	if operationName == "med" {
		sort.Float64s(arrayOfNumbers)
		n := len(arrayOfNumbers)
		if n%2 == 0 {
			result = (arrayOfNumbers[n/2-1] + arrayOfNumbers[n/2]) / 2
		} else {
			result = arrayOfNumbers[n/2]
		}
	}
	return result
}
