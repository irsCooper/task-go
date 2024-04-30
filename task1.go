package main

import (
	"fmt"
)

func GetComputerString(num int) string {
	if num < 0 {
		return "Компьютеров не может быть меньше 0"
	}
    
	//преобразуем число в строку
	numStr := fmt.Sprintf("%d", num)
	length := len(numStr)

	var lastTwoDigits int

	//в зависимости от количества цифр в цисле назначим переменную
	if length > 2 {
		lastTwoDigits = num % 100
	} else {
		lastTwoDigits = num
	}

	lastDigit := num % 10

	if lastTwoDigits >= 11 && lastTwoDigits <= 19 {
		return fmt.Sprintf("%d компьютеров", num)
	}

	switch lastDigit {
	case 1:
		return fmt.Sprintf("%d компьютер", num)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", num)
	default:
		return fmt.Sprintf("%d компьютеров", num)
	}
}