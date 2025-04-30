package gspl

import (
	"fmt"
)

func (lp *LinearProgram) String() string {
	stringBuilder := ""
	if lp.Sense == LpMinimise {
		stringBuilder += "Min: "
	} else {
		stringBuilder += "Max: "
	}

	stringBuilder += "\n"

	stringBuilder += "Objective: "
	stringBuilder += lp.ObjectiveFunc.String()

	stringBuilder += "\n"

	stringBuilder += "Constraints: "
	for i, v := range lp.Constraints.Values {
		stringBuilder += fmt.Sprintf("C%d: ", i)
		for j, val := range v {
			if val != 0 {
				stringBuilder += fmt.Sprintf("%f * %s + ", val, lp.VariablesMap[j])
			}
		}
		stringBuilder += fmt.Sprintf("= %f\n", lp.RHS.Values[i][0])
	}

	return stringBuilder
}
