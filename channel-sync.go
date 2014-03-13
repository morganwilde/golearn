package main

import "fmt"
//import "time"

func total(in chan int, out chan int) {
    res := 0
    for iter := range in {
        res += iter
    }
    out <- res // sends back the result
}

func main() {

    ch := make(chan int)
    rch := make(chan int)
    
    go total(ch, rch)

    ch <- 1
    ch <- 2
    ch <- 3

    close(ch) // ends the loop in total()
    result := <- rch
    fmt.Println("Total is", result)

}
