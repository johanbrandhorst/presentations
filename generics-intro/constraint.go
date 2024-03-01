package main

// START OMIT
type myConstraint interface {
	~int | ~float64
	GetNumber() int
}

func MyFunc[T myConstraint](t T) int {
	return t.GetNumber()
}

func MyInterfaceFunc(t myConstraint) int { // Compilation error
	return t.GetNumber()
}

// END OMIT
