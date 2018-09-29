package intset

import (
	"math/rand"
	"testing"
	"time"
)

const testSize = 10

var random = rand.New(rand.NewSource(time.Now().Unix()))

func TestHas(t *testing.T) {
	var set IntSet
	numbers := randomNumbers(1, 100, testSize)
	for _, n := range numbers {
		set.Add(n)
	}
	for _, n := range numbers {
		if !set.Has(n) {
			t.Errorf("%d missing in set\n", n)
		}
	}
}

func TestUnionWith(t *testing.T) {
	a, b := &IntSet{}, &IntSet{}
	numbersA := randomNumbers(100, 200, testSize)
	numbersB := randomNumbers(150, 250, testSize)
	for i := 0; i < testSize; i++ {
		a.Add(numbersA[i])
		b.Add(numbersB[i])
	}
	a.UnionWith(b)
	for i := 0; i < testSize; i++ {
		if !a.Has(numbersA[i]) {
			t.Errorf("%d missing in union set\n", numbersA[i])
		}
		if !a.Has(numbersB[i]) {
			t.Errorf("%d missing in union set\n", numbersB[i])
		}
	}
}

func TestLen(t *testing.T) {
	var set IntSet
	assertLen(t, 0, set)
	var expected int
	for _, n := range randomNumbers(1, 100, testSize) {
		if !set.Has(n) {
			expected++
			set.Add(n)
		}
	}
	assertLen(t, expected, set)
}

func TestRemove(t *testing.T) {
	var set IntSet
	numbers := randomNumbers(1, 100, testSize)
	for _, n := range numbers {
		set.Add(n)
	}
	for _, n := range numbers {
		if set.Has(n) {
			set.Remove(n)
			if set.Has(n) {
				t.Errorf("%d was supposed to be removed\n", n)
			}
		}
	}
	if set.Len() > 0 {
		t.Errorf("set was not emptied as expected\n")
	}
}

func TestClear(t *testing.T) {
	var set IntSet
	numbers := randomNumbers(1, 100, testSize)
	for _, n := range numbers {
		set.Add(n)
	}
	set.Clear()
	if set.Len() > 0 {
		t.Errorf("set was not cleared\n")
	}
}

func TestCopy(t *testing.T) {
	var set IntSet
	numbers := randomNumbers(1, 100, testSize)
	for _, n := range numbers {
		set.Add(n)
	}
	cpy := set.Copy()
	if set.Len() != cpy.Len() {
		t.Errorf("copied set should have length %d, but has %d\n",
			set.Len(), cpy.Len())
	}
	set.Add(999) // must not affect original
	if cpy.Has(999) {
		t.Error("modification on original affected copy")
	}
	for _, n := range numbers {
		if !cpy.Has(n) {
			t.Errorf("number %d missing in copied set\n", n)
		}
	}
}

func assertLen(t *testing.T, expected int, set IntSet) {
	if actual := set.Len(); actual != expected {
		t.Errorf("expected set len %d, but was %d\n", expected, actual)
	}
}

func randomNumbers(min, max, n int) []int {
	var numbers []int
	for i := 0; i < n; i++ {
		numbers = append(numbers, random.Intn(max-min+1)+min)
	}
	return numbers
}
