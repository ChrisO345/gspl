package solver

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/internal/common"
)

func TestWithTolerance(t *testing.T) {
	cfg := NewSolverConfig(WithTolerance(1e-5))
	assert.Equal(t, cfg.Tolerance, 1e-5)
}

func TestWithMaxIterations(t *testing.T) {
	cfg := NewSolverConfig(WithMaxIterations(100))
	assert.Equal(t, cfg.MaxIterations, 100)
}

func TestWithGapSensitivity(t *testing.T) {
	cfg := NewSolverConfig(WithGapSensitivity(0.01))
	assert.Equal(t, cfg.GapSensitivity, 0.01)
}

func TestWithLogging(t *testing.T) {
	cfg := NewSolverConfig(WithLogging(true))
	assert.True(t, cfg.Logging)
}

func TestNewSolverConfig_Defaults(t *testing.T) {
	cfg := NewSolverConfig()
	defaults := common.DefaultSolverConfig()
	assert.Equal(t, cfg.Tolerance, defaults.Tolerance)
	assert.Equal(t, cfg.MaxIterations, defaults.MaxIterations)
	assert.Equal(t, cfg.GapSensitivity, defaults.GapSensitivity)
	assert.Equal(t, cfg.Logging, defaults.Logging)
}

func TestWithThreads_Panic(t *testing.T) {
	assert.Panic(t, func() {
		_ = WithThreads(4)
	})
}

func TestUnimplementedStrategyOptions_Panic(t *testing.T) {
	assert.Panic(t, func() { _ = WithBranch(nil) })
	assert.Panic(t, func() { _ = WithHeuristic(nil) })
	assert.Panic(t, func() { _ = WithCut(nil) })
}
