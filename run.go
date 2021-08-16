package main

import (
	"math/rand"
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

func main() {
	workerCount := 1
	counter := make(chan bool, workerCount*5)
	equations := make(chan []string, workerCount*5)
	solutions = make(map[string]bool, workerCount*5)
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
	goals = []float64{85.895, 34.940000000000005}

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go findSolutions(counter, equations, &wg)
	}

	go attemptCounter(counter)
	go simplifyEquation(equations)
	go printSolutions()

	wg.Wait()
}
