package common

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestDefaultSolverConfig(t *testing.T) {
	cfg := DefaultSolverConfig()
	assert.False(t, cfg.Logging)
	assert.Equal(t, cfg.Tolerance, 1e-6)
	assert.Equal(t, cfg.MaxIterations, 1000)
	assert.Equal(t, cfg.GapSensitivity, 0.05)
	assert.NotNil(t, cfg.Branch) // FIXME: should be nil??
	assert.NotNil(t, cfg.Heuristic)
	assert.NotNil(t, cfg.Cut)
	assert.Equal(t, cfg.Threads, 0)
}

func TestValidateSolverConfig(t *testing.T) {
	cfg := DefaultSolverConfig()
	err := ValidateSolverConfig(cfg)
	assert.Nil(t, err)
}
