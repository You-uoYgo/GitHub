package main

import std "fmt"

func main() {
	A()
	B()
	C()
}
func A() {
	std.Println("Func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			std.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}
func C() {
	std.Println("Fun C")
}
