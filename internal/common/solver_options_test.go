package common_test

import (
	"testing"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/testutils/assert"
)

func TestDefaultSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	assert.AssertEqual(t, cfg.Tolerance, 1e-6)
	assert.AssertEqual(t, cfg.MaxIterations, 1000)
	assert.AssertEqual(t, cfg.GapSensitivity, 0.05)
	assert.AssertEqual(t, cfg.UseCuttingPlanes, false)
	assert.AssertEqual(t, cfg.BranchingStrategy, common.FirstFractional)
	assert.AssertEqual(t, cfg.HeuristicStrategy, common.RandomHeuristic)
	assert.AssertEqual(t, cfg.StrongBranchingDepth, 0)
	assert.AssertEqual(t, cfg.Threads, 0)
	assert.AssertEqual(t, cfg.RandomSeed, 42)
	assert.AssertEqual(t, cfg.SolverMethod, common.SimplexMethod)
}

func TestValidateSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	err := common.ValidateSolverConfig(cfg)
	assert.AssertNil(t, err)
}
