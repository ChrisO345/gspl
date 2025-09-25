package common

// SolverStatus represents the status of the solver
type SolverStatus int

const (
	SolverStatusNotSolved SolverStatus = iota
	SolverStatusOptimal
	SolverStatusInfeasible
	SolverStatusUnbounded
)
