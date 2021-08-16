package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func attemptCounter(counter <-chan bool) {
	for range counter {
		count += 1
	}
}

func printSolutions() {
	previousCount := 0
	for {
		//CallClear()
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
