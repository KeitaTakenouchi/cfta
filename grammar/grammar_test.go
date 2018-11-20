package grammar

import (
	"testing"
)

func TestNewGrammar(t *testing.T) {
	grammar := createGrammar()
	grammar.dump()
}

func createGrammar() *Grammar {
	grammar := NewGrammar()

	nSymbol := grammar.CreateNonTerminalSymbol("N")
	tSymbol := grammar.CreateNonTerminalSymbol("T")

	idToken := grammar.CreateToken("id", 1)
	addToken := grammar.CreateToken("+", 2)
	multToken := grammar.CreateToken("*", 2)

	xToken := grammar.CreateToken("x", 0)
	twoToken := grammar.CreateToken("2", 0)
	threeToken := grammar.CreateToken("3", 0)

	grammar.AddProdoctionRule(nSymbol, idToken, xToken)
	grammar.AddProdoctionRule(nSymbol, addToken, nSymbol, tSymbol)
	grammar.AddProdoctionRule(nSymbol, multToken, nSymbol, tSymbol)

	grammar.AddProdoctionRule(tSymbol, twoToken)
	grammar.AddProdoctionRule(tSymbol, threeToken)

	grammar.SetStartSymbol(nSymbol)

	return grammar
}
