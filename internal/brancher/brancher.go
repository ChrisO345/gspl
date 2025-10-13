package brancher

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/simplex"
)

func BranchAndBound(ip *common.IntegerProgram) error {
	// Define the strategies to be used in tree traversal
	defineStrategies(ip)

	// Solve at the root
	rootNode := &common.Node{
		SCF: ip.SCF,
		// ID:       0,
		// ParentID: -1,
		// Depth:    0,
	}

	err := simplex.Simplex(rootNode.SCF)
	if err != nil {
		return fmt.Errorf("error solving root node: %v", err)
	}

	ip.BestObj = *rootNode.SCF.ObjectiveValue
	ip.BestSolution = rootNode.SCF.PrimalSolution
	rootNode.IsInteger = isIntegerFeasible(rootNode.SCF)

	// If the root node is not optimal, the IP is infeasible or unbounded
	if *rootNode.SCF.Status != common.SolverStatusOptimal {
		*ip.SCF.Status = *rootNode.SCF.Status
		return nil
	}

	fmt.Printf("[DEBUG] Primal Solution: %v\n", rootNode.SCF.PrimalSolution)

	// Check if the root solution is integer feasible
	if rootNode.IsInteger {
		*ip.SCF.Status = common.SolverStatusOptimal
		return nil
	}

	err = branchAndBound(ip, rootNode)

	ip.SCF.ObjectiveValue = &ip.BestObj
	ip.SCF.PrimalSolution = ip.BestSolution
	return err
}

// isIntegerFeasible checks if a solution is currently integer feasible
func isIntegerFeasible(scf *common.StandardComputationalForm) bool {
	sol := scf.PrimalSolution
	for i := 0; i < sol.Len(); i++ {
		val := sol.AtVec(i)
		// If the current value is not integer and it the index is not in the slack indices
		isSlack := scf.SlackIndices[i]

		if math.Floor(val) != val && isSlack == -1 {
			return false
		}
	}
	return true
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
