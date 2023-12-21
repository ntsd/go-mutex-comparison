package benchmark

import "sync"

const MaxIterations = 100000

type ScenarioName string

const (
	ScenarioWrite      ScenarioName = "write"
	ScenarioRandomRead ScenarioName = "random_read"
	ScenarioRangeRead  ScenarioName = "range_read"
)

type Scenario struct {
	concurrents  int
	iterations   int
	scenarioName ScenarioName
}

var StringNumbers [MaxIterations]string
var RandomNumbers [MaxIterations]string

var TestMethods = map[ScenarioName]func(concurrents, iterations int, method func(key string)){
	ScenarioWrite: func(concurrents, iterations int, method func(key string)) {
		wg := new(sync.WaitGroup)
		wg.Add(concurrents)
		for t := 0; t < concurrents; t++ {
			go func() {
				for i := 0; i < iterations; i++ {
					method(StringNumbers[i])
				}
				wg.Done()
			}()
		}
		wg.Wait()
	},
	ScenarioRandomRead: func(concurrents, iterations int, method func(key string)) {
		wg := new(sync.WaitGroup)
		wg.Add(concurrents)
		for t := 0; t < concurrents; t++ {
			go func() {
				for i := 0; i < iterations; i++ {
					method(RandomNumbers[i])
				}
				wg.Done()
			}()
		}
		wg.Wait()
	},
	ScenarioRangeRead: func(concurrents, iterations int, method func(key string)) {
		wg := new(sync.WaitGroup)
		wg.Add(concurrents)
		for t := 0; t < concurrents; t++ {
			go func() {
				for i := 0; i < iterations; i++ {
					method("0")
				}
				wg.Done()
			}()
		}
		wg.Wait()
	},
}
var TestScenarios = []Scenario{
	{
		concurrents:  10,
		iterations:   MaxIterations,
		scenarioName: ScenarioWrite,
	},
	{
		concurrents:  10,
		iterations:   MaxIterations,
		scenarioName: ScenarioRandomRead,
	},
	{
		concurrents:  10,
		iterations:   100, // Range read will read all per interation
		scenarioName: ScenarioRangeRead,
	},
	{
		concurrents:  100,
		iterations:   MaxIterations,
		scenarioName: ScenarioWrite,
	},
	{
		concurrents:  100,
		iterations:   MaxIterations,
		scenarioName: ScenarioRandomRead,
	},
	{
		concurrents:  100,
		iterations:   10, // Range read will read all per interation
		scenarioName: ScenarioRangeRead,
	},
}
