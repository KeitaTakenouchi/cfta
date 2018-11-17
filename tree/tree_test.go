package tree

import (
	"testing"
)

func TestSyntaxTree_addSubTree(t *testing.T) {
	tests := []struct {
		name  string
		trees []*SyntaxTree
		want  int
	}{
		{
			trees: []*SyntaxTree{
				&SyntaxTree{
					Symbol: "sub_A",
				},
			},
			want: 1,
		},
		{
			trees: []*SyntaxTree{
				&SyntaxTree{
					Symbol: "sub_A",
				},
				&SyntaxTree{
					Symbol: "sub_B",
				},
			},
			want: 2,
		},
		{
			trees: []*SyntaxTree{
				&SyntaxTree{
					Symbol: "sub_A",
				},
				&SyntaxTree{
					Symbol: "sub_B",
				},
				&SyntaxTree{
					Symbol: "sub_C",
				},
			},
			want: 3,
		},
		{
			trees: []*SyntaxTree{
				&SyntaxTree{
					Symbol: "sub_A",
					SubTrees: []*SyntaxTree{
						&SyntaxTree{
							Symbol: "sub_A_A",
						},
						&SyntaxTree{
							Symbol: "sub_A_B",
						},
					},
				},
				&SyntaxTree{
					Symbol: "sub_B",
				},
				&SyntaxTree{
					Symbol: "sub_C",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := NewSyntaxTree("test")
			sn.AddSubTree(tt.trees...)
			if tt.want != len(sn.SubTrees) {
				t.Errorf("wrong size of children. expected=%d, but actual=%d", tt.want, len(sn.SubTrees))
			}
		})
	}
}
