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
	A
	B
}

// AC is a composition of A and C.
type AC struct {
	A
	C
}

// BC is a composition of B and C.
type BC struct {
	B
	C
}

// ABBC is a composition of AB and BC.
type ABBC struct {
	AB
	BC
}

func newAbbc() *ABBC {
	return &ABBC{
		AB: AB{
			A: A{AI: 1},
			B: B{BI: 2},
		},
		BC: BC{
			B: B{BI: 3},
			C: C{CI: 4},
		},
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
