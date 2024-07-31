package main

import (
	"fmt"
	"reflect"
)

func test1() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)
	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"

}

func test2() {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"
	d.Set(reflect.ValueOf(4))
	fmt.Println(x) // "4"
}

func test3() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)               // OK, x = 2
	rx.Set(reflect.ValueOf(3)) // OK, x = 3
	// rx.SetString("hello")            // panic: string is not assignable to int
	// rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2)               // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3)) // OK, y = int(3)
	// ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"

}

func main() {
	test1()
	test2()
	test3()
}
