package main

type TrianglePosition struct {
	vertex0, vertex1, vertex2 Tuple
}

type Triangle struct {
	position TrianglePosition
	material Material
	normal   Tuple
}
