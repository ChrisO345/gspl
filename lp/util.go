package lp

import (
	"fmt"
	"math"
)

func (lp *LinearProgram) String() string {
	stringBuilder := lp.Description
	stringBuilder += "\n"
	if lp.Sense == LpMinimise {
		stringBuilder += "Min: "
	} else {
		stringBuilder += "Max: "
	}

	stringBuilder += "\n"

	stringBuilder += "Objective: "
	for i, v := range lp.ObjectiveFunc.RawVector().Data {
		if v != 0 {
			if i > 0 && v > 0 {
				stringBuilder += " + "
			} else if v < 0 {
				stringBuilder += " - "
			}
			stringBuilder += fmt.Sprintf("%f * %s", math.Abs(v), lp.VariablesMap[i])
		}
	}

	stringBuilder += "\n"

	stringBuilder += "Constraints: \n"
	for i, val := range lp.Constraints.RawMatrix().Data {
		stringBuilder += fmt.Sprintf("C%d: ", i)
		if val != 0 {
			if i > 0 && val > 0 {
				stringBuilder += " + "
			} else if val < 0 {
				stringBuilder += " - "
			}
			stringBuilder += fmt.Sprintf("%f * %s", math.Abs(val), lp.VariablesMap[i])
		}
	}

	return stringBuilder
}

func (lp *LinearProgram) PrintSolution() {
	fmt.Println(lp.Status.String())
	fmt.Println(lp.Solution)

	for i, v := range lp.VariablesMap {
		fmt.Printf("%s: %f\n", v, lp.Variables.At(i, 0))
	}
}
