package main

import "fmt"

func main() {
    i := 1 // This infers int type, same as var i int = 1
    for i < 7 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    for {
        fmt.Println("loop")
        break
    }
}
