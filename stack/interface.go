package stack

// Stack interface
type Stack interface {
	Push(any)
	Pop() (any, error)
	Peek() (any, error)
	IsEmpty() bool
	Size() int
}
