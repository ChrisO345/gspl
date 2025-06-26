package lp

// LpExpression represents the LHS of a linear expression
type LpExpression struct {
	Terms []LpTerm
}

// NewExpression creates a new LpExpression with the given terms
func NewExpression(terms []LpTerm) LpExpression {
	return LpExpression{terms}
}

// LpTerm represents a term in a linear expression, consisting of a coefficient and a variable.
type LpTerm struct {
	Coefficient float64
	Variable    LpVariable // These get added to the variable list in the LinearProgram??
}

// NewTerm creates a new LpTerm with the given coefficient and variable.
func NewTerm(coefficient float64, variable LpVariable) LpTerm {
	return LpTerm{coefficient, variable}
}

// LpVariable represents a variable in a linear programming problem.
type LpVariable struct {
	Name         string
	Value        float64
	IsSlack      bool
	IsArtificial bool
}

// NewVariable creates a new LpVariable with the given name.
func NewVariable(name string) LpVariable {
	return LpVariable{name, 0, false, false}
}

// LpCategory represents the category of a linear programming variable, such as continuous, integer, or binary.
type LpCategory string

const (
	LpContinuous = LpCategory("Continuous")
	LpInteger    = LpCategory("Integer") // TODO: Implement integer solving
	LpBinary     = LpCategory("Binary")  // TODO: Implement binary solving
)

// LpSense represents the sense of the linear programming problem, either minimization or maximization.
type LpSense int

const (
	LpMinimise = LpSense(-1)
	LpMaximise = LpSense(1)
)

// LpStatus represents the current status of solving the linear programming problem.
type LpStatus int

const (
	LpStatusNotSolved      = LpStatus(0)
	LpStatusOptimal        = LpStatus(1)
	LpStatusInfeasible     = LpStatus(2)
	LpStatusUnbounded      = LpStatus(3)
	LpStatusUndefined      = LpStatus(4)
	LpStatusNotImplemented = LpStatus(5)
)

// LpStatusMap maps LpStatus values to their string representations.
var LpStatusMap = map[LpStatus]string{
	LpStatusNotSolved:      "Not Solved",
	LpStatusOptimal:        "Optimal",
	LpStatusInfeasible:     "Infeasible",
	LpStatusUnbounded:      "Unbounded",
	LpStatusUndefined:      "Undefined",
	LpStatusNotImplemented: "Not Implemented",
}

// String returns the string representation of the LpStatus.
func (s *LpStatus) String() string {
	return LpStatusMap[*s]
}

// LpConstraintType represents the type of a constraint in a linear programming problem.
type LpConstraintType int

const (
	LpConstraintLE = LpConstraintType(-1)
	LpConstraintEQ = LpConstraintType(0)
	LpConstraintGE = LpConstraintType(1)
)
