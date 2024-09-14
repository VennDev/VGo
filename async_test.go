package vgo

import (
	"fmt"
	"strconv"
	"testing"
)

func doAsync() *Async {
	return &Async{
		Callback: func() interface{} {
			return 30
		},
	}
}

func TestAsync(t *testing.T) {
	async := doAsync()
	result := Await(async)
	fmt.Println("Result async:", strconv.Itoa(result.(int))) // Output: Result async: 30
}

func TestDeferred(t *testing.T) {
	deferred := &Deferred{
		Callback: func() interface{} {
			return 50
		},
	}
	result := deferred.Await()
	fmt.Println("Result deferred:", strconv.Itoa(result.(int))) // Output: Result deferred: 50
}

func TestDeferredAll(t *testing.T) {
	deferred := &Deferred{}
	result := deferred.All(
		func() interface{} {
			return 10
		},
		func() interface{} {
			return 20
		},
		func() interface{} {
			return 30
		},
	).Await()
	fmt.Println("Result deferred all:", result) // Output: Result deferred all: [10 20 30]
}

func TestDeferredAny(t *testing.T) {
	deferred := &Deferred{}
	result := deferred.Any(
		func() interface{} {
			return 10
		},
		func() interface{} {
			return 20
		},
		func() interface{} {
			return 30
		},
	).Await()
	fmt.Println("Result deferred any:", result) // Output: Result deferred any: 10
}
