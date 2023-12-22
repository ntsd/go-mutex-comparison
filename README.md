# Go Mutex Comparision

Mutex vs RWMutex vs sync.Map, Concurrency map benchmark in different concurrent scenarios.

## Methods

Check all the methods and code here [benchmark/methods.go](benchmark/methods.go)

### 1. sync.Map

Write

```go
syncMap.Store(key, key)
```

Random Read

```go
v, ok := syncMap.Load(key)
if !ok {
}
v = v
```

Rang Read All

```go
syncMap.Range(func(k, v interface{}) bool {
    k = k
    v = v
    return true
})
```

### 2. sync.Mutex

Write

```go
mutex.Lock()
mutexMap[key] = key
mutex.Unlock()
```

Random Read

```go
mutex.Lock()
v, ok := mutexMap[key]
if !ok {
}
v = v
mutex.Unlock()
```

Rang Read All

```go
mutex.Lock()
for k, v := range mutexMap {
    k = k
    v = v
}
mutex.Unlock()
```

### 3. sync.RWMutex

Write

```go
rwMutex.Lock()
rwMutexMap[key] = key
rwMutex.Unlock()
```

Random Read

```go
rwMutex.RLock()
v, ok := rwMutexMap[key]
if !ok {
}
v = v
rwMutex.RUnlock()
```

Rang Read All

```go
rwMutex.RLock()
for k, v := range rwMutexMap {
    k = k
    v = v
}
rwMutex.RUnlock()
```

### 4. orcaman/concurrent-map

Write

```go
cMap.Set(key, key)
```

Random Read

```go
v, ok := cMap.Get(key)
if !ok {
}
v = v
```

Rang Read All

```go
for i := range cMap.IterBuffered() {
    i = i
}
```

## Test scenarios

Check all the test scenarios here [benchmark/scenarios.go](benchmark/scenarios.go)

### 10 concurrents 100k write

| Method                 | Test                                  | Time         |
| ---------------------- | ------------------------------------- | ------------ |
| orcaman/concurrent-map | write (10 concurrents) (100000 iters) | 62.288221ms  |
| RWMutex                | write (10 concurrents) (100000 iters) | 113.321432ms |
| Mutex                  | write (10 concurrents) (100000 iters) | 121.2995ms   |
| sync.Map               | write (10 concurrents) (100000 iters) | 203.216039ms |

### 10 concurrents 100k random read

| Method                 | Test                                        | Time        |
| ---------------------- | ------------------------------------------- | ----------- |
| orcaman/concurrent-map | random_read (10 concurrents) (100000 iters) | 15.110771ms |
| RWMutex                | random_read (10 concurrents) (100000 iters) | 30.6894ms   |
| sync.Map               | random_read (10 concurrents) (100000 iters) | 38.139147ms |
| Mutex                  | random_read (10 concurrents) (100000 iters) | 85.278314ms |

### 10 concurrents 100 range read all

| Method                 | Test                                    | Time         |
| ---------------------- | --------------------------------------- | ------------ |
| RWMutex                | range_read (10 concurrents) (100 iters) | 143.054227ms |
| Mutex                  | range_read (10 concurrents) (100 iters) | 905.993286ms |
| sync.Map               | range_read (10 concurrents) (100 iters) | 949.263514ms |
| orcaman/concurrent-map | range_read (10 concurrents) (100 iters) | 2.366558255s |

### 50 concurrents 20k write

| Method                 | Test                                 | Time         |
| ---------------------- | ------------------------------------ | ------------ |
| sync.Map               | write (50 concurrents) (20000 iters) | 22.123396ms  |
| orcaman/concurrent-map | write (50 concurrents) (20000 iters) | 30.682173ms  |
| Mutex                  | write (50 concurrents) (20000 iters) | 122.591997ms |
| RWMutex                | write (50 concurrents) (20000 iters) | 134.832297ms |

### 50 concurrents 20k random read

| Method                 | Test                                       | Time         |
| ---------------------- | ------------------------------------------ | ------------ |
| sync.Map               | random_read (50 concurrents) (20000 iters) | 7.055942ms   |
| orcaman/concurrent-map | random_read (50 concurrents) (20000 iters) | 12.597161ms  |
| RWMutex                | random_read (50 concurrents) (20000 iters) | 32.372341ms  |
| Mutex                  | random_read (50 concurrents) (20000 iters) | 118.862146ms |

### 50 concurrents 20 range read all

| Method                 | Test                                   | Time         |
| ---------------------- | -------------------------------------- | ------------ |
| RWMutex                | range_read (50 concurrents) (20 iters) | 104.010416ms |
| sync.Map               | range_read (50 concurrents) (20 iters) | 577.479455ms |
| Mutex                  | range_read (50 concurrents) (20 iters) | 863.322488ms |
| orcaman/concurrent-map | range_read (50 concurrents) (20 iters) | 1.715899163s |

### 100 concurrents 10k write

| Method                 | Test                                   | Time         |
| ---------------------- | -------------------------------------- | ------------ |
| sync.Map               | write (100 concurrents) (100000 iters) | 260.946224ms |
| orcaman/concurrent-map | write (100 concurrents) (100000 iters) | 349.440739ms |
| Mutex                  | write (100 concurrents) (100000 iters) | 1.630196524s |
| RWMutex                | write (100 concurrents) (100000 iters) | 1.680058588s |

### 100 concurrents 10k random read

| Method                 | Test                                         | Time         |
| ---------------------- | -------------------------------------------- | ------------ |
| sync.Map               | random_read (100 concurrents) (100000 iters) | 119.658652ms |
| orcaman/concurrent-map | random_read (100 concurrents) (100000 iters) | 126.84391ms  |
| RWMutex                | random_read (100 concurrents) (100000 iters) | 280.359023ms |
| Mutex                  | random_read (100 concurrents) (100000 iters) | 1.644604438s |

### 100 concurrents 10 range read all

| Method                 | Test                                    | Time         |
| ---------------------- | --------------------------------------- | ------------ |
| RWMutex                | range_read (100 concurrents) (10 iters) | 95.519575ms  |
| sync.Map               | range_read (100 concurrents) (10 iters) | 836.398188ms |
| Mutex                  | range_read (100 concurrents) (10 iters) | 972.755566ms |
| orcaman/concurrent-map | range_read (100 concurrents) (10 iters) | 2.795446594s |
