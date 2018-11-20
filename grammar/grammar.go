package grammar

import (
	"bytes"
	"fmt"
	"strconv"
)

// Grammar represents a grammar of DSL.
type Grammar struct {
	Terminals       []Token
	NonTerminals    []NonTerminalSymbol
	ProductionRules map[NonTerminalSymbol][]RightHandSide
	StartSymbol     []NonTerminalSymbol
}

// NewGrammar is a  constructor of Grammar.
func NewGrammar() *Grammar {
	g := &Grammar{
		Terminals:       make([]Token, 0),
		NonTerminals:    make([]NonTerminalSymbol, 0),
		ProductionRules: make(map[NonTerminalSymbol][]RightHandSide),
		StartSymbol:     make([]NonTerminalSymbol, 0),
	}
	return g
}

// AddProdoctionRule adds a prodoction rule to the grammar.
func (g *Grammar) AddProdoctionRule(left NonTerminalSymbol, right Token, args ...Symbol) {
	if right.Arity != len(args) {
		msg := fmt.Sprintf("arity is %d but length of args is %d", right.Arity, len(args))
		panic(msg)
	}

	rules, ok := g.ProductionRules[left]
	if !ok {
		rules = make([]RightHandSide, 0)
	}
	rhs := newRightHandSide(right, args...)
	g.ProductionRules[left] = append(rules, rhs)
}

// SetStartSymbol sets a symbol as a topmost one of the grammar.
func (g *Grammar) SetStartSymbol(symbol NonTerminalSymbol) {
	g.StartSymbol = append(g.StartSymbol, symbol)
}

// CreateToken is a constructor of Symbol.
func (g *Grammar) CreateToken(symbol string, arity int) Token {
	tk := NewToken(symbol, arity)
	g.Terminals = append(g.Terminals, tk)
	return tk
}

// CreateNonTerminalSymbol is a constructor of NonTerminalSymbol.
func (g *Grammar) CreateNonTerminalSymbol(symbol string) NonTerminalSymbol {
	s := NonTerminalSymbol{
		text: symbol,
	}
	g.NonTerminals = append(g.NonTerminals, s)
	return s
}

func (g *Grammar) dump() {
	fmt.Println("------- terminals -------")
	for _, t := range g.Terminals {
		fmt.Println(t.String())
	}

	fmt.Println("----- non terminals -----")
	for _, t := range g.NonTerminals {
		for _, start := range g.StartSymbol {
			if start.GetText() == t.GetText() {
				fmt.Println(t.GetText() + " (Start)")
			} else {
				fmt.Println(t.GetText())
			}
		}
	}

	fmt.Println("------ productions ------")
	for left, rhss := range g.ProductionRules {
		for _, rhs := range rhss {
			fmt.Println(left.GetText() + " -> " + rhs.String())
		}
	}
}

// Symbol is a symbol in DSL.
type Symbol interface {
	GetText() string
}

// NonTerminalSymbol is a non terminal symbol of DSL
type NonTerminalSymbol struct {
	text string
}

// GetText is for Symbol interface.
func (nt NonTerminalSymbol) GetText() string {
	return nt.text
}

// Token is a symbol with arity.
type Token struct {
	text  string
	Arity int
}

// NewToken is a constructor of Alphabet.
func NewToken(symbol string, arity int) Token {
	return Token{
		text:  symbol,
		Arity: arity,
	}
}

// GetText is for Symbol interface.
func (tk Token) GetText() string {
	return tk.text
}

func (tk Token) String() string {
	return tk.GetText() + "[" + strconv.Itoa(tk.Arity) + "]"
}

// RightHandSide is a right hand side of production rules.
type RightHandSide struct {
	Symbol Token
	Args   []Symbol
}

func newRightHandSide(symbol Token, args ...Symbol) RightHandSide {
	return RightHandSide{
		Symbol: symbol,
		Args:   args,
	}
}

func (rhs *RightHandSide) String() string {
	var buf bytes.Buffer
	buf.WriteString(rhs.Symbol.GetText())

	if len(rhs.Args) != 0 {
		buf.WriteString("(")
	}
	for i, arg := range rhs.Args {
		if i == len(rhs.Args)-1 {
			buf.WriteString(arg.GetText())
		} else {
			buf.WriteString(arg.GetText() + ", ")
		}
	}
	if len(rhs.Args) != 0 {
		buf.WriteString(")")
	}

	return buf.String()
}
