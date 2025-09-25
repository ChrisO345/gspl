package brancher

import "github.com/chriso345/gspl/internal/common"

// DefaultBranch represents the default branching strategy.
//
// This branches on the first variable found that is not integer in the current node
func DefaultBranch(node *common.Node) ([]*common.Node, error) {
	return nil, nil
}

// DefaultHeuristic represents the default heuristic strategy.
//
// This does not implement any heuristic and simply returns nil
func DefaultHeuristic(node *common.Node) ([]float64, float64, bool) {
	return nil, 0, false
}

// DefaultCut represents the default cutting planes strategy.
//
// This does not implement any cutting planes and simply returns nil
func DefaultCut(node *common.Node) [][]float64 {
	return nil
}
