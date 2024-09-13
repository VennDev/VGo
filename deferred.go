package vgo

type Deferred struct {
	callback func() interface{}
}

type ResultDeferred struct {
	result chan interface{}
}

/**
 * Deferred creates a new Deferred instance and runs the callback function in a new goroutine.
 * @return *Derrered The Deferred instance.
 */
func (p *Deferred) Run() *ResultDeferred {
	result := make(chan interface{})
	go func() {
		result <- p.callback()
	}()
	return &ResultDeferred{result}
}

/**
 * Await waits for the Deferred instance to finish and returns the result.
 * @return interface{} The result of the Deferred instance.
 */
func (p *Deferred) Await() interface{} {
	return p.Run().Await()
}

/**
 * Await waits for the Deferred instance to finish and returns the result.
 * @return interface{} The result of the Deferred instance.
 * @deprecated Use Run().Await() instead.
 */
func (p *ResultDeferred) Await() interface{} {
	return <-p.result
}

/**
 * Then waits for the Deferred instance to finish and returns the result.
 * @param callbacks The callback functions to run after the Deferred instance is finished.
 * @return interface{} The result of the Deferred instance.
 */
func (p *Deferred) All(callbacks ...func() interface{}) *ResultDeferred {
	result := make(chan interface{})
	go func() {
		var results []interface{}
		for _, callback := range callbacks {
			results = append(results, callback())
		}
		result <- results
	}()
	return &ResultDeferred{result}
}

/**
 * Any waits for the Deferred instance to finish and returns the result.
 * @param callbacks The callback functions to run after the Deferred instance is finished.
 * @return interface{} The result of the Deferred instance.
 */
func (p *Deferred) Any(callbacks ...func() interface{}) *ResultDeferred {
	result := make(chan interface{})
	go func() {
		for _, callback := range callbacks {
			if callback() != nil {
				result <- callback()
				return
			}
		}
		result <- nil
	}()
	return &ResultDeferred{result}
}
