package solver

// solverConfig holds the actual configuration with no pointers.
type solverConfig struct {
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
}

// DefaultSolverConfig returns the default solver configuration.
func DefaultSolverConfig() *solverConfig {
	return &solverConfig{
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
	}
}

// validateSolverConfig checks if the solverConfig is valid.
func validateSolverConfig(cfg *solverConfig) error {
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

// SolverMethod represents the method used for solving linear programming problems.
type SolverMethod string

const (
	SimplexMethod SolverMethod = "simplex" // Currently the only option is the simplex method
)

// SolverOption defines a function that modifies solverConfig.
type SolverOption func(*solverConfig)

// WithTolerance sets the tolerance.
func WithTolerance(t float64) SolverOption {
	return func(cfg *solverConfig) {
		cfg.Tolerance = t
	}
}

// WithMaxIterations sets the maximum number of iterations.
func WithMaxIterations(max int) SolverOption {
	return func(cfg *solverConfig) {
		cfg.MaxIterations = max
	}
}

// WithGapSensitivity sets the gap sensitivity.
func WithGapSensitivity(gap float64) SolverOption {
	return func(cfg *solverConfig) {
		cfg.GapSensitivity = gap
	}
}

// WithUseCuttingPlanes enables or disables cutting planes.
func WithUseCuttingPlanes(enabled bool) SolverOption {
	return func(cfg *solverConfig) {
		cfg.UseCuttingPlanes = enabled
	}
}

// WithBranchingStrategy sets the branching strategy.
func WithBranchingStrategy(bs BranchingStrategy) SolverOption {
	return func(cfg *solverConfig) {
		cfg.BranchingStrategy = bs
	}
}

// WithHeuristicStrategy sets the heuristic strategy.
func WithHeuristicStrategy(hs HeuristicStrategy) SolverOption {
	return func(cfg *solverConfig) {
		cfg.HeuristicStrategy = hs
	}
}

// WithStrongBranchingDepth sets the strong branching depth.
func WithStrongBranchingDepth(depth int) SolverOption {
	return func(cfg *solverConfig) {
		cfg.StrongBranchingDepth = depth
	}
}

// WithThreads sets the number of threads to use.
func WithThreads(n int) SolverOption {
	return func(cfg *solverConfig) {
		cfg.Threads = n
	}
}

// WithRandomSeed sets the random seed.
func WithRandomSeed(seed int) SolverOption {
	return func(cfg *solverConfig) {
		cfg.RandomSeed = seed
	}
}

// WithSolverMethod sets the solver method.
func WithSolverMethod(m SolverMethod) SolverOption {
	return func(cfg *solverConfig) {
		cfg.SolverMethod = m
	}
}

// NewSolverConfig builds a solverConfig applying all options on defaults.
func NewSolverConfig(opts ...SolverOption) *solverConfig {
	cfg := DefaultSolverConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
