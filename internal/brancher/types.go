package brancher

// FIXME: The following code is all placeholder, and is partially indicative of the future implementation

type IntegerProblem struct {
	NumVars        int
	NumConstraints int

	Objective     []float64
	Constraints   [][]float64
	RHS           []float64
	ConstraintDir []string
	Sense         string // "max" or "min"

	// Best known solution
	BestSolution []float64
	BestObj      float64

	// User-supplied strategy functions
	Branch    BranchFunc
	Heuristic HeuristicFunc
	Cut       CutFunc
}

// Branching: takes a node, returns two child nodes (or more if you generalize).
type BranchFunc func(node *Node) ([]*Node, error)

// Heuristic: try to find a feasible integer solution quickly.
type HeuristicFunc func(node *Node) ([]float64, float64, bool)

// Cutting planes: generate additional constraints for a node.
type CutFunc func(node *Node) [][]float64

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
