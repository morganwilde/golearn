package main

import "fmt"
import "math"

type geometry interface {
    area()  float64
    perim() float64
}

type square struct {
    width, height float64
}
type circle struct {
    radius float64
}

// Implement "geometry" interface for square
func (s square) area() float64 {
    return s.width * s.height
}
func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

// Implement "geometry" interface for circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// A generic measure, works on all type structs that implement interface
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {

    s := square{width: 5, height: 5}
    c := circle{radius: 3}

    measure(s)
    measure(c)
}
