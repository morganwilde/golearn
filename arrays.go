package main

import "fmt"
import "math"

func main() {

    var a [5]int // initializes to 0 for each element
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    var matrix [2][3]int
    for i := 2; i < 4; i++ {
        for j := 0; j < 3; j++ {
            matrix[i-2][j] = int(math.Pow(float64(i), float64(j)))
        }
    }
    fmt.Println("matrix(2x3):", matrix)
}
