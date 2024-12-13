package forms

import "math"

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	c := a * a
	d := b * b
	n := math.Sqrt(c + d)
	return n
}
