package main

import "fmt"

type rect struct {
	width, height int
}

type point struct {
	x, y int
}

func (r rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return (2 * r.height) + (2 * r.width)
}

func (p point) updateX(a int) {
	p.x = a
	fmt.Println("x inside", p.x)
}

func (p *point) updateY(a int) {
	p.y = a
	fmt.Println("y inside", p.y)
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("perim:", r.perim())
	fmt.Println("area: ", r.area())

	s := rect{width: 20, height: 15}
	st := &s
	fmt.Println("perim:", st.perim())
	fmt.Println("area: ", st.area())

	p1 := point{x: 10, y: 20}
	fmt.Println("x before", p1.x)
	p1.updateX(5)
	fmt.Println("x after", p1.x)
	p2 := &p1
	fmt.Println("x before", p2.x)
	p1.updateX(5)
	fmt.Println("x after", p2.x)

	k1 := point{x: 100, y: 200}
	fmt.Println("y before", k1.y)
	k1.updateY(500)
	fmt.Println("y after", k1.y)

	k := point{x: 150, y: 250}
	k2 := &k

	fmt.Println("y before", k2.y)
	k2.updateY(600)
	fmt.Println("y after", k2.y)
}
