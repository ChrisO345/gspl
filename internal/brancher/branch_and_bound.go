package brancher

import (
	"errors"
	"fmt"

	"github.com/chriso345/gspl/internal/common"
	"github.com/chriso345/gspl/internal/simplex"
)

func branchAndBound(ip *common.IntegerProgram, rootNode *common.Node) error {
	nodes, err := branchFunc(rootNode)
	if err != nil {
		return fmt.Errorf("error in branching function: %v", err)
	}

	for _, node := range nodes {
		node.Depth = rootNode.Depth + 1
		fmt.Printf("[DEBUG] Branching to new node at depth %d\n", node.Depth)
		err := simplex.Simplex(node.SCF)
		if err != nil {
			return fmt.Errorf("error solving child node: %v", err)
		}

		if *node.SCF.Status != common.SolverStatusOptimal {
			// Node is infeasible, or unbounded, so it can be pruned
			continue
		}

		node.IsInteger = isIntegerFeasible(node.SCF)
		fmt.Printf("[DEBUG] Node Objective: %.4f, IsInteger: %v\n\n", *node.SCF.ObjectiveValue, node.IsInteger)
		if node.IsInteger {
			objVal := *node.SCF.ObjectiveValue
			if objVal < ip.BestObj {
				ip.BestObj = objVal
				ip.BestSolution = node.SCF.PrimalSolution
				fmt.Printf("[DEBUG] New Best Obj: %.4f\n", ip.BestObj)
			}
			continue
		}

		// If not integer feasible, continue branching
		err = branchAndBound(ip, node)
		if err != nil {
			return err
		}
	}

	return errors.New("Branch and Bound algorithm not yet implemented")
}
