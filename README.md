# gspl

**gspl** (short for **Go Simplex Programming Library**, pronounced *gospel*) is a lightweight solver for linear programming problems, built in Go using the revised simplex method. It emphasizes clarity, performance, and seamless integration into Go applications.

---

## Features

* Solves optimal, unbounded, infeasible, and degenerate linear problems.
* Implements an efficient revised simplex algorithm.
* Clean and idiomatic API for modeling and solving LPs.
* Focused on numerical stability and usability.
* Performs basic branch-and-bouund techniques for pure integer problems.

`gspl` is ideal for embedding linear optimization into Go-based software.

---

## Installation

Ensure you have Go installed, then run:

```bash
go get github.com/chriso345/gspl
```

---

## Usage

> [!WARNING]
> Currently undergoing a syntax overhaul. The API will be changing soon. Please check back later.

### What's Happening?

We're setting up a linear program to:

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
variables := []lp.LpVariable{
    lp.NewVariable("x1"),
    lp.NewVariable("x2"),
}
```

You can access and pass them as pointers:

```go
x1 := &variables[0]
```

Forcing integer constraints is done at the variable level:

```go
variables := []lp.LpVariable{
    lp.NewVariable("x1", lp.LpCategoryInteger),
    lp.NewVariable("x2", lp.LpCategoryInteger),
}
```

### Objective Function

Build the objective using terms:

```go
objective := lp.NewExpression([]lp.LpTerm{
    lp.NewTerm(5, *x1),
    lp.NewTerm(3, *x2),
})
```

Add it to the LP:

```go
lp := lp.NewLinearProgram("My LP", variables)
lp.AddObjective(lp.LpMaximise, objective)
```

### Constraints

Each constraint uses an expression, a comparison type, and a right-hand side:

```go
lp.AddConstraint(lp.NewExpression([]lp.LpTerm{
    lp.NewTerm(2, *x1),
    lp.NewTerm(3, *x2),
}), lp.LpConstraintLE, 10)
```

Constraint types:

* `gspl.LpConstraintLE` - less than or equal
* `gspl.LpConstraintGE` - greater than or equal
* `gspl.LpConstraintEQ` - equality

### Solving

```go
solution := solver.Solve(&lp)
solution.PrintSolution()
```

This solves the model and prints variable values and the objective result.

---

## License

Licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
