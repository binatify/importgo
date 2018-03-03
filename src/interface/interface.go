package main

import (
	"fmt"
	"math"
)

type tmp interface {
	area() float32
}

type geometry interface {
	// area() float32
	tmp
	perim() float32
}

type rect struct {
	len, width float32
}

func (r *rect) area() float32 {
	return r.len * r.width
}

func (r *rect) perim() float32 {
	return 2 * (r.len + r.width)
}

type circle struct {
	radius float32
}

func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func show(name string, params interface{}) {
	switch params.(type) {
	case geometry:
		fmt.Printf("area of %v is %v \n", name, params.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, params.(geometry).perim())
	default:
		fmt.Println("Wrong type!!")
	}
}

func main() {
	r := &rect{
		len:   1,
		width: 2,
	}
	show("rect", r)

	cir := &circle{
		radius: 1,
	}
	show("circle", cir)
}
