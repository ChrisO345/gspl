package simplex

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestSimplexMethodStruct(t *testing.T) {
	sm := &simplexMethod{
		m: 2,
		n: 3,
	}
	sm.flag = 42
	assert.Equal(t, sm.m, 2)
	assert.Equal(t, sm.n, 3)
	assert.Equal(t, sm.flag, 42)
}
