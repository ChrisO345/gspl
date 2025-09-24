package brancher

import "gonum.org/v1/gonum/mat"

// Tree represents the branch-and-bound tree structure used to solve integer linear programming problems.
type Tree struct {
	root       *Branch
	incumbentZ float64
	incumbentX *mat.VecDense

	vars []string // Variable names for the problem
}

// Branch represents a node in the branch-and-bound tree.
type Branch struct {
	left         *Branch
	right        *Branch
	value        float64
	branchStatus BranchStatus

	highestLower *float64
	lowestUpper  *float64 // Incumbent Solution of the Tree
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

// BranchSense represents the sense of branching in the branch-and-bound algorithm.
type BranchSense int

const (
	BranchSenseMinimize = BranchSense(-1)
	BranchSenseMaximize = BranchSense(1)
)
