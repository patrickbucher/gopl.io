package ex09

import (
	"testing"
)

var testInput = `
Yesterday I was eating some $food
I like $food -- especially with some $drink
However, if I eat too much $food and drink too much $drink
I'm getting a bit $eatenTooMuch and $drankTooMuch
But as long as I don't overdue it, some $food and some $drink
neither make me $eatenTooMuch nor $drankTooMuch
`

var testExpected = `
Yesterday I was eating some fried chicken
I like fried chicken -- especially with some lager beer
However, if I eat too much fried chicken and drink too much lager beer
I'm getting a bit bloated and shitfaced
But as long as I don't overdue it, some fried chicken and some lager beer
neither make me bloated nor shitfaced
`

func testFunc(str string) string {
	switch str {
	case "food":
		return "fried chicken"
	case "drink":
		return "lager beer"
	case "eatenTooMuch":
		return "bloated"
	case "drankTooMuch":
		return "shitfaced"
	default:
		return ""
	}
}

func TestExpand(t *testing.T) {
	got := Expand(testInput, testFunc)
	if got != testExpected {
		t.Errorf(`expected "%s" but was "%s"\n`, testExpected, got)
	}
}
