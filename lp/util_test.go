package lp

import (
	"bytes"
	"os"
	"testing"

	"github.com/chriso345/gore/assert"
	"gonum.org/v1/gonum/mat"
)

func TestPrintSolution(t *testing.T) {
	x1 := NewVariable("x1")
	x2 := NewVariable("x2")
	lp := NewLinearProgram("Test LP", []LpVariable{x1, x2})

	lp.AddObjective(LpMinimise, NewExpression([]LpTerm{
		NewTerm(2, x1),
		NewTerm(3, x2),
	}))

	lp.AddConstraint(NewExpression([]LpTerm{
		NewTerm(1, x1),
		NewTerm(1, x2),
	}), LpConstraintLE, 5)

	lp.PrimalSolution = mat.NewVecDense(3, []float64{1, 2, 0})
	lp.ObjectiveValue = 8.0

	// Redirect stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	lp.PrintSolution()

	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = oldStdout

	assert.True(t, buf.Len() > 0)
	assert.StringContains(t, buf.String(), "Test LP")
	assert.StringContains(t, buf.String(), "ObjectiveValue")
	assert.StringContains(t, buf.String(), "x1")
	assert.StringContains(t, buf.String(), "x2")
}
