package benchmark

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func measure(name string, f func()) {
	// fmt.Println("start ", name, "...")
	start := time.Now()
	f()
	taken := time.Since(start)
	fmt.Printf("%s | %v\n", name, taken)
}

func Benchmark() {
	// prepare string numbers
	for i := 0; i < MaxIterations; i++ {
		StringNumbers[i] = strconv.Itoa(i)
	}
	// prepare random numbers
	for i := 0; i < MaxIterations; i++ {
		RandomNumbers[i] = strconv.Itoa(rand.Intn(MaxIterations))
	}

	for _, scenario := range TestScenarios {
		for name, method := range Methods {
			measure(
				fmt.Sprintf(
					"%s | %s (%d concurrents) (%d iters)",
					name,
					scenario.scenarioName,
					scenario.concurrents,
					scenario.iterations,
				), func() {
					TestMethods[scenario.scenarioName](scenario.concurrents, scenario.iterations, method[scenario.scenarioName])
				})
		}
	}
}
