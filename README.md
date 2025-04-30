# gspl

**gspl** (short for **Go Simplex Programming Library**, pronounced *gospel*) is a lightweight solver for linear programming problems, based on the revised simplex method and implemented in Go. It is designed for clarity, numerical stability, and seamless integration into Go applications.
___

## Features

- Solves optimal, unbounded, infeasible and degenerate linear problems.
- Implements an effective revised simplex algorithm.
- Clean and simple API for defining and solving linear problems.
- Emphasis on numerical stability and performance.

`gspl` provides a solid foundation for applications requiring reliable and efficient linear optimisation within the Go ecosystem.
___

## Installation

To install `gspl`, you need to have Go installed on your machine. Then, you can run the following command:

```bash
go get github.com/chriso345/gspl
```

___

## Usage

Lets set up a basic linear programming problem using **gspl**:

```go
package main

import "github.com/chriso345/gspl"

func main() {
    // Create decision variables
    x1 := gspl.NewVariable("x1")
    x2 := gspl.NewVariable("x2")
    x3 := gspl.NewVariable("x3")

    // Objective function: Minimize -6 * x1 + 7 * x2 + 4 * x3
    objective := gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(-6, x1),
        gspl.NewTerm(7, x2),
        gspl.NewTerm(4, x3),
    })

    // Set up the LP problem
    lp := gspl.NewLinearProgram()
    lp.AddObjective(gspl.LpMinimise, objective)

    // Add constraints
    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(2, x1),
        gspl.NewTerm(5, x2),
        gspl.NewTerm(-1, x3),
    }), gspl.LpConstraintLE, 18)

    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(1, x1),
        gspl.NewTerm(-1, x2),
        gspl.NewTerm(-2, x3),
    }), gspl.LpConstraintLE, -14)

    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(3, x1),
        gspl.NewTerm(2, x2),
        gspl.NewTerm(2, x3),
    }), gspl.LpConstraintEQ, 26)

    // Solve it
    lp.Solve().PrintSolution()
}
```

### What's Happening Here?

We're defining a linear programming problem where:

1. **Objective Function**: We want to minimize $-6x_1 + 7x_2 + 4x_3$
2. **Constraints**:
    - $2x_1 + 5x_2 - x_3 \leq 18$
    - $x_1 - x_2 - 2x_3 \leq -14$
    - $3x_1 + 2x_2 + 2x_3 = 26$

When we run this program, `gspl` will solve the linear programming problem and print the solution, and optimal values for $x_1$, $x_2$, and $x_3$.

___

## API Overview

### Decision Variables

These are the variables whose values we wish to determine.

```go
x1 := gspl.NewVariable("x1")
x2 := gspl.NewVariable("x2")    
```

Each variable is created using `gspl.NewVariable()` with a name that uniquely identifies it. In this case `x1` and `x2` represent the decision variables.

### Objective Function

The objective function is the mathematical expression that we want to minimize or maximize.

```go
objective := gspl.NewExpression([]gspl.LpTerm{
    gspl.NewTerm(-6, x1),
    gspl.NewTerm(7, x2),
    gspl.NewTerm(4, x3),
})
```

- We use `gspl.NewTerm()` to specify each term in the objective function. Each term consists of:
  - A coefficient (the weight or importance of the variable in the objective function)
  - A variable (the decision variable it corresponds to)
- We then use `gspl.NewExpression()` to create the objective function using the terms we defined.

```go
lp := gspl.NewLinearProgram()
lp.AddObjective(gspl.LpMinimise, objective)
```

- After creating the objective function, we add it to the linear program using `lp.AddObjective()`.
- The first argument specifies whether we want to minimize or maximize the objective function. We use `gspl.LpMinimise` or `gspl.LpMaximise` for this.

### Constraints

Constraints are the conditions that the decision variables must satisfy. They can be equality or inequality constraints.

```go
lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
    gspl.NewTerm(2, x1),
    gspl.NewTerm(5, x2),
    gspl.NewTerm(-1, x3),
}), gspl.LpConstraintLE, 18)
```

Each constraint is composed of:
- **Expression**: Defined similarly to the objective function using `gspl.NewTerm()` and `gspl.NewExpression()`.
- **Type**: The type of constraint. We use `gspl.LpConstraintLE` for less than or equal to ($\leq$), `gspl.LpConstraintGE` for greater than or equal to ($\qeq$), and `gspl.LpConstraintEQ` for equality ($\eq$).
- **Right-hand Side**: The value on the right-hand side of the constraint.

### Solving the Problem

```go
solution := lp.Solve()
solution.PrintSolution()
```

- `lp.Solve()` will run the simplex algorithm to find the optimal solution based on the objective function and constraints.
- `solution.PrintSolution()` will print the optimal values of the decision variables and the optimal value of the objective function.

___ 

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
