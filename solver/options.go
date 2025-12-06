package solver

import (
	"context"

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

// WithContext sets a context for cancellation of long-running solves.
func WithContext(ctx context.Context) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.Ctx = ctx
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
	panic("multi-threading not yet implemented")
}

// WithLogging enables or disables logging.
func WithLogging(enabled bool) SolverOption {
	return func(cfg *common.SolverConfig) {
		cfg.Logging = enabled
	}
}

/// Strategy Functions Options

// WithBranch sets the branching strategy function.
func WithBranch(common.BranchFunc) SolverOption {
	panic("branching not yet implemented")
}

// WithHeuristic sets the heuristic strategy function.
func WithHeuristic(common.HeuristicFunc) SolverOption {
	panic("heuristics not yet implemented")
}

// WithCut sets the cut generation function.
func WithCut(common.CutFunc) SolverOption {
	panic("cutting planes not yet implemented")
}

/// Helpers

// NewSolverConfig builds a SolverConfig applying all options on defaults.
func NewSolverConfig(opts ...SolverOption) *common.SolverConfig {
	cfg := common.DefaultSolverConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
