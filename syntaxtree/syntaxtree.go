package syntaxtree

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

func (sn *SyntaxNode) addSubNode(node ...*SyntaxNode) {
	sn.SubNodes = append(sn.SubNodes, node...)
}
