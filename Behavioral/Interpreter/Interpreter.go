package Interpreter

import (
	"fmt"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func Calculate(o string) (int, error) {

	stack := polishNotationStack{}
	operators := strings.Split(o, " ")

	for _, operatorString := range operators {

		if isOperator(operatorString) {

			right := stack.Pop()
			left := stack.Pop()

			mathFunc := getOperationFunc(operatorString)

			res := mathFunc(left, right)
			stack.Push(res)
		} else {
			val, err := strconv.Atoi(operatorString)

			if err != nil {
				return 0, err
			}

			stack.Push(val)
		}
	}

	return int(stack.Pop()), nil
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return 0
}

func isOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}
	return false
}

func getOperationFunc(o string) func(a, b int) int {

	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}

	case SUB:
		return func(a, b int) int {
			return a - b
		}

	case MUL:
		return func(a, b int) int {
			return a * b
		}

	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}

	return nil
}

/*
	Complex Way to User Interpreter Design Pattern
*/
type Interpreter interface {
	Read() int
	/*
		Value or Symbol
		harus mengextends
	*/
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (s *operationSubtract) Read() int {
	return s.Left.Read() - s.Right.Read()
}

type polishNotationStack []Interpreter

func (p *polishNotationStack) PushX(s Interpreter) {
	*p = append(*p, s)
}

func (p *polishNotationStack) PopX() Interpreter {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}
