package vgo

import (
	"fmt"
	"testing"
)

func doAsync() *AsyncResult {
	return Async(func() interface{} {
		return 10 + 20
	})
}

func TestAsync(t *testing.T) {
	async := doAsync()
	result := Await(async)

	fmt.Println(result) // Output: 30
}

func doAwait() *AsyncResult {
	return Async(func() interface{} {
		result := Await(doAsync())
		return result.(int) + 30
	})
}

func TestAwait(t *testing.T) {
	async := doAwait()
	result := Await(async)

	fmt.Println(result) // Output: 60
}
