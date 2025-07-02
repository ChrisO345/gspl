package common

// SolverConfig holds the actual configuration with no pointers.
type SolverConfig struct {
	// TimeLimit        // TODO: Add time.Duration field later
	Tolerance     float64
	MaxIterations int

	GapSensitivity       float64
	UseCuttingPlanes     bool
	BranchingStrategy    BranchingStrategy
	HeuristicStrategy    HeuristicStrategy
	StrongBranchingDepth int
	Threads              int

	RandomSeed   int
	SolverMethod SolverMethod
	Logging      bool
}

// DefaultSolverConfig returns the default solver configuration.
func DefaultSolverConfig() *SolverConfig {
	return &SolverConfig{
		Tolerance:            1e-6,
		MaxIterations:        1000,
		GapSensitivity:       0.05,
		UseCuttingPlanes:     false,
		BranchingStrategy:    FirstFractional,
		HeuristicStrategy:    RandomHeuristic,
		StrongBranchingDepth: 0,
		Threads:              0, // 0 means use all available cores
		RandomSeed:           42,
		SolverMethod:         SimplexMethod,
		Logging:              false, // Default logging is off
	}
}

// validateSolverConfig checks if the SolverConfig is valid.
func ValidateSolverConfig(cfg *SolverConfig) error {
	// TODO: Impement validation logic
	_ = cfg
	return nil
}

// BranchingStrategy represents the strategy used for branching in integer programming.
type BranchingStrategy string

const (
	FirstFractional BranchingStrategy = "first-fractional"
	MostFractional  BranchingStrategy = "most-fractional"
	LeastFractional BranchingStrategy = "least-fractional"
	RandomBranching BranchingStrategy = "random-branching"
)

// HeuristicStrategy represents the strategy used for heuristics in integer programming.
type HeuristicStrategy string

const (
	RandomHeuristic      HeuristicStrategy = "random"
	LargestInfeasibility HeuristicStrategy = "largest-infeasibility"
)

// SolverMethod represents the method used for solving linear programming problems or the
// LP relaxation of integer programming problems.
type SolverMethod string

const (
	SimplexMethod SolverMethod = "simplex" // Currently the only option is the simplex method
)
