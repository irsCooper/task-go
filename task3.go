package main

func GetPrimes(min, max int) []int {
    //создаем слайс для хранения простых чисел
    primes := make([]int, 0)

    //проходим по всем числам в диапазоне
    for num := min; num <= max; num++ {
        // Пропускаем числа меньше 2, так как они не являются простыми
        if num < 2 {
            continue
        }

        //проверяем, является ли текущее число простым
        isPrime := true
        for i := 2; i*i <= num; i++ {
            if num % i == 0 {
                isPrime = false
                break
            }
        }

        //если число простое, добавляем его в слайс
        if isPrime {
            primes = append(primes, num)
        }
    }

    return primes
}
