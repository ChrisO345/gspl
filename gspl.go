package gspl

func Gspl() {
	// Objective: Minimise - 6 * x1 + 7 * x2 + 4 * x3
	// Constraints: 2 * x1 + 5 * x2 - 1 * x3 <= 18
	// Constraints: 1 * x1 - 1 * x2 - 2 * x3 <= -14
	// Constraints: 3 * x1 + 2 * x2 + 2 * x3 = 26

	variables := []LpVariable{
		NewVariable("x1"),
		NewVariable("x2"),
		NewVariable("x3"),
		NewVariable("x4"),
		NewVariable("x5"),
	}

	terms := []LpTerm{
		NewTerm(15, variables[0]),
		NewTerm(10, variables[1]),
		NewTerm(-10, variables[2]),
		NewTerm(1, variables[3]),
		NewTerm(2, variables[4]),
	}
	objective := NewExpression(terms)

	terms2 := []LpTerm{
		NewTerm(1, variables[0]),
		NewTerm(1, variables[1]),
		NewTerm(1, variables[2]),
		NewTerm(1, variables[3]),
		NewTerm(0, variables[4]),
	}

	terms3 := []LpTerm{
		NewTerm(0, variables[0]),
		NewTerm(-1, variables[1]),
		NewTerm(0, variables[2]),
		NewTerm(-1, variables[3]),
		NewTerm(1, variables[4]),
	}

	terms4 := []LpTerm{
		NewTerm(1, variables[0]),
		NewTerm(0, variables[1]),
		NewTerm(1, variables[2]),
		NewTerm(0, variables[3]),
		NewTerm(1, variables[4]),
	}

	lp := NewLinearProgram()
	lp.AddObjective(LpMinimise, objective).
		AddConstraint(NewExpression(terms2), LpConstraintEQ, 4).
		AddConstraint(NewExpression(terms3), LpConstraintEQ, 4).
		AddConstraint(NewExpression(terms4), LpConstraintEQ, 8)

	//fmt.Println(lp.String())

	lp.Solve().PrintSolution()

	// lp.Shadows()
}
