package main

import (
	"github.com/Knetic/govaluate"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

func generateEquation() []string {
	tempVariables := make([]string, len(variables))
	copy(tempVariables, variables)

	var equation []string
	parenthesesCounter := len(variables)
	openParenthesesCounter := 0

	rand.Shuffle(len(tempVariables), func(i, j int) {
		tempVariables[i], tempVariables[j] = tempVariables[j], tempVariables[i]
	})

	for _, i := range tempVariables {
		if parenthesesCounter > 0 {
			openParens := 1
			if parenthesesCounter != 1 {
				openParens = rand.Intn(parenthesesCounter-1) + 1
			}
			for a := 0; a < openParens; a++ {
				equation = append(equation, "(")
			}
			openParenthesesCounter += openParens
			parenthesesCounter -= openParens
		}

		equation = append(equation, i)

		if openParenthesesCounter > 0 {
			closedParens := rand.Intn(openParenthesesCounter + 1)
			for a := 0; a < closedParens; a++ {
				equation = append(equation, ")")
			}
			openParenthesesCounter -= closedParens
		}

		if i != tempVariables[len(tempVariables)-1] {
			op := operators[rand.Intn(len(operators))]
			equation = append(equation, op)
		} else {
			for a := 0; a < openParenthesesCounter; a++ {
				equation = append(equation, ")")
			}
		}
	}

	return equation
}

func validateEquation(equation []string) bool {
	expression, err := govaluate.NewEvaluableExpression(strings.Join(equation, ""))
	if err != nil {
		log.Println(err)
	}

	for i, _ := range values {
		result, err := expression.Evaluate(values[i])
		if err != nil {
			log.Println(err)
		}

		if result != goals[i] {
			return false
		}
	}

	return true
}

func simplifyEquation(equations <-chan []string) {
	for equation := range equations {
		log.Println(equation)
		cmd := exec.Command("python", "-u", "./simplify.py", strings.Join(equation, ""))
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
		}
		solutions[string(out)] = true
	}
}
