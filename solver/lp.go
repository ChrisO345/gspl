package solver

import (
	"fmt"
	"math"

	"github.com/chriso345/gspl/internal/matrix"
	"github.com/chriso345/gspl/internal/simplex"
	"gonum.org/v1/gonum/mat"
)

type LinearProgram struct {
	// Solution
	Solution float64 // z

	// Matrix Representation
	Variables     *mat.VecDense // x
	ObjectiveFunc *mat.VecDense // c
	Constraints   *mat.Dense    // A
	RHS           *mat.VecDense // b

	// Simplex Internal Variables
	indices  *mat.VecDense // Indices of basic variables
	pivalues *mat.VecDense // Pivot values
	bMatrix  *mat.Dense    // Basis matrix
	cb       *mat.VecDense // Coefficients of the basis variables

	// Others
	Description      string
	VariablesMap     []string
	Status           LpStatus
	Sense            LpSense
	ConstraintVector []LpConstraintType
}

// NewLinearProgram Create a new Linear Program
func NewLinearProgram(desc string, vars []LpVariable) LinearProgram {

	lp := LinearProgram{
		Description:  desc,
		VariablesMap: make([]string, len(vars)),
	}

	for i, v := range vars {
		lp.VariablesMap[i] = v.Name
	}

	lp.Status = LpStatusNotSolved
	return lp
}

// AddObjective Add an objective to the linear program
func (lp *LinearProgram) AddObjective(sense LpSense, objective LpExpression) *LinearProgram {
	lp.Sense = sense
	// lp.ObjectiveFunction = objective
	lp.ObjectiveFunc = mat.NewVecDense(len(objective.Terms), nil)
	for i, v := range objective.Terms {
		mappedIndex := -1
		for j, varName := range lp.VariablesMap {
			if varName == v.Variable.Name {
				mappedIndex = j
				break
			}
		}
		if mappedIndex == -1 {
			panic(fmt.Sprintf("Variable %s not found in Linear Program", v.Variable.Name))
		}
		lp.ObjectiveFunc.SetVec(i, v.Coefficient)
	}

	if lp.Sense == LpMaximise {
		for i := range lp.ObjectiveFunc.RawVector().N {
			lp.ObjectiveFunc.SetVec(i, -lp.ObjectiveFunc.At(i, 0))
		}
	}

	return lp
}

// AddConstraint Add a constraint to the linear program
func (lp *LinearProgram) AddConstraint(constraint LpExpression, constraintType LpConstraintType, rightHandSide float64) *LinearProgram {
	// Panic if objective function is not set
	if lp.ObjectiveFunc == nil {
		panic("Objective function not set")
	}

	lp.ConstraintVector = append(lp.ConstraintVector, constraintType)

	if rightHandSide < 0 {
		// Multiply the constraint by -1, flip equality sign
		rightHandSide = math.Abs(rightHandSide)
		for i := range constraint.Terms {
			constraint.Terms[i].Coefficient *= -1
		}
		constraintType = -constraintType
	}

	currentRow := 0
	if lp.Constraints == nil {
		lp.Constraints = mat.NewDense(1, len(lp.VariablesMap), nil)
		lp.RHS = mat.NewVecDense(1, nil)
	} else {
		currentRow = lp.Constraints.RawMatrix().Rows
		lp.Constraints = matrix.ResizeMatDense(lp.Constraints, currentRow+1, len(lp.VariablesMap))
		lp.RHS = matrix.ResizeVecDense(lp.RHS, currentRow+1)
	}

	newRow := make([]float64, len(lp.VariablesMap))
	for _, v := range constraint.Terms {
		mappedIndex := -1
		for j, varName := range lp.VariablesMap {
			if varName == v.Variable.Name {
				mappedIndex = j
				break
			}
		}
		if mappedIndex == -1 {
			panic(fmt.Sprintf("Variable %s not found in Linear Program", v.Variable.Name))
		}
		newRow[mappedIndex] = v.Coefficient
	}

	lp.Constraints.SetRow(currentRow, newRow)
	lp.RHS.SetVec(currentRow, rightHandSide)

	return lp
}

func (lp *LinearProgram) Solve() *LinearProgram {
	// Add slacks for non-equality constraints
	for i, constraintType := range lp.ConstraintVector {
		if constraintType != LpConstraintEQ {
			slack := NewVariable(fmt.Sprintf("s%d", i))
			lp.VariablesMap = append(lp.VariablesMap, slack.Name)
			unitVector := mat.NewDense(lp.Constraints.RawMatrix().Rows, 1, nil)
			one := 1.0
			if constraintType == LpConstraintGE {
				one = -1.0
			}
			unitVector.Set(i, 0, one)
			lp.Constraints = matrix.ResizeMatDense(lp.Constraints, lp.Constraints.RawMatrix().Rows, len(lp.VariablesMap))

			lp.ObjectiveFunc = matrix.ResizeVecDense(lp.ObjectiveFunc, lp.ObjectiveFunc.RawVector().N+1)
			lp.ObjectiveFunc.SetVec(lp.ObjectiveFunc.RawVector().N-1, 0)
		}
	}

	m := lp.Constraints.RawMatrix().Rows
	n := len(lp.VariablesMap)

	z, x, _, idx, flag := simplex.Simplex(lp.Constraints, lp.RHS, lp.ObjectiveFunc, m, n)

	lp.Solution = z
	lp.Variables = x
	lp.indices = idx

	switch flag {
	case 0:
		lp.Status = LpStatusOptimal
	case 1:
		lp.Status = LpStatusInfeasible
	case -1:
		lp.Status = LpStatusUnbounded
	}

	if lp.Sense == LpMaximise {
		lp.Solution *= -1
	}

	return lp
}

func (lp *LinearProgram) PrintSolution() {
	fmt.Println(lp.Status.String())
	fmt.Println(lp.Solution)

	for i, v := range lp.VariablesMap {
		fmt.Printf("%s: %f\n", v, lp.Variables.At(i, 0))
	}
}

/* #####################################################################################################################
TO BE MOVED TO SEPARATE FILES, SOME OF THE STRUCTS ARE PRIVATE
##################################################################################################################### */

type LpExpression struct {
	Terms []LpTerm
}

func NewExpression(terms []LpTerm) LpExpression {
	return LpExpression{terms}
}

type LpTerm struct {
	Coefficient float64
	Variable    LpVariable // These get added to the variable list in the LinearProgram??
}

func NewTerm(coefficient float64, variable LpVariable) LpTerm {
	return LpTerm{coefficient, variable}
}

type LpVariable struct {
	Name         string
	Value        float64
	IsSlack      bool
	IsArtificial bool
}

func NewVariable(name string) LpVariable {
	return LpVariable{name, 0, false, false}
}

// LpCategory Note that this is currently not used
type LpCategory string

const (
	LpContinuous = LpCategory("Continuous")
	LpInteger    = LpCategory("Integer")
	LpBinary     = LpCategory("Binary")
)

type LpSense int

const (
	LpMinimise = LpSense(-1)
	LpMaximise = LpSense(1)
)

type LpStatus int

const (
	LpStatusNotSolved      = LpStatus(0)
	LpStatusOptimal        = LpStatus(1)
	LpStatusInfeasible     = LpStatus(2)
	LpStatusUnbounded      = LpStatus(3)
	LpStatusUndefined      = LpStatus(4)
	LpStatusNotImplemented = LpStatus(5)
)

var LpStatusMap = map[LpStatus]string{
	LpStatusNotSolved:      "Not Solved",
	LpStatusOptimal:        "Optimal",
	LpStatusInfeasible:     "Infeasible",
	LpStatusUnbounded:      "Unbounded",
	LpStatusUndefined:      "Undefined",
	LpStatusNotImplemented: "Not Implemented",
}

func (s *LpStatus) String() string {
	return LpStatusMap[*s]
}

type LpConstraintType int

const (
	LpConstraintLE = LpConstraintType(-1)
	LpConstraintEQ = LpConstraintType(0)
	LpConstraintGE = LpConstraintType(1)
)

// Utils

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
