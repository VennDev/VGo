# VGo
The library helps you have the async-await syntax in Go that runs on goroutines.

# Example
```go
package vgo

import (
	"fmt"
	"testing"
	"time"
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

func TestDeferred(t *testing.T) {
	deferred := NewDeferred()
	go func() {
		deferred.callback <- func() interface{} {
			return "Hello, World!"
		}
		deferred.Close()
	}()
	deferred.Run()
	result := <-deferred.result

	fmt.Println(result) // Output: Hello, World!
}

func TestDeferredWithAny(t *testing.T) {
	deferred1 := NewDeferred()
	go func() {
		deferred1.callback <- func() interface{} {
			time.Sleep(1 * time.Second)
			return 10
		}
		deferred1.Close()
	}()
	deferred1.Run()

	deferred2 := NewDeferred()
	go func() {
		deferred2.callback <- func() interface{} {
			return 20
		}
		deferred2.Close()
	}()
	deferred2.Run()

	async := Any(deferred1, deferred2)
	result := Await(async)

	fmt.Println(result) // Output: 20
}

func TestDeferredWithAll(t *testing.T) {
	deferred1 := NewDeferred()
	go func() {
		deferred1.callback <- func() interface{} {
			return 10
		}
		deferred1.Close()
	}()
	deferred1.Run()

	deferred2 := NewDeferred()
	go func() {
		deferred2.callback <- func() interface{} {
			return 20
		}
		deferred2.Close()
	}()
	deferred2.Run()

	async := All(deferred1, deferred2)
	result := Await(async)

	fmt.Println(result) // Output: [10 20]
}
```
