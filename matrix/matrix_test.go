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

func TestSize(t *testing.T) {

	row := 2
	col := 3

	a := matrix.New(row, col, []float64{1, 2, 2, 2, 3, 4})
	total := a.Size()

	assert.Equal(t, row*col, total, "Deberia de regresar el producto de las filas y columnas")
}
