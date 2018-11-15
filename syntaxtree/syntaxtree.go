package syntaxtree

import (
	"bytes"
)

// SyntaxNode is a node of syntax tree.
type SyntaxNode struct {
	Symbol   string
	SubNodes []*SyntaxNode
}

// NewSyntaxNode is a constructor for SyntaxNode.
func NewSyntaxNode(symbol string) *SyntaxNode {
	return &SyntaxNode{
		Symbol:   symbol,
		SubNodes: make([]*SyntaxNode, 0),
	}
}

// NewSyntaxNodeWithSubNodes is a constructor for SyntaxNode with defalult sub nodes.
func NewSyntaxNodeWithSubNodes(symbol string, node ...*SyntaxNode) *SyntaxNode {
	return &SyntaxNode{
		Symbol:   symbol,
		SubNodes: make([]*SyntaxNode, 0),
	}
}

// AddSubNode adds sub nodes to syntax node.
func (sn *SyntaxNode) AddSubNode(node ...*SyntaxNode) {
	sn.SubNodes = append(sn.SubNodes, node...)
}

func (sn *SyntaxNode) String() string {
	var buf bytes.Buffer

	buf.WriteString(sn.Symbol)

	if len(sn.SubNodes) != 0 {
		buf.WriteString("(")
	}

	for i, sub := range sn.SubNodes {
		if i == len(sn.SubNodes)-1 {
			buf.WriteString(sub.String())
		} else {
			buf.WriteString(sub.String() + ", ")
		}
	}

	if len(sn.SubNodes) != 0 {
		buf.WriteString(")")
	}

	return buf.String()
}
