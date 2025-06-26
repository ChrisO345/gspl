// Package lp provides foundational types and utilities for defining and
// manipulating linear programming (LP) models.
//
// It defines core abstractions such as variables, expressions, constraints,
// objective functions, and problem categories (e.g., continuous, integer).
//
// The package supports constructing LP problems programmatically using
// types like LinearProgram, LpVariable, LpTerm, and LpExpression.
//
// Example usage:
//
//	vars := []lp.LpVariable{
//	    lp.NewVariable("x1"),
//	    lp.NewVariable("x2"),
//	}
//	prog := lp.NewLinearProgram("Example LP", vars)
package lp
