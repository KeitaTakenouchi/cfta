package automaton

import (
	"github.com/KeitaTakenouchi/cfta/tree"
	"testing"
)

func TestCFTA_Evaluate(t *testing.T) {
	cfta := createLogicCFTA()

	tests := []struct {
		name string
		tree tree.SyntaxTree
		want bool
	}{
		{
			name: "0",
			tree: *tree.NewSyntaxTree("0"),
			want: true,
		},
		{
			name: "1",
			tree: *tree.NewSyntaxTree("1"),
			want: false,
		},
		{
			name: "AND(0, 1)",
			tree: *tree.NewSyntaxTreeWithSubs("AND",
				tree.NewSyntaxTree("0"),
				tree.NewSyntaxTree("1"),
			),
			want: true,
		},
		{
			name: "AND(1, 1)",
			tree: *tree.NewSyntaxTreeWithSubs("AND",
				tree.NewSyntaxTree("1"),
				tree.NewSyntaxTree("1"),
			),
			want: false,
		},
		{
			name: "AND(0, 0)",
			tree: *tree.NewSyntaxTreeWithSubs("AND",
				tree.NewSyntaxTree("0"),
				tree.NewSyntaxTree("0"),
			),
			want: true,
		},
		{
			name: "AND(0, NOT(1))",
			tree: *tree.NewSyntaxTreeWithSubs("AND",
				tree.NewSyntaxTree("0"),
				tree.NewSyntaxTreeWithSubs("NOT",
					tree.NewSyntaxTree("1"),
				),
			),
			want: true,
		},
		{
			name: "NOT(AND(0, NOT(1)))",
			tree: *tree.NewSyntaxTreeWithSubs("NOT",
				tree.NewSyntaxTreeWithSubs("AND",
					tree.NewSyntaxTree("0"),
					tree.NewSyntaxTreeWithSubs("NOT",
						tree.NewSyntaxTree("1"),
					),
				),
			),
			want: false,
		},
		{
			name: "AND(AND(1, 0), AND(1, 1))",
			tree: *tree.NewSyntaxTreeWithSubs("AND",
				tree.NewSyntaxTreeWithSubs("AND",
					tree.NewSyntaxTree("1"),
					tree.NewSyntaxTree("0"),
				),
				tree.NewSyntaxTreeWithSubs("AND",
					tree.NewSyntaxTree("1"),
					tree.NewSyntaxTree("1"),
				),
			),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cfta.Evaluate(tt.tree); got != tt.want {
				t.Errorf("CFTA.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createLogicCFTA() *CFTA {
	cfta := NewCFTA()

	// add transitions
	and := NewAlphabet("AND", 2)
	cfta.AddTransition(and, []int{0, 0}, 0)
	cfta.AddTransition(and, []int{0, 1}, 0)
	cfta.AddTransition(and, []int{1, 0}, 0)
	cfta.AddTransition(and, []int{1, 1}, 1)

	one := NewAlphabet("1", 0)
	cfta.AddTransition(one, []int{}, 1)

	zero := NewAlphabet("0", 0)
	cfta.AddTransition(zero, []int{}, 0)

	not := NewAlphabet("NOT", 1)
	cfta.AddTransition(not, []int{0}, 1)
	cfta.AddTransition(not, []int{1}, 0)

	// final states
	cfta.AddFinalState(0)

	return cfta
}
