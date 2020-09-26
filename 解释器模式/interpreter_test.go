package Interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	expression := "w x z +"
	sentence := NewEvaluator(expression)
	variables := make(map[string]Expression)
	variables["w"] = &Integer{6}
	variables["x"] = &Integer{10}
	variables["z"] = &Integer{41}
	result := sentence.Interpret(variables)
	fmt.Println(result)
}
