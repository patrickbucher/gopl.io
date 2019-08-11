# Lessons Learned

## Data Structures

- When composite data types (structs) inside the map need to be updated, store
  pointers, not struct values.

## Concurrency

- When working with `async.WaitGroup`, always call `wg.Add(1)` outside the
  goroutine that defers the call to `wg.Done()`.
- Never write to a closed channel, and never close a channel twice.

## Benchmarking

- Always run `b.N` (`*benchmark.B`) iterations per benchnark instead of calling
  the function at hand only once.
