package common

import "gonum.org/v1/gonum/mat"

type IntegerProgram struct {
	SCF *StandardComputationalForm

	// IP Specific fields
	ConstraintDir []string

	// Best known solution
	BestSolution *mat.VecDense // x*
	BestObj      float64

	// User-supplied strategy functions
	Branch    BranchFunc
	Heuristic HeuristicFunc
	Cut       CutFunc
}

// FIXME: This is just a placeholder struct for Node. This will change.
type Node struct {
	SCF *StandardComputationalForm

	ID       int
	ParentID int
	Depth    int

	Bounds [][2]float64

	RelaxedSol []float64
	RelaxedObj float64

	IsFeasible bool
	IsInteger  bool

	BranchVar   int
	BranchValue float64

	LowerBound float64
}
