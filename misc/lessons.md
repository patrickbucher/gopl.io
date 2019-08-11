# Lessons Learned

## Data Structures

- When composite data types (structs) inside the map need to be updated, store
  pointers, not struct values.

## Concurrency

- When working with `async.WaitGroup`, always call `wg.Add(1)` outside the
  goroutine that defers the call to `wg.Done()`.
- Never write to a closed channel, and never close a channel twice.
- To make sure a mutex is released for all paths, defer the `Unlock()` call to
  the mutex at the beginning of the function.
- A _monitor_ is the arrangement of exported functions accessing an unexported
  variable through a mutex.
- The code section between the mutexe's lock and unlock is called the _critical
  section_ of a concurrent function.

## Benchmarking

- Always run `b.N` (`*benchmark.B`) iterations per benchnark instead of calling
  the function at hand only once.
