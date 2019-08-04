package arrsum

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// SumConcurrently sums up the numbers array using p goroutines.
// The function panics if p is bigger than len(numbers).
func SumConcurrently(numbers []int, p int) int {
	if p > len(numbers) || p < 1 {
		msg := fmt.Sprintf("cannot sum up %d numbers using %d goroutines", len(numbers), p)
		panic(msg)
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	bounds := calcLimits(len(numbers), p)

	for i := 0; i < p; i++ {
		wg.Add(1)
		go func(numbers []int, l limits, ch chan<- int) {
			defer wg.Done()
			var sum int
			for i := l.lower; i <= l.upper; i++ {
				sum += numbers[i]
			}
			ch <- sum
		}(numbers, bounds[i], ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for partialSum := range ch {
		sum += partialSum
	}

	return sum
}

type limits struct {
	lower int
	upper int
}

func (l limits) String() string {
	return fmt.Sprintf("[%d,%d]", l.lower, l.upper)
}

func calcLimits(size, n int) []limits {
	limits := make([]limits, n)
	upper := make([]float64, n)
	delta := float64(size) / float64(n)
	lower := 0.0
	for i := 0; i < n; i++ {
		upper[i] = lower + delta - 1
		lower = upper[i] + 1.0
	}
	lower = 0.0
	for i := 0; i < len(upper); i++ {
		limits[i].lower = int(math.Round(lower))
		limits[i].upper = int(math.Round(upper[i]))
		lower = float64(limits[i].upper) + 1
	}
	return limits
}

// Sum sums up the numbers.
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// RandonNumbers generates n pseudo-random numbers in the range [min,max].
func RandomNumbers(n, min, max int) []int {
	rand.Seed(int64(time.Now().Nanosecond()))
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = RandomNumber(min, max)
	}
	return numbers
}

// RandomNumber generates a pseudo-random number in the range [min,max].
func RandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
