package benchmark

import (
	"sync"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	syncMap    sync.Map
	mutexMap   = make(map[string]string)
	mutex      sync.Mutex
	rwMutexMap = make(map[string]string)
	rwMutex    sync.RWMutex
	cMap       = cmap.New[string]()
)

var Methods = map[string]map[ScenarioName]func(key string){
	"sync.Map": {
		ScenarioWrite: (func(key string) {
			syncMap.Store(key, key)
		}),
		ScenarioRandomRead: (func(key string) {
			v, ok := syncMap.Load(key)
			if !ok {
			}
			v = v
		}),
		ScenarioRangeRead: (func(key string) {
			syncMap.Range(func(k, v interface{}) bool {
				k = k
				v = v
				return true
			})
		}),
	},
	"orcaman/concurrent-map": {
		ScenarioWrite: (func(key string) {
			cMap.Set(key, key)
		}),
		ScenarioRandomRead: (func(key string) {
			v, ok := cMap.Get(key)
			if !ok {
			}
			v = v
		}),
		ScenarioRangeRead: (func(key string) {
			for i := range cMap.IterBuffered() {
				i = i
			}
		}),
	},
	"Mutex": {
		ScenarioWrite: (func(key string) {
			mutex.Lock()
			mutexMap[key] = key
			mutex.Unlock()
		}),
		ScenarioRandomRead: (func(key string) {
			mutex.Lock()
			v, ok := mutexMap[key]
			if !ok {
			}
			v = v
			mutex.Unlock()
		}),
		ScenarioRangeRead: (func(key string) {
			mutex.Lock()
			for k, v := range mutexMap {
				k = k
				v = v
			}
			mutex.Unlock()
		}),
	},
	"RWMutex": {
		ScenarioWrite: (func(key string) {
			rwMutex.Lock()
			rwMutexMap[key] = key
			rwMutex.Unlock()
		}),
		ScenarioRandomRead: (func(key string) {
			rwMutex.RLock()
			v, ok := rwMutexMap[key]
			if !ok {
			}
			v = v
			rwMutex.RUnlock()
		}),
		ScenarioRangeRead: (func(key string) {
			rwMutex.RLock()
			for k, v := range rwMutexMap {
				k = k
				v = v
			}
			rwMutex.RUnlock()
		}),
	},
}
