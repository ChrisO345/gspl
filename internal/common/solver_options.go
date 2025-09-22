package common

// SolverConfig holds the actual configuration with no pointers.
type SolverConfig struct {
	// TimeLimit        // TODO: Add time.Duration field later
	Tolerance     float64
	MaxIterations int

	GapSensitivity float64
	Threads        int

	RandomSeed   int
	SolverMethod SolverMethod
	Logging      bool
}

// DefaultSolverConfig returns the default solver configuration.
func DefaultSolverConfig() *SolverConfig {
	return &SolverConfig{
		Tolerance:      1e-6,
		MaxIterations:  1000,
		GapSensitivity: 0.05,
		Threads:        0, // 0 means use all available cores
		RandomSeed:     42,
		SolverMethod:   SimplexMethod,
		Logging:        false, // Default logging is off
	}
}

// validateSolverConfig checks if the SolverConfig is valid.
func ValidateSolverConfig(cfg *SolverConfig) error {
	// TODO: Impement validation logic
	_ = cfg
	return nil
}

// SolverMethod represents the method used for solving linear programming problems or the
// LP relaxation of integer programming problems.
type SolverMethod string

const (
	SimplexMethod SolverMethod = "simplex" // Currently the only option is the simplex method
)
