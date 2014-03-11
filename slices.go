package main

import "fmt"

func main() {

    s := make([]string, 3)
    fmt.Println("emp:", s)
    fmt.Println("len(s):", len(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len(s):", len(s))

    // Append
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Copy
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Slice
    l := s[2:4]
    fmt.Println("sl1:", l)
    l = s[:5]
    fmt.Println("sl2:", l)
    l = s[2:]
    fmt.Println("sl3:", l)

    // Declare and initialize a slice in a single line
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // 2D slices
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d:", twoD)
}
