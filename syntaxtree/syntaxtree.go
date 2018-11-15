package syntaxtree

import (
	"bytes"
)

// SyntaxTree is a tree of syntax tree.
type SyntaxTree struct {
	Symbol   string
	SubTrees []*SyntaxTree
}

// NewSyntaxTree is a constructor for SyntaxTree.
func NewSyntaxTree(symbol string) *SyntaxTree {
	return &SyntaxTree{
		Symbol:   symbol,
		SubTrees: make([]*SyntaxTree, 0),
	}
}

// NewSyntaxTreeWithSubs is a constructor for SyntaxTree with defalult sub trees.
func NewSyntaxTreeWithSubs(symbol string, trees ...*SyntaxTree) *SyntaxTree {
	tree := &SyntaxTree{
		Symbol:   symbol,
		SubTrees: make([]*SyntaxTree, 0),
	}
	tree.AddSubTree(trees...)
	return tree
}

// AddSubTree adds sub trees to syntax tree.
func (sn *SyntaxTree) AddSubTree(tree ...*SyntaxTree) {
	sn.SubTrees = append(sn.SubTrees, tree...)
}

func (sn *SyntaxTree) String() string {
	var buf bytes.Buffer

	buf.WriteString(sn.Symbol)

	if len(sn.SubTrees) != 0 {
		buf.WriteString("(")
	}

	for i, sub := range sn.SubTrees {
		if i == len(sn.SubTrees)-1 {
			buf.WriteString(sub.String())
		} else {
			buf.WriteString(sub.String() + ", ")
		}
	}

	if len(sn.SubTrees) != 0 {
		buf.WriteString(")")
	}

	return buf.String()
}
