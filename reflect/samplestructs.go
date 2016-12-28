package reflect

// The A is a simple test struct.
type A struct {
	AI int
}

// B is a simple test struct.
type B struct {
	BI int
}

// C is a simple test struct.
type C struct {
	CI int
}

// AB is a composition of A and B.
type AB struct {
	a A
	b B
}

// AC is a composition of A and C.
type AC struct {
	a A
	c C
}

// BC is a composition of B and C.
type BC struct {
	b B "out"
	c C "in"
}

// ABBC is a composition of AB and BC.
type ABBC struct {
	ab AB
	bc BC
}

func newAb() *AB {
	return &AB{
		a: A{AI: 1},
		b: B{BI: 2},
	}
}

func newBc() *BC {
	return &BC{
		b: B{BI: 3},
		c: C{CI: 4},
	}
}

func newAbbc() *ABBC {
	return &ABBC{
		ab: *newAb(),
		bc: *newBc(),
	}
}

// IFS is a test struct containing an int, a float32, and a string.
type IFS struct {
	I int
	F float32
	S string
}

// BIFS is a tagged test struct composed of a bool and an IFS struct.
type BIFS struct {
	B   bool `direction:"out"`
	IFS `direction:"in"`
}
