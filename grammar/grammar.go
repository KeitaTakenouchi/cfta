package grammar

import (
	"fmt"
	"strconv"
)

// Grammar represents a grammar of DSL.
type Grammar struct {
	terminals       []Token
	nonTerminals    []NonTerminalSymbol
	productionRules map[NonTerminalSymbol][]rightHandSide
	startSymbol     []NonTerminalSymbol
}

// NewGrammar is a  constructor of Grammar.
func NewGrammar() *Grammar {
	g := &Grammar{
		terminals:       make([]Token, 0),
		nonTerminals:    make([]NonTerminalSymbol, 0),
		productionRules: make(map[NonTerminalSymbol][]rightHandSide),
		startSymbol:     make([]NonTerminalSymbol, 0),
	}
	return g
}

// AddProdoctionRule adds a prodoction rule to the grammar.
func (g *Grammar) AddProdoctionRule(left NonTerminalSymbol, right Token, args ...Symbol) {
	if right.Arity != len(args) {
		msg := fmt.Sprintf("arity is %d but length of args is %d", right.Arity, len(args))
		panic(msg)
	}

	rules, ok := g.productionRules[left]
	if !ok {
		rules = make([]rightHandSide, 0)
	}
	rhs := newRightHandSide(right, args...)
	g.productionRules[left] = append(rules, rhs)
}

// SetStartSymbol sets a symbol as a topmost one of the grammar.
func (g *Grammar) SetStartSymbol(symbol NonTerminalSymbol) {
	g.startSymbol = append(g.startSymbol, symbol)
}

// CreateToken is a constructor of Symbol.
func (g *Grammar) CreateToken(symbol string, arity int) Token {
	tk := Token{
		Symbol{symbol},
		arity,
	}
	g.terminals = append(g.terminals, tk)
	return tk
}

// CreateNonTerminalSymbol is a constructor of NonTerminalSymbol.
func (g *Grammar) CreateNonTerminalSymbol(symbol string) NonTerminalSymbol {
	s := NonTerminalSymbol{
		Symbol{symbol},
	}
	g.nonTerminals = append(g.nonTerminals, s)
	return s
}

// Symbol is a symbol in DSL.
type Symbol struct {
	Text string
}

// NonTerminalSymbol is a non terminal symbol of DSL
type NonTerminalSymbol struct {
	Symbol
}

// Token is a symbol with arity.
type Token struct {
	Symbol
	Arity int
}

func (al *Token) String() string {
	return al.Text + "[" + strconv.Itoa(al.Arity) + "]"
}

type rightHandSide struct {
	symbol Token
	args   []Symbol
}

func newRightHandSide(symbol Token, args ...Symbol) rightHandSide {
	return rightHandSide{
		symbol: symbol,
		args:   args,
	}
}
