package main

import "fmt"

type geometry interface {
	area() float32
	perim() float32
}

type rect struct {
	len, wid float32
}

func (r rect) area() float32 {
	return r.len * r.wid
}

func (r rect) perim() float32 {
	return 2 * (r.len + r.wid)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * 3.14 * c.radius
}

func show(name string, g geometry) {
	fmt.Printf("area of %v is %v \n", name, g.area())
	fmt.Printf("perim of %v is %v \n", name, g.perim())
}

func main() {
	rec := rect{
		len: 1,
		wid: 2,
	}
	show("rect", rec)

	cir := circle{
		radius: 1,
	}
	show("circle", cir)

	return
}
