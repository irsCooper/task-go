package main

import "fmt"

func MultiplicationTable(n int) {
    //выводим заголовок таблицы
    fmt.Printf("%5s", " ")

    for i := 1; i <= n; i++ {
        fmt.Printf("%5d", i)
    }

    fmt.Println()

    //выводим строки таблицы умножения
    for i := 1; i <= n; i++ {
        fmt.Printf("%5d", i)

        for j := 1; j <= n; j++ {
            fmt.Printf("%5d", i*j)
        }

        fmt.Println()
    }
}