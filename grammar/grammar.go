package grammar

import (
	"strconv"
)

// Grammar represents a grammar of DSL.
type Grammar struct {
	terminals    []Token
	nonTerminals []NonTerminalSymbol
}

// NonTerminalSymbol is a non terminal symbol of DSL
type NonTerminalSymbol struct {
	Text string
}

// Token is a symbol with arity.
type Token struct {
	Symbol string
	Arity  int
}

// NewToken is a constructor of Alphabet.
func NewToken(symbol string, arity int) Token {
	return Token{
		Symbol: symbol,
		Arity:  arity,
	}
}

func (al *Token) String() string {
	return al.Symbol + "[" + strconv.Itoa(al.Arity) + "]"
}
