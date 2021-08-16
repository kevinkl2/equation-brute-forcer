package main

import (
	"github.com/Knetic/govaluate"
	"log"
)

// TESTING FILE

func main() {
	expression, err := govaluate.NewEvaluableExpression("h*(d*e*(a + b + c) + f + g) + i")
	if err != nil {
		log.Println(err)
	}

	values := []map[string]interface{}{
		{
			"a":10, "b":5, "c":3, "d":1.5, "e":1.75, "f":15, "g":20, "h":1.02, "i":2,
		},
	}

	log.Println(expression.Evaluate(values[0]))
}
