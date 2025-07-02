package lp

import (
	"fmt"
	"math"
	"strings"
)

// String returns a string representation of the LinearProgram
// String returns a string representation of the LinearProgram.
func (lp *LinearProgram) String() string {
	var sb strings.Builder

	sb.WriteString(lp.Description + "\n")

	if lp.Sense == LpMinimise {
		sb.WriteString("Minimize: ")
	} else {
		sb.WriteString("Maximize: ")
	}

	// Objective function
	first := true
	for i, coef := range lp.ObjectiveFunc.RawVector().Data {
		if coef == 0 {
			continue
		}
		if !first {
			if coef > 0 {
				sb.WriteString(" + ")
			} else {
				sb.WriteString(" - ")
			}
		} else {
			if coef < 0 {
				sb.WriteString("-")
			}
			first = false
		}
		varName := lp.VariablesMap[i]
		sb.WriteString(fmt.Sprintf("%.2f * %s", math.Abs(coef), varName.Name))
	}
	sb.WriteString("\n")

	// Constraints
	sb.WriteString("Subject to:\n")
	for row := range lp.Constraints.RawMatrix().Rows {
		sb.WriteString(fmt.Sprintf("  C%d: ", row+1))

		first = true
		for col := range lp.Constraints.RawMatrix().Cols {
			coef := lp.Constraints.At(row, col)
			if coef == 0 {
				continue
			}
			if !first {
				if coef > 0 {
					sb.WriteString(" + ")
				} else {
					sb.WriteString(" - ")
				}
			} else {
				if coef < 0 {
					sb.WriteString("-")
				}
				first = false
			}
			varName := lp.VariablesMap[col]
			sb.WriteString(fmt.Sprintf("%.2f * %s", math.Abs(coef), varName.Name))
		}

		sb.WriteString(" <= ")
		sb.WriteString(fmt.Sprintf("%.3f\n", lp.RHS.AtVec(row)))
	}

	return sb.String()
}

// PrintSolution prints the solution of the linear program in a human-readable format.
func (lp *LinearProgram) PrintSolution() {
	fmt.Println(lp.Status.String())
	fmt.Println(lp.Solution)

	for i, v := range lp.VariablesMap {
		fmt.Printf("%s: %f\n", v.Name, lp.Variables.At(i, 0))
	}
}
