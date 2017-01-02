package reflect

// The A is a simple test struct.
type A struct {
	AI int
}

// B is a simple test struct.
type B struct {
	E
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

type D struct {
	DI int
}

type E struct {
	D  "out"
	EI int "out"
}

// BC is a composition of B and C.
type BC struct {
	B "out"
	C "in"
}

// ABBC is a composition of AB and BC.
type ABBC struct {
	AB
	BC
}

// ABII is a composition of AB and two integers.
type ABII struct {
	AB
	C int
	D int
}

func newB(i int) *B {
	return &B{
		E:  *newE(),
		BI: i,
	}
}

func newAb() *AB {
	return &AB{
		A: A{AI: 1},
		B: *newB(2),
	}
}

func newBc() *BC {
	return &BC{
		B: *newB(3),
		C: C{CI: 4},
	}
}

func newD() *D {
	return &D{
		DI: 5,
	}
}

func newE() *E {
	return &E{
		D:  *newD(),
		EI: 6,
	}
}

func newAbbc() *ABBC {
	return &ABBC{
		AB: *newAb(),
		BC: *newBc(),
	}
}

func newAbii() *ABII {
	return &ABII{
		AB: *newAb(),
		C:  3,
		D:  4,
	}
}

// IFS is a test struct containing an int, a float32, and a string.
type IFS struct {
	I int     "out"
	F float32 "in"
	S string  "in"
}

// BIFS is a tagged test struct composed of a bool and an IFS struct.
type BIFS struct {
	B   bool "bar"
	IFS `direction:"in"`
}

func newIfs() *IFS {
	return &IFS{
		I: 2,
		F: 3.14,
		S: "hello",
	}
}

func newBifs() *BIFS {
	return &BIFS{
		B:   true,
		IFS: *newIfs(),
	}
}
