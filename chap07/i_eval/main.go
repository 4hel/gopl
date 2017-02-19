package main

import (
	"./eval"
	"fmt"
	"os"
)

func main() {

	env := map[eval.Var]float64{"x": 2}
	exp, err := eval.Parse("5 + pow(x,3)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing expression: %v", err)
	}
	fmt.Println(exp.Eval(env))

}
