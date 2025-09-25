package brancher

import (
	"errors"

	"github.com/chriso345/gspl/internal/common"
)

func BranchAndBound(ip *common.IntegerProgram) error {
	// Define the strategies to be used in tree traversal
	defineStrategies(ip)

	return errors.New("Branch and Bound algorithm not yet implemented")
}

// defineStrategies sets the strategies to be used in the Branch and Bound algorithm
func defineStrategies(ip *common.IntegerProgram) {
	if ip.Branch == nil {
		branchFunc = DefaultBranch
	} else {
		branchFunc = ip.Branch
	}

	if ip.Heuristic == nil {
		heuristicFunc = DefaultHeuristic
	} else {
		heuristicFunc = ip.Heuristic
	}

	if ip.Cut == nil {
		cutFunc = DefaultCut
	} else {
		cutFunc = ip.Cut
	}
}

// Strategy function variables
var branchFunc common.BranchFunc
var heuristicFunc common.HeuristicFunc
var cutFunc common.CutFunc
