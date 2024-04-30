package main

func DivisorsForNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	//находим наименьшее положительное число в массиве
    min := numbers[0]
    for _, num := range numbers {
        if num < min {
            min = num
        }
    }

	//cоздаем слайс для хранения общих делителей
    divisors := make([]int, 0)

    //проверяем каждое число от 1 до min
    for i := 2; i <= min; i++ {
        //фaлаг для проверки, является ли число общим делителем
        isCommonDivisor := true

        //проверяем, является ли число делителем для каждого числа в массиве
        for _, num := range numbers {
            if num % i != 0 {
                isCommonDivisor = false
                break
            }
        }

        //если число является общим делителем, добавляем его в слайс
        if isCommonDivisor {
            divisors = append(divisors, i)
        }
    }

    return divisors
}