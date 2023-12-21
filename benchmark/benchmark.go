package benchmark

import (
	"fmt"
	"math/rand"
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
	// pre add random number
	for i := 0; i < MaxIterations; i++ {
		RandomNumbers[i] = rand.Intn(MaxIterations)
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
