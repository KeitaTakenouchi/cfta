package automaton

import (
	"testing"
)

func TestNewCFTA(t *testing.T) {
	cfta := NewCFTA()

	// add transitions
	and := NewAlphabet("AND", 2)
	cfta.AddTransition(and, []int{0, 0}, 1)
	cfta.AddTransition(and, []int{0, 1}, 0)
	cfta.AddTransition(and, []int{1, 0}, 0)
	cfta.AddTransition(and, []int{1, 1}, 1)

	one := NewAlphabet("1", 0)
	cfta.AddTransition(one, []int{}, 1)

	zero := NewAlphabet("0", 0)
	cfta.AddTransition(zero, []int{}, 0)

	not := NewAlphabet("NOT", 1)
	cfta.AddTransition(not, []int{0}, 1)
	cfta.AddTransition(not, []int{1}, 0)

	// final states
	cfta.AddFinalState(1)

	cfta.dump()
}
