package simplex

// ExitFlag represents the status code returned by the simplex solver.
type ExitFlag int

const (
	ExitUnbounded ExitFlag = iota - 1
	ExitOptimal
	ExitInfeasible
)
