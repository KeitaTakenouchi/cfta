package automaton

import (
	"fmt"
	"github.com/KeitaTakenouchi/cfta/tree"
	"strconv"
	"strings"
)

// CFTA represents Concrete Finite Tree Automata.
type CFTA struct {
	States      []State
	Alphabets   []Alphabet
	FinalStates []State
	Transitions map[TransitionInput]State
}

// NewCFTA is a constructor of CFTA.
func NewCFTA() *CFTA {
	return &CFTA{
		States:      make([]State, 0),
		Alphabets:   make([]Alphabet, 0),
		FinalStates: make([]State, 0),
		Transitions: make(map[TransitionInput]State),
	}
}

// AddTransition adds transition.
func (cfta *CFTA) AddTransition(f Alphabet, fromStateIds []int, toStateId int) {
	var parameterStates []State

	for _, id := range fromStateIds {
		parameterStates = append(parameterStates, cfta.getState(id))
	}

	// add states if not exist.
	for _, paramState := range parameterStates {
		exists := false
		for _, existingState := range cfta.States {
			if paramState == existingState {
				exists = true
				break
			}
		}
		if !exists {
			cfta.States = append(cfta.States, paramState)
		}
	}

	// add alphabets if not exist.
	exists := false
	for _, existingAlphabet := range cfta.Alphabets {
		if f == existingAlphabet {
			exists = true
		}
	}
	if !exists {
		cfta.Alphabets = append(cfta.Alphabets, f)
	}

	// add a transition
	input := newTransitionInput(f, parameterStates)
	_, ok := cfta.Transitions[input]
	if !ok {
		cfta.Transitions[input] = cfta.getState(toStateId)
	} else {
		fmt.Printf("Key \"%s\" already exists.\n", input.String())
	}
}

//AddFinalState adds a final state
func (cfta *CFTA) AddFinalState(finalStateId int) {
	state := cfta.getState(finalStateId)
	cfta.FinalStates = append(cfta.FinalStates, state)
}

// Evaluate returns true if the given tree is accepted or false otherwise.
func (cfta *CFTA) Evaluate(tree tree.SyntaxTree) bool {
	state := cfta.eval(tree)
	for _, fstate := range cfta.FinalStates {
		if fstate == state {
			return true
		}
	}
	return false
}

func (cfta *CFTA) eval(tree tree.SyntaxTree) State {
	alphabet := cfta.searchAlphabetBySymbol(tree.Symbol)
	if alphabet == nil {
		msg := fmt.Sprintf("Invalid alphabet: %s.", tree.Symbol)
		panic(msg)
	}

	subStates := make([]State, 0)
	for _, subTree := range tree.SubTrees {
		subStates = append(subStates, cfta.eval(*subTree))
	}
	if len(subStates) != alphabet.arity {
		msg := fmt.Sprintf("Invalid tree. arity of %s is not %d.", alphabet.symbol, len(subStates))
		panic(msg)
	}

	transInput := newTransitionInput(*alphabet, subStates)
	nextState, ok := cfta.Transitions[transInput]
	if !ok {
		msg := fmt.Sprintf("Invalid transition %s.", transInput.String())
		panic(msg)
	}
	return nextState
}

func (cfta *CFTA) searchAlphabetBySymbol(symbol string) *Alphabet {
	for _, alphabet := range cfta.Alphabets {
		if symbol == alphabet.symbol {
			return &alphabet
		}
	}
	return nil
}

func (cfta *CFTA) getState(id int) State {
	for _, state := range cfta.States {
		if id == state.id {
			return state
		}
	}
	s := NewState(id)
	cfta.States = append(cfta.States, s)
	return s
}

func (cfta *CFTA) dump() {
	fmt.Println("------ states ------")
	for _, state := range cfta.States {
		fmt.Print(state.String())
		for _, final := range cfta.FinalStates {
			if final == state {
				fmt.Print(" (final)")
			}
		}
		fmt.Println()
	}
	fmt.Println("------ alphabets ------")
	for _, alpha := range cfta.Alphabets {
		fmt.Println(alpha.String())
	}

	fmt.Println("------ transactions ------")
	for key, value := range cfta.Transitions {
		fmt.Println(key.String(), " -> ", value.String())
	}
}

// State is a state of automaton.
type State struct {
	id int
}

// NewState returns a new state.
func NewState(id int) State {
	return State{
		id: id,
	}
}

func (s *State) String() string {
	return "q_" + strconv.Itoa(s.id)
}

// Alphabet is a symbol with arity.
type Alphabet struct {
	symbol string
	arity  int
}

// NewAlphabet is a constructor of Alphabet.
func NewAlphabet(symbol string, arity int) Alphabet {
	return Alphabet{
		symbol: symbol,
		arity:  arity,
	}
}

func (al *Alphabet) String() string {
	return al.symbol + "[" + strconv.Itoa(al.arity) + "]"
}

// TransitionInput is a pair of alphabet and parameters.
type TransitionInput struct {
	f      Alphabet
	params parameters
}

func newTransitionInput(f Alphabet, states []State) TransitionInput {
	if f.arity != len(states) {
		panic("illegal arguments.")
	}

	var params parameters
	switch {
	case len(states) == 0:
		params = param0{}
	case len(states) == 1:
		params = param1{
			prm: states[0],
		}
	case len(states) == 2:
		params = param2{
			fst: states[0],
			snd: states[1],
		}
	}
	return TransitionInput{
		f:      f,
		params: params,
	}
}

func (ti *TransitionInput) String() string {
	var buf strings.Builder
	buf.WriteString(ti.f.symbol)

	buf.WriteString("(")
	states := ti.params.getParams()
	for i, state := range states {
		if i == len(states)-1 {
			buf.WriteString(state.String())
		} else {
			buf.WriteString(state.String())
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")

	return buf.String()
}

type parameters interface {
	getParams() []State
}

type param0 struct{}

func (p param0) getParams() []State {
	return make([]State, 0)
}

type param1 struct {
	prm State
}

func (p param1) getParams() []State {
	ret := make([]State, 0)
	ret = append(ret, p.prm)
	return ret
}

type param2 struct {
	fst State
	snd State
}

func (p param2) getParams() []State {
	ret := make([]State, 0)
	ret = append(ret, p.fst)
	ret = append(ret, p.snd)
	return ret
}
