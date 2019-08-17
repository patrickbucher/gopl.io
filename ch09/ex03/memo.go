// Package memo provides a concurrency-safe
// memoization of a function of type Func.
package memo

import (
	"errors"
)

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	done     <-chan struct{} // the client wants to cancel the operation
}

// A Memo caches the results of calling a Func.
type Memo struct{ requests chan request }

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{make(chan request)}
	go memo.server(f)
	return memo
}

// Get is concurrency-safe.
func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

type entry struct {
	res    result
	ready  chan struct{} // closed when res is ready
	done   <-chan struct{}
	cancel chan struct{}
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{
				ready:  make(chan struct{}),
				done:   req.done,
				cancel: make(chan struct{}),
			}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key, done)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function.
	finished := make(chan result)
	go func() {
		val, err := f(key, e.cancel)
		finished <- result{val, err}
	}()
	select {
	case r := <-finished:
		e.res.value = r.value
		e.res.err = r.err
		break
	case <-e.done:
		close(e.cancel)
		e.res = result{
			value: nil,
			err:   errors.New("operation cancelled"),
		}
		break
	}
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
