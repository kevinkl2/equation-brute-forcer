package main

import (
	"github.com/Knetic/govaluate"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	variables []string
	values    []map[string]interface{}
	operators []string
	goals     []float64
	clear     map[string]func()
	count     int
	solutions map[string]bool
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

func findSolutions(counter chan<- bool, equations chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		equation := generateEquation()
		counter <- true
		if validateEquation(equation) == true {
			equations <- equation
		}
	}
}

func attemptCounter(counter <-chan bool) {
	for range counter {
		count += 1
	}
}

func printSolutions() {
	previousCount := 0
	for {
		CallClear()
		log.Printf("Attempts: %v, %v/s", count, count-previousCount)
		previousCount = count

		var keys []string
		for solution, _ := range solutions {
			keys = append(keys, solution)
		}
		log.Printf("Found Solutions: %v", strings.Join(keys, ","))

		time.Sleep(1 * time.Second)
	}
}

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	//start := time.Now()
	counter := make(chan bool, 15)
	equations := make(chan []string, 15)
	solutions = make(map[string]bool)
	//variables = []string{"baseDamage", "smite", "vigor", "agilityConst", "agilityMult"}
	//values = []map[string]interface{}{
	//	{
	//		"baseDamage":   3,
	//		"smite":        5,
	//		"vigor":        1.02,
	//		"agilityConst": 1,
	//		"agilityMult":  1.1,
	//	},
	//}
	variables = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	values = []map[string]interface{}{
		{
			"a": 10, "b": 5, "c": 3, "d": 1.5, "e": 1.75, "f": 15, "g": 20, "h": 1.02, "i": 2,
		},
		{
			"a": 5, "b": 4, "c": 3, "d": 1.5, "e": 1.5, "f": 3.5, "g": 0, "h": 1.08, "i": 2,
		},
	}
	operators = []string{"+", "*"}
	//goals = []float64{10.076}
	goals = []float64{85.895, 35.33}

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go findSolutions(counter, equations, &wg)
	}

	go attemptCounter(counter)
	go simplifyEquation(equations)
	go printSolutions()

	wg.Wait()

	//duration := time.Since(start)
	//log.Println(duration)
}
