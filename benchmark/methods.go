package benchmark

import "sync"

var (
	syncMap    sync.Map
	mutexMap   = make(map[int]int)
	mutex      sync.Mutex
	rwMutexMap = make(map[int]int)
	rwMutex    sync.RWMutex
)

var Methods = map[string]map[ScenarioName]func(key int){
	"sync.Map": {
		ScenarioWrite: (func(key int) {
			syncMap.Store(key, key)
		}),
		ScenarioRandomRead: (func(key int) {
			v, ok := syncMap.Load(key)
			if !ok {
			}
			v = v
		}),
		ScenarioRangeRead: (func(key int) {
			syncMap.Range(func(k, v interface{}) bool {
				k = k
				v = v
				return true
			})
		}),
	},
	"Mutex": {
		ScenarioWrite: (func(key int) {
			mutex.Lock()
			mutexMap[key] = key
			mutex.Unlock()
		}),
		ScenarioRandomRead: (func(key int) {
			mutex.Lock()
			v, ok := mutexMap[key]
			if !ok {
			}
			v = v
			mutex.Unlock()
		}),
		ScenarioRangeRead: (func(key int) {
			mutex.Lock()
			for k, v := range mutexMap {
				k = k
				v = v
			}
			mutex.Unlock()
		}),
	},
	"RWMutex": {
		ScenarioWrite: (func(key int) {
			rwMutex.Lock()
			rwMutexMap[key] = key
			rwMutex.Unlock()
		}),
		ScenarioRandomRead: (func(key int) {
			rwMutex.RLock()
			v, ok := rwMutexMap[key]
			if !ok {
			}
			v = v
			rwMutex.RUnlock()
		}),
		ScenarioRangeRead: (func(key int) {
			rwMutex.RLock()
			for k, v := range rwMutexMap {
				k = k
				v = v
			}
			rwMutex.RUnlock()
		}),
	},
}
