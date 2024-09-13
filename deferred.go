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
 * close(deferred.callback)
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
