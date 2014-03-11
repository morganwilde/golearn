package main

import "fmt"

func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n - 1)
}

func fact() func() int {
    i := 0
    return func() int {
        i += 1
        return factorial(i)
    }
}

func main() {

    fact1 := factorial(1)
    fact2 := factorial(2)
    fact5 := factorial(5)

    fmt.Println("fact1:", fact1)
    fmt.Println("fact2:", fact2)
    fmt.Println("fact5:", fact5)

    f := fact()
    for i := 1; i < 10; i++ {
        fmt.Println("f:", f())
    }
}
