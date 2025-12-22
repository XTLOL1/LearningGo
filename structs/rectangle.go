package main

import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	sizeY float64
	sizeX float64	
}

func (r Rectangle) area() float64 {
	return r.sizeY * r.sizeX
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	height float64
	base float64
}

func (t Triangle) area() float64 {
	return (t.height * t.base) / 2
}
