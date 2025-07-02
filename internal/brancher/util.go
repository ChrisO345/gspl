package brancher

import "fmt"

func (b *Branch) String() string {
	stringBuilder := "Brancher {\n"
	if b.left != nil {
		stringBuilder += "  Left: " + b.left.String() + "\n"
	} else {
		stringBuilder += "  Left: nil\n"
	}

	if b.right != nil {
		stringBuilder += "  Right: " + b.right.String() + "\n"
	} else {
		stringBuilder += "  Right: nil\n"
	}

	stringBuilder += "  Node: " + fmt.Sprintf("%f", b.node) + "\n"

	stringBuilder += "  BranchStatus: " + b.branchStatus.String() + "\n"
	if b.highestLower != nil {
		stringBuilder += "  HighestLower: " + fmt.Sprintf("%f", *b.highestLower) + "\n"
	} else {
		stringBuilder += "  HighestLower: nil\n"
	}
	if b.lowestUpper != nil {
		stringBuilder += "  LowestUpper: " + fmt.Sprintf("%f", *b.lowestUpper) + "\n"
	} else {
		stringBuilder += "  LowestUpper: nil\n"
	}

	return stringBuilder
}

func (s BranchStatus) String() string {
	switch s {
	case BranchStatusUnexplored:
		return "Unexplored"
	case BranchStatusInfeasible:
		return "Infeasible"
	case BranchStatusFeasible:
		return "Feasible"
	case BranchStatusIncumbent:
		return "Incumbent"
	default:
		return "Unknown Status"
	}
}
