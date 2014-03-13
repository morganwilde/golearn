package main

import "fmt"

func f(from string) {
    var i int64 = 0

    for i = 0; i < 40; i++ {
        fmt.Println(from, "fib(", i, "):", fib(i))
    }
}

func fib(n int64) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main() {

    go f("|||")

    go f("---")

    go func(msg string) {
        fmt.Println(msg)
    }("this should be three")

    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
