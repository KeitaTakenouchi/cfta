package syntaxtree

import (
	"fmt"
	"testing"
)

func TestSyntaxNode_addSubNode(t *testing.T) {
	tests := []struct {
		name  string
		nodes []*SyntaxNode
		want  int
	}{
		{
			nodes: []*SyntaxNode{
				&SyntaxNode{
					Symbol: "sub_A",
				},
			},
			want: 1,
		},
		{
			nodes: []*SyntaxNode{
				&SyntaxNode{
					Symbol: "sub_A",
				},
				&SyntaxNode{
					Symbol: "sub_B",
				},
			},
			want: 2,
		},
		{
			nodes: []*SyntaxNode{
				&SyntaxNode{
					Symbol: "sub_A",
				},
				&SyntaxNode{
					Symbol: "sub_B",
				},
				&SyntaxNode{
					Symbol: "sub_C",
				},
			},
			want: 3,
		},
		{
			nodes: []*SyntaxNode{
				&SyntaxNode{
					Symbol: "sub_A",
					SubNodes: []*SyntaxNode{
						&SyntaxNode{
							Symbol: "sub_A_A",
						},
						&SyntaxNode{
							Symbol: "sub_A_B",
						},
					},
				},
				&SyntaxNode{
					Symbol: "sub_B",
				},
				&SyntaxNode{
					Symbol: "sub_C",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := NewSyntaxNode("test")
			sn.AddSubNode(tt.nodes...)
			if tt.want != len(sn.SubNodes) {
				t.Errorf("wrong size of children. expected=%d, but actual=%d", tt.want, len(sn.SubNodes))
			}
		})
	}
}
