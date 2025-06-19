# gspl

**gspl** (short for **Go Simplex Programming Library**, pronounced *gospel*) is a lightweight solver for linear programming problems, built in Go using the revised simplex method. It emphasizes clarity, performance, and seamless integration into Go applications.

---

## Features

* Solves optimal, unbounded, infeasible, and degenerate linear problems.
* Implements an efficient revised simplex algorithm.
* Clean and idiomatic API for modeling and solving LPs.
* Focused on numerical stability and usability.

`gspl` is ideal for embedding linear optimization into Go-based software.

---

## Installation

Ensure you have Go installed, then run:

```bash
go get github.com/chriso345/gspl
```

---

## Usage

Here's how to define and solve a basic linear programming problem with **gspl**:

```go
package main

import "github.com/chriso345/gspl"

func main() {
    // Create decision variables
    variables := []gspl.LpVariable{
        gspl.NewVariable("x1"),
        gspl.NewVariable("x2"),
        gspl.NewVariable("x3"),
    }

    x1 := &variables[0]
    x2 := &variables[1]
    x3 := &variables[2]

    // Objective: Minimize -6*x1 + 7*x2 + 4*x3
    objective := gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(-6, *x1),
        gspl.NewTerm(7, *x2),
        gspl.NewTerm(4, *x3),
    })

    // Initialize the LP with a name and variables
    lp := gspl.NewLinearProgram("Minimisation Example", variables)
    lp.AddObjective(gspl.LpMinimise, objective)

    // Add constraints
    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(2, *x1),
        gspl.NewTerm(5, *x2),
        gspl.NewTerm(-1, *x3),
    }), gspl.LpConstraintLE, 18)

    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(1, *x1),
        gspl.NewTerm(-1, *x2),
        gspl.NewTerm(-2, *x3),
    }), gspl.LpConstraintLE, -14)

    lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
        gspl.NewTerm(3, *x1),
        gspl.NewTerm(2, *x2),
        gspl.NewTerm(2, *x3),
    }), gspl.LpConstraintEQ, 26)

    // Solve the problem and print the solution
    lp.Solve().PrintSolution()
}
```

### What’s Happening?

We’re setting up a linear program to:

1. **Minimize**:
   $-6x_1 + 7x_2 + 4x_3$
2. **Subject to constraints**:

   * $2x_1 + 5x_2 - x_3 \leq 18$
   * $x_1 - x_2 - 2x_3 \leq -14$
   * $3x_1 + 2x_2 + 2x_3 = 26$

The solution will print the optimal values of the variables and the minimized objective value.

See the [examples](examples) directory for other scenarios.

---

## API Overview

### Variables

Create decision variables as a slice:

```go
variables := []gspl.LpVariable{
    gspl.NewVariable("x1"),
    gspl.NewVariable("x2"),
}
```

You can access and pass them as pointers:

```go
x1 := &variables[0]
```

### Objective Function

Build the objective using terms:

```go
objective := gspl.NewExpression([]gspl.LpTerm{
    gspl.NewTerm(5, *x1),
    gspl.NewTerm(3, *x2),
})
```

Add it to the LP:

```go
lp := gspl.NewLinearProgram("My LP", variables)
lp.AddObjective(gspl.LpMaximise, objective)
```

### Constraints

Each constraint uses an expression, a comparison type, and a right-hand side:

```go
lp.AddConstraint(gspl.NewExpression([]gspl.LpTerm{
    gspl.NewTerm(2, *x1),
    gspl.NewTerm(3, *x2),
}), gspl.LpConstraintLE, 10)
```

Constraint types:

* `gspl.LpConstraintLE` – less than or equal
* `gspl.LpConstraintGE` – greater than or equal
* `gspl.LpConstraintEQ` – equality

### Solving

```go
solution := lp.Solve()
solution.PrintSolution()
```

This solves the model and prints variable values and the objective result.

---

## License

Licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
