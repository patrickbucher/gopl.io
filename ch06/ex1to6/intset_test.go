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

func TestAddAll(t *testing.T) {
	var set IntSet
	numbers := randomNumbers(1, 100, testSize)
	set.AddAll(numbers...)
	for _, n := range numbers {
		if !set.Has(n) {
			t.Errorf("value %d is missing in set\n", n)
		}
	}
}

func TestIntersect(t *testing.T) {
	a, b := &IntSet{}, &IntSet{}
	a.AddAll(5, 10, 15, 20, 25, 30)
	b.AddAll(10, 20, 30, 40, 50, 60)
	positive := []int{10, 20, 30}
	negative := []int{5, 15, 25, 40, 50, 60}
	a.IntersectWith(b)
	testHas(a, negative, positive, "intersection", t)
}

func TestDifference(t *testing.T) {
	a, b := &IntSet{}, &IntSet{}
	a.AddAll(3, 6, 9, 12, 15, 18, 21)
	b.AddAll(6, 12, 18, 24, 30, 36)
	positive := []int{3, 9, 15, 21}
	negative := []int{6, 12, 18, 24, 30, 36}
	a.DifferenceWith(b)
	testHas(a, negative, positive, "difference", t)
}

func TestSymmetricDifference(t *testing.T) {
	a, b := &IntSet{}, &IntSet{}
	a.AddAll(2, 4, 6, 8, 10, 12, 14)
	b.AddAll(3, 6, 9, 12, 15, 18)
	positive := []int{2, 3, 4, 8, 9, 10, 14, 15, 18}
	negative := []int{6, 12}
	a.SymmetricDifference(b)
	testHas(a, negative, positive, "symmetric difference", t)
}

func TestElem(t *testing.T) {
	var s IntSet
	numbers := []int{2, 4, 8, 16, 32, 64, 128, 256}
	for _, n := range numbers {
		s.Add(n)
	}
	elems := s.Elem()
	if len(elems) != s.Len() {
		t.Errorf("set and element slide have different lengths: %d != %d\n",
			len(elems), s.Len())
	}
	for _, e := range elems {
		if !s.Has(e) {
			t.Errorf("%d must not be in the set\n", e)
		}
	}
}

func testHas(set *IntSet, negative, positive []int, desc string, t *testing.T) {
	for i := range positive {
		if !set.Has(positive[i]) {
			t.Errorf("%d missing in %s\n", positive[i], desc)
		}
	}
	for i := range negative {
		if set.Has(negative[i]) {
			t.Errorf("%d too much in %s\n", negative[i], desc)
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
