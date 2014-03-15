package main

import "fmt"

func genNaturals(from, to int) []int {
    // Generates a slice of all natural numbers in [from, to)
    count := to - from
    naturals := make([]int, count)

    num := from
    for i := 0; i < count; i++ {
        naturals[i] = num
        num++
    }

    return naturals
}

func main() {

    fmt.Println("naturals(2, 10):", genNaturals(2, 1000001))
}
