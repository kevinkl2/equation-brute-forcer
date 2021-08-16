package main

import (
	"github.com/Knetic/govaluate"
	"log"
)

// TESTING FILE

func main1() {
	expression, err := govaluate.NewEvaluableExpression("h*(d*e*(a + b + c) + f + g) + i")
	if err != nil {
		log.Println(err)
	}

	values := []map[string]interface{}{
		{
			"a": 5, "b": 4, "c": 3, "d": 1.5, "e": 1.5, "f": 3.5, "g": 0, "h": 1.08, "i": 2,
		},
	}

	log.Println(expression.Evaluate(values[0]))
}
