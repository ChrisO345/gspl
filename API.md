# Defines the API for **gspl**

```go
lp := NewLP()
lp.AddObjective(LpSense, Constraint)
lp.AddConstraint(Constraint, LpConstraintType, RHS)
lp.Solve()

// Print the solution
fmt.Println(lp.Status)
fmt.Println(lp.Solution)

// May require a switch to a true simplex method, how does solver do this???
lp.GenerateShadows()
```
