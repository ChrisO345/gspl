package common

// SolverConfig holds the actual configuration with no pointers.
type SolverConfig struct {
	Logging       bool
	Tolerance     float64
	MaxIterations int

	// IP Specific Options
	GapSensitivity float64
	Branch         BranchFunc
	Heuristic      HeuristicFunc
	Cut            CutFunc

	// Not Yet Implemented
	Threads int
}

// DefaultSolverConfig returns the default solver configuration.
func DefaultSolverConfig() *SolverConfig {
	return &SolverConfig{
		Logging:       false, // Default logging is off
		Tolerance:     1e-6,
		MaxIterations: 1000,

		GapSensitivity: 0.05,
		Branch:         nil, // Default branching strategy defined in `brancher`
		Heuristic:      nil, // Default heuristic defined in `brancher`
		Cut:            nil, // Default cutting planes defined in `brancher`

		Threads: 0, // 0 means use all available cores
	}
}

// validateSolverConfig checks if the SolverConfig is valid.
func ValidateSolverConfig(cfg *SolverConfig) error {
	// TODO: Impement validation logic
	_ = cfg
	return nil
}
