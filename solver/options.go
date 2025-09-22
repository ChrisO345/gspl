package solver

import (
	"github.com/chriso345/gspl/internal/common"
)

// SolverOption defines a function that modifies SolverConfig.
type SolverOption func(*common.SolverConfig)

// WithTolerance sets the tolerance.
func WithTolerance(t float64) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.Tolerance = t
	}
}

// WithMaxIterations sets the maximum number of iterations.
func WithMaxIterations(max int) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.MaxIterations = max
	}
}

// WithGapSensitivity sets the gap sensitivity.
func WithGapSensitivity(gap float64) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.GapSensitivity = gap
	}
}

// WithThreads sets the number of threads to use.
func WithThreads(n int) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.Threads = n
	}
}

// WithRandomSeed sets the random seed.
func WithRandomSeed(seed int) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.RandomSeed = seed
	}
}

// WithSolverMethod sets the solver method.
func WithSolverMethod(m common.SolverMethod) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.SolverMethod = m
	}
}

// WithLogging enables or disables logging.
func WithLogging(enabled bool) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.Logging = enabled
	}
}

// NewSolverConfig builds a SolverConfig applying all options on defaults.
func NewSolverConfig(opts ...SolverOption) *common.SolverConfig {
	cfg := common.DefaultSolverConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
