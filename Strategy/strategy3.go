package strategy

import "fmt"

type SolutionApproach interface {
	solve()
}

type Recursivesolver struct{}

type Iterativesolver struct{}

func (rs *Recursivesolver) solve() {
	fmt.Printf("\nrecursive solver strategy\n")
}

func (is *Iterativesolver) solve() {
	fmt.Printf("\niterative solver strategy\n")
}

type Wrapper struct {
	approaches SolutionApproach
}

func (w *Wrapper) SetStrategy(strategy SolutionApproach) {
	w.approaches = strategy
}

func (w *Wrapper) Solve() {
	if w.approaches == nil {
		fmt.Printf("\nInvalid approach")
	}
	w.approaches.solve()
}
