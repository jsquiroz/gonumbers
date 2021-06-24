package matrix_test

import (
	"testing"

	"github.com/jsquiroz/gonumbers/matrix"
	"github.com/stretchr/testify/assert"
)

func TestNewValid(t *testing.T) {
	assert.NotPanics(t, func() { matrix.New(2, 2, []float64{1, 2, 3, 4}) })
}

func TestNewInvalid(t *testing.T) {
	assert.Panics(t, func() { matrix.New(2, 3, []float64{2, 3, 4, 5}) })
}
