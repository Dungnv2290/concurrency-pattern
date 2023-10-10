package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5
	startTime := time.Now()

	for i := 0; i < N; i++ {
		result := addFoo(addQuote(square(multiplyTwo(i))))
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Printf("Elapsed time without concurrency: %s\n", time.Since(startTime))

	outC := NewPipeline(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < N; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) {
			return multiplyTwo(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return square(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addQuote(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addFoo(in.(string)), nil
		}).
		Merge()

	startTimeC := time.Now()
	for result := range outC {
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Printf("Elapsed time with concurrency: %s\n", time.Since(startTimeC))
}

func multiplyTwo(v int) int {
	time.Sleep(1 * time.Second)
	return v * 2
}

func square(v int) int {
	time.Sleep(2 * time.Second)
	return v * v
}

func addQuote(v int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("'%d'", v)
}

func addFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
