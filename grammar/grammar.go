package grammar

import (
	"strconv"
)

// Grammar represents a grammar of DSL.
type Grammar struct {
	terminals       []Token
	nonTerminals    []NonTerminalSymbol
	productionRules map[NonTerminalSymbol]rhs
}

// Symbol is a symbol in DSL.
type Symbol struct {
	Text string
}

// NewSymbol is a constructor of Symbol.
func NewSymbol(symbol string) Symbol {
	return Symbol{
		Text: symbol,
	}
}

// NonTerminalSymbol is a non terminal symbol of DSL
type NonTerminalSymbol struct {
	Symbol
}

// NewNonTerminalSymbol is a constructor of NonTerminalSymbol.
func NewNonTerminalSymbol(symbol string) NonTerminalSymbol {
	return NonTerminalSymbol{
		Symbol: Symbol{
			Text: symbol,
		},
	}
}

// Token is a symbol with arity.
type Token struct {
	Symbol
	Arity int
}

// NewToken is a constructor of Alphabet.
func NewToken(symbol string, arity int) Token {
	return Token{
		Symbol: Symbol{
			Text: symbol,
		},
		Arity: arity,
	}
}

func (al *Token) String() string {
	return al.Text + "[" + strconv.Itoa(al.Arity) + "]"
}

type rhs struct {
	Symbol Symbol
	args   []Symbol
}

func newRhs(symbol Symbol, args ...Symbol) rhs {
	return rhs{
		Symbol: symbol,
		args:   args,
	}
}
