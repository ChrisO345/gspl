package brancher

import "gonum.org/v1/gonum/mat"

// Branch represents a node in the branch-and-bound tree.
type Branch struct {
	left         *Branch
	right        *Branch
	node         float64
	branchStatus BranchStatus
	highestLower *float64
	lowestUpper  *float64 // Incumbent Solution of the Tree
	vars         []string

	// Global tracking (in root node only)
	incumbentZ float64
	incumbentX *mat.VecDense
}

// Constraint represents an additional integer constraint in the branch-and-bound algorithm.
type Constraint struct {
	varIndex  int
	value     int
	direction ConstraintDirection
}

// BranchStatus represents the status of a branch in the branch-and-bound algorithm.
type BranchStatus int

const (
	BranchStatusUnexplored BranchStatus = iota
	BranchStatusInfeasible
	BranchStatusFeasible
	BranchStatusIncumbent
)

// ConstraintDirection represents the direction of the constraint in the branch-and-bound algorithm.
type ConstraintDirection int

const (
	ConstraintDirectionLE ConstraintDirection = iota
	ConstraintDirectionGE
)
