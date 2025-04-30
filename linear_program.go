package gspl

import (
	"fmt"
	"math"
)

type LinearProgram struct {
	// Solution
	Solution float64 // z

	// Matrix Representation
	Variables     *Matrix // x
	ObjectiveFunc *Matrix // c
	Constraints   *Matrix // A
	RHS           *Matrix // b

	// Simplex Internal Variables
	indices  *Matrix
	pivalues *Matrix
	bMatrix  *Matrix
	cb       *Matrix

	// Others
	Description  string
	VariablesMap []string
	Status       LpStatus
	Sense        LpSense
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
	if sense != LpMinimise {
		panic("Not Implemented, use LpMinimise")
	}
	lp.Sense = sense
	// lp.ObjectiveFunction = objective
	lp.ObjectiveFunc = NewMatrix(len(objective.Terms), 1)
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
		lp.ObjectiveFunc.Set(i, 0, v.Coefficient)
	}

	return lp
}

// AddConstraint Add a constraint to the linear program
func (lp *LinearProgram) AddConstraint(constraint LpExpression, constraintType LpConstraintType, rightHandSide float64) *LinearProgram {
	// Panic if objective function is not set
	if lp.ObjectiveFunc == nil {
		panic("Objective function not set")
	}

	if constraintType != LpConstraintEQ {
		panic("Only LpConstraintEQ is supported")
	}

	if rightHandSide < 0 {
		// Multiply the constraint by -1, flip equality sign
		rightHandSide = math.Abs(rightHandSide)
		for i := range constraint.Terms {
			constraint.Terms[i].Coefficient *= -1
		}
		constraintType = -constraintType
	}

	// Add Artificial Variables
	// if constraintType == LpConstraintEQ || constraintType == LpConstraintGE {
	// 	variable := NewArtificialVariable(fmt.Sprintf("a%d", len(lp.Constraints)+1))
	// 	constraint.Terms = append(constraint.Terms, NewTerm(1, variable))
	// 	lp.ObjectiveFunction.Terms = append(lp.ObjectiveFunction.Terms, NewTerm(-1e20, variable))
	// }

	// Add Slack Variables
	// if constraintType == LpConstraintLE || constraintType == LpConstraintGE {
	// 	variable := NewSlackVariable(fmt.Sprintf("s%d", len(lp.Constraints)+1))
	// 	sign := 1.0
	// 	if constraintType == LpConstraintGE {
	// 		sign = -1.0
	// 	}
	// 	constraint.Terms = append(constraint.Terms, NewTerm(sign, variable))
	// 	lp.ObjectiveFunction.Terms = append(lp.ObjectiveFunction.Terms, NewTerm(0, variable))
	// 	constraintType = LpConstraintEQ
	// }

	// lp.Constraints = append(lp.Constraints, _constraint{constraintType, constraint.Terms, rightHandSide})
	currentRow := 0
	if lp.Constraints == nil {
		lp.Constraints = NewMatrix(1, len(lp.VariablesMap))
		lp.RHS = NewMatrix(1, 1)
	} else {
		currentRow = lp.Constraints.Rows
		lp.Constraints.Resize(currentRow+1, len(lp.VariablesMap))
		lp.RHS.Resize(currentRow+1, 1)
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
	lp.RHS.Set(currentRow, 0, rightHandSide)

	return lp
}

func (lp *LinearProgram) Solve() *LinearProgram {
	m := len(lp.Constraints.Values)
	n := len(lp.VariablesMap)

	z, x, _, idx, flag := Simplex(lp.Constraints, lp.RHS, lp.ObjectiveFunc, m, n)

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

	return lp
}

func (lp *LinearProgram) PrintSolution() {
	fmt.Println(lp.Status.String())
	fmt.Println(lp.Solution)

	for i, v := range lp.VariablesMap {
		fmt.Printf("%s: %f\n", v, lp.Variables.Get(i, 0))
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

func NewSlackVariable(name string) LpVariable {
	return LpVariable{name, 0, true, false}
}

func NewArtificialVariable(name string) LpVariable {
	return LpVariable{name, 0, false, true}
}

type _constraint struct {
	ConstraintType LpConstraintType
	Terms          []LpTerm
	RightHandSide  float64
}

/* #####################################################################################################################
TO BE MOVED TO SEPARATE FILES
##################################################################################################################### */

// LpCategory Note that this is currently not used
type LpCategory string

const (
	LpContinuous = LpCategory("Continuous")
	LpInteger    = LpCategory("Integer")
	LpBinary     = LpCategory("Binary")
)

// TODO: Category Mapping ?? dictionary??

type LpSense int

const (
	LpMinimise = LpSense(-1)
	LpMaximise = LpSense(1)
)

// TODO: Sense Mapping ?? dictionary??

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
