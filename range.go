package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    // range nums returns "key" and "value" pairs
    // _ - the blank identifier
    for _, val := range nums {
        sum += val
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }

}
