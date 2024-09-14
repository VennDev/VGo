package vgo

type Async struct {
	callback func() interface{}
}

/**
 * Async creates a new Deferred instance and runs the callback function in a new goroutine.
 * @return *Derrered The Deferred instance.
 */
func (p *Async) Run() *Deferred {
	return &Deferred{
		callback: p.callback,
	}
}

/**
 * Await waits for the Deferred instance to finish and returns the result.
 * @param value The Deferred instance to wait for.
 * @return interface{} The result of the Deferred instance.
 */
func Await(value interface{}) interface{} {
	switch v := value.(type) {
	case *Async:
		return v.Run().Run().Await()
	case *Deferred:
		return v.Run().Await()
	case func() interface{}:
		// Create a new Async instance and run the callback function.
		return (&Async{callback: v}).Run().Run().Await()
	default:
		panic("Await expects an Async function, callback or a Deferred instance!")
	}
}
