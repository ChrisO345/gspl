package common_test

import (
	"testing"

	"github.com/chriso345/gspl/internal/common"
)

func TestDefaultSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	if cfg.Tolerance != 1e-6 {
		t.Errorf("Expected Tolerance to be 1e-6, got %v", cfg.Tolerance)
	}
	if cfg.MaxIterations != 1000 {
		t.Errorf("Expected MaxIterations to be 1000, got %d", cfg.MaxIterations)
	}
	if cfg.GapSensitivity != 0.05 {
		t.Errorf("Expected GapSensitivity to be 0.05, got %v", cfg.GapSensitivity)
	}
	if cfg.UseCuttingPlanes != false {
		t.Errorf("Expected UseCuttingPlanes to be false, got %v", cfg.UseCuttingPlanes)
	}
	if cfg.BranchingStrategy != common.FirstFractional {
		t.Errorf("Expected BranchingStrategy to be %v, got %v", common.FirstFractional, cfg.BranchingStrategy)
	}
	if cfg.HeuristicStrategy != common.RandomHeuristic {
		t.Errorf("Expected HeuristicStrategy to be %v, got %v", common.RandomHeuristic, cfg.HeuristicStrategy)
	}
	if cfg.StrongBranchingDepth != 0 {
		t.Errorf("Expected StrongBranchingDepth to be 0, got %d", cfg.StrongBranchingDepth)
	}
	if cfg.Threads != 0 {
		t.Errorf("Expected Threads to be 0, got %d", cfg.Threads)
	}
	if cfg.RandomSeed != 42 {
		t.Errorf("Expected RandomSeed to be 42, got %d", cfg.RandomSeed)
	}
	if cfg.SolverMethod != common.SimplexMethod {
		t.Errorf("Expected SolverMethod to be %v, got %v", common.SimplexMethod, cfg.SolverMethod)
	}
}

func TestValidateSolverConfig(t *testing.T) {
	cfg := common.DefaultSolverConfig()

	err := common.ValidateSolverConfig(cfg)
	if err != nil {
		t.Errorf("Expected ValidateSolverConfig to return nil, got %v", err)
	}
}
