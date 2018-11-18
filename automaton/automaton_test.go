package automaton

import (
	"github.com/KeitaTakenouchi/cfta/grammar"
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

func TestCFTA_Evaluate2(t *testing.T) {
	cfta := createCalcCFTA()

	tests := []struct {
		name string
		tree tree.SyntaxTree
		want bool
	}{
		{
			name: "(id(1) + 2) * 3",
			tree: *tree.NewSyntaxTreeWithSubs("*",
				tree.NewSyntaxTreeWithSubs("+",
					tree.NewSyntaxTreeWithSubs("id",
						tree.NewSyntaxTree("1"),
					),
					tree.NewSyntaxTree("2"),
				),
				tree.NewSyntaxTree("3"),
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
	and := grammar.NewToken("AND", 2)
	cfta.AddTransition(and, []int{0, 0}, 0)
	cfta.AddTransition(and, []int{0, 1}, 0)
	cfta.AddTransition(and, []int{1, 0}, 0)
	cfta.AddTransition(and, []int{1, 1}, 1)

	one := grammar.NewToken("1", 0)
	cfta.AddTransition(one, []int{}, 1)

	zero := grammar.NewToken("0", 0)
	cfta.AddTransition(zero, []int{}, 0)

	not := grammar.NewToken("NOT", 1)
	cfta.AddTransition(not, []int{0}, 1)
	cfta.AddTransition(not, []int{1}, 0)

	// final states
	cfta.AddFinalState(0)

	return cfta
}

func createCalcCFTA() *CFTA {
	cfta := NewCFTA()

	// states
	const (
		x1 = iota
		t2
		t3
		n1
		n3
		n9
	)

	add := grammar.NewToken("+", 2)
	cfta.AddTransition(add, []int{n1, t2}, n3)

	mult := grammar.NewToken("*", 2)
	cfta.AddTransition(mult, []int{n3, t3}, n9)

	id := grammar.NewToken("id", 1)
	cfta.AddTransition(id, []int{x1}, n1)

	one := grammar.NewToken("1", 0)
	cfta.AddTransition(one, []int{}, x1)

	two := grammar.NewToken("2", 0)
	cfta.AddTransition(two, []int{}, t2)

	three := grammar.NewToken("3", 0)
	cfta.AddTransition(three, []int{}, t3)

	cfta.AddFinalState(n9)

	return cfta
}
