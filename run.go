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
	workerCount := 4
	counter := make(chan bool, workerCount*5)
	equations := make(chan []string, workerCount*5)
	solutions = make(map[string]bool, workerCount*5)
	variables = []string{"a", "b"}
	values = []map[string]interface{}{
		{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 0,
		},
	}
	operators = []string{"**"}
	goals = []float64{2}

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
