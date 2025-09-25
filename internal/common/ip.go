package common

type IntegerProgram struct {
	SCF *StandardComputationalForm

	// IP Specific fields
	ConstraintDir []string

	// Best known solution
	BestSolution []float64
	BestObj      float64

	// User-supplied strategy functions
	Branch    BranchFunc
	Heuristic HeuristicFunc
	Cut       CutFunc
}

// FIXME: This is just a placeholder struct for Node. This will change.
type Node struct {
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
