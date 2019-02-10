package palindrome

import "sort"

type word string

func (w word) Len() int           { return len(w) }
func (w word) Less(i, j int) bool { return w[i] < w[j] }
func (w word) Swap(i, j int)      {}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		eq := !s.Less(i, j) && !s.Less(j, i)
		if !eq {
			return false
		}
	}
	return true
}
