package main

// Color struct holds three color values
type Color struct {
	r, g, b float64
}

// AddColor adds color values from other color values
func (c Color) AddColor(c1 Color) Color {
	return Color{c.r + c1.r, c.g + c1.g, c.b + c1.b}
}

// SubtractColor subtracts color values from other color values
func (c Color) SubtractColor(c1 Color) Color {
	return Color{c.r - c1.r, c.g - c1.g, c.b - c1.b}
}

// MulScalarColor multiplies color values by a scalar
func (c Color) MulScalarColor(s float64) Color {
	return Color{c.r * s, c.g * s, c.b * s}
}

// DivScalarColor divides color values by a scalar
func (c Color) DivScalarColor(s float64) Color {
	return Color{c.r / s, c.g / s, c.b / s}
}

// MulColor returns Hadamard product of two colors
func (c Color) MulColor(c1 Color) Color {
	return Color{c.r * c1.r, c.g * c1.g, c.b * c1.b}
}
