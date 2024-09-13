package vgo

type Deferred struct {
	callback chan func() interface{}
	result   chan interface{}
}

func NewDeferred() *Deferred {
	return &Deferred{
		callback: make(chan func() interface{}),
		result:   make(chan interface{}),
	}
}

/**
 * Run starts the Deferred instance.
 * @example
 * deferred := NewDeferred()
 * go func() {
 *  deferred.callback <- func() interface{} {
 *  return "Hello, World!"
 * }
 * deferred.Close()
 * }()
 * deferred.Run()
 * result := <-deferred.result
 * fmt.Println(result) // Output: Hello, World!
 */
func (p *Deferred) Run() {
	go func() {
		for fn := range p.callback {
			p.result <- fn()
		}
		close(p.result)
	}()
}

/**
 * Close closes the Deferred instance.
 */
func (p *Deferred) Close() {
	close(p.callback)
}

/*
 * Create one new Deferred instance for each value in the values array.
 * The first Deferred instance to finish will return its result.
 * @param values The Deferred instances to wait for.
 * @return A new Deferred instance.
 * @example
 * deferred1 := NewDeferred(func() interface{} {
 *  return 10
 * })
 * deferred2 := NewDeferred(func() interface{} {
 *  return 20
 * })
 * async := Any(deferred1, deferred2)
 * result := Await(async)
 * fmt.Println(result) // Output: 10
 */
func Any(values ...*Deferred) *Deferred {
	deferred := NewDeferred()
	go func() {
		results := make(chan interface{}, len(values))
		for _, value := range values {
			go func(d *Deferred) {
				result := <-d.result
				results <- result
			}(value)
		}
		for range values {
			result := <-results
			if result != nil {
				deferred.result <- result
				break
			}
		}
		close(results)
	}()
	return deferred
}

/**
 * Create one new Deferred instance for each value in the values array.
 * All Deferred instances must finish before the new Deferred instance returns.
 * @param values The Deferred instances to wait for.
 * @return A new Deferred instance.
 * @example
 * deferred1 := NewDeferred(func() interface{} {
 *  return 10
 * })
 * deferred2 := NewDeferred(func() interface{} {
 *  return 20
 * })
 * async := All(deferred1, deferred2)
 * result := Await(async)
 * fmt.Println(result) // Output: [10, 20]
 */
func All(values ...*Deferred) *Deferred {
	deferred := NewDeferred()
	go func() {
		results := make([]interface{}, len(values))
		for i, value := range values {
			results[i] = <-value.result
		}
		deferred.result <- results
		close(deferred.result)
	}()
	return deferred
}
