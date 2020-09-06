package main

import "fmt"

type triangle struct{
    base float64
    height float64
}

type square struct{
    sideLength float64
}

type shape interface{
    getArea() float64
}

func (t triangle) getArea() float64 {
    return (t.base * t.height * .5)
}

func (s square) getArea() float64 {
    return (s.sideLength * s.sideLength)
}


func main() {
    tri := triangle{
        base: 3.0,
        height: 4.0,
    }

    sq := square{
        sideLength: 3.0,
    }

    printArea(tri)
    printArea(sq)
}


func printArea(s shape) {
    fmt.Println(s.getArea())
}

