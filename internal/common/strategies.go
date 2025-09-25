package common

// Branching: takes a node, returns two child nodes (or more if you generalize).
type BranchFunc func(node *Node) ([]*Node, error)

// Heuristic: try to find a feasible integer solution quickly.
type HeuristicFunc func(node *Node) ([]float64, float64, bool)

// Cutting planes: generate additional constraints for a node.
type CutFunc func(node *Node) [][]float64
