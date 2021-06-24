package matrix

type matrix struct {
	r    int
	c    int
	data []float64
}

// New, retorna un puntero de matriz, las columnas y las filas deben ser mayores a 0 y
// la longitud de los datos deben corresponder con el producto de las columnas y filas
func New(r, c int, data []float64) *matrix {

	if r <= 0 || c <= 0 {
		panic("Fuera de rango")
	}

	if r*c != len(data) {
		panic("La longitud de la columna y filas no corresponden con la longitud de la informacion")
	}

	return &matrix{
		r:    r,
		c:    c,
		data: data,
	}
}

// Shape regresa la longitud de las filas y columnas de la matriz correspondiente
func (m *matrix) Shape() (r, c int) { return m.r, m.c }
