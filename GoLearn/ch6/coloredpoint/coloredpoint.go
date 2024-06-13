package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

//type ColoredPoint struct {
//	Point
//	Color color.RGBA
//}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"

	//red := color.RGBA{255, 0, 0, 255}
	//blue := color.RGBA{0, 0, 255, 255}
	//var p = ColoredPoint{Point{1, 1}, red}
	//var q = ColoredPoint{Point{5, 4}, blue}
	//// NOTE: 可以把ColoredPoint类型当作接收器来调用Point里的方法，即使ColoredPoint里没有声明这些方法
	//fmt.Println(p.Distance(q.Point)) // "5"
	//p.ScaleBy(2)
	//q.ScaleBy(2)
	//fmt.Println(p.Distance(q.Point)) // "10"

	//p := ColoredPoint{&Point{1, 1}, red}
	//q := ColoredPoint{&Point{5, 4}, blue}
	//fmt.Println(p.Distance(*q.Point)) // "5"
	//q.Point = p.Point                 // p and q now share the same Point
	//p.ScaleBy(2)
	//fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"

	p := Point{1, 2}
	q := Point{4, 6}
	// distanceFromP叫方法值，本质是一个函数，var distanceFromP func(q Point) float64 = p. Distance
	//distanceFromP := p.Distance        // method value
	//fmt.Println(distanceFromP(q))      // "5"
	//var origin Point                   // {0, 0}
	//fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)
	//
	//scaleP := p.ScaleBy // method value
	//scaleP(2)           // p becomes (2, 4)
	//scaleP(3)           //      then (6, 12)
	//scaleP(10)          //      then (60, 120)

	// distance是一个方法表达式，var distance func(Point, Point) float64 = Point. Distance
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
}
