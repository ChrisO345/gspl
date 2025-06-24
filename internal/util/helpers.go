package util

//
// import (
// 	"fmt"
// 	"github.com/chriso345/gspl/solver"
// )
//
// func (lp *solver.LinearProgram) String() string {
// 	stringBuilder := lp.Description
// 	stringBuilder += "\n"
// 	if lp.Sense == LpMinimise {
// 		stringBuilder += "Min: "
// 	} else {
// 		stringBuilder += "Max: "
// 	}
//
// 	stringBuilder += "\n"
//
// 	stringBuilder += "Objective: "
// 	stringBuilder += lp.ObjectiveFunc.String()
//
// 	stringBuilder += "\n"
//
// 	stringBuilder += "Constraints: \n"
// 	for i, v := range lp.Constraints.Values {
// 		stringBuilder += fmt.Sprintf("C%d: ", i)
// 		for j, val := range v {
// 			if val != 0 {
// 				stringBuilder += fmt.Sprintf("%f * %s + ", val, lp.VariablesMap[j])
// 			}
// 		}
// 		stringBuilder += fmt.Sprintf("= %f\n", lp.RHS.Values[i][0])
// 	}
//
// 	return stringBuilder
// }
