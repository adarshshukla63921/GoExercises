package main

import "fmt"

func NewCounter(name string) (increment func(), getValue func() int) {
	type Counter struct {
		name  string
		value int
	}
	// No need to create pointer here.
	// Closure already forces 'counter' to be captured by reference.
	counter := Counter{name, 0}
	return func() {
			counter.value++
		}, 
		func() int {
			return counter.value
		}
}

func main() {
	increment, getValue := NewCounter("FirstCounter")
	increment() // 1
	fmt.Println(getValue())
	increment() // 2
	fmt.Println(getValue())
	increment() // 3
	fmt.Println(getValue())
}