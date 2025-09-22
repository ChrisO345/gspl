package common_test

import (
	"testing"

	"github.com/chriso345/gore/assert"
	"github.com/chriso345/gspl/internal/common"
)

func TestDefaultSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	assert.Equal(t, cfg.Tolerance, 1e-6)
	assert.Equal(t, cfg.MaxIterations, 1000)
	assert.Equal(t, cfg.GapSensitivity, 0.05)
	assert.Equal(t, cfg.Threads, 0)
	assert.Equal(t, cfg.RandomSeed, 42)
	assert.Equal(t, cfg.SolverMethod, common.SimplexMethod)
}

func TestValidateSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	err := common.ValidateSolverConfig(cfg)
	assert.Nil(t, err)
}
