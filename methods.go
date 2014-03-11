package main

import "fmt"

type car struct {
    people  int
    speed   int
}

func (c *car) seatPerson() { // Changes c
    c.people++
}

func (c car) seatPeople(count int) { // Doesn't change c
    c.people += count
}

func (c *car) showPeople() int {
    return c.people
}

func main() {

    c := car{people: 1, speed: 20}

    fmt.Println("People:", c.showPeople())
    c.seatPerson()
    fmt.Println("People:", c.showPeople())
    c.seatPeople(3)
    fmt.Println("People:", c.showPeople())
}
