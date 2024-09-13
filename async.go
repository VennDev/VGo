package vgo

type AsyncResult struct {
	deferred *Deferred
}

/**
 * Async creates a new Deferred instance and runs the callback function in a new goroutine.
 * The Deferred instance is returned as an AsyncResult instance.
 * @param callback The callback function to run asynchronously.
 * @return An AsyncResult instance.
 * @example
 * async := Async(func() interface{} {
 *    return "Hello, World!"
 * })
 */
func Async(callback func() interface{}) *AsyncResult {
	deferred := NewDeferred()
	go func() {
		deferred.callback <- callback
		close(deferred.callback)
	}()
	deferred.Run()
	return &AsyncResult{
		deferred: deferred,
	}
}

/**
 * Await waits for the Deferred instance to finish and returns the result.
 * @param value The Deferred instance to wait for.
 * @return The result of the Deferred instance.
 * @example
 * async := Async(func() interface{} {
 *   return "Hello, World!"
 * })
 * result := Await(async)
 * fmt.Println(result) // Output: Hello, World!
 */
func Await(value interface{}) interface{} {
	switch v := value.(type) {
	case *AsyncResult:
		return <-v.deferred.result
	case *Deferred:
		return <-v.result
	case func() interface{}:
		return v()
	default:
		panic("Await expects an Async instance or a Deferred instance!")
	}
}
