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

## Test scenarios

Check all the test scenarios here [benchmark/scenarios.go](benchmark/scenarios.go)

### 10 concurrents 100k write

| Method   | Test                                  | Time         |
| -------- | ------------------------------------- | ------------ |
| Mutex    | write (10 concurrents) (100000 iters) | 84.885211ms  |
| RWMutex  | write (10 concurrents) (100000 iters) | 90.861559ms  |
| sync.Map | write (10 concurrents) (100000 iters) | 192.916029ms |

### 10 concurrents 100k random read

| Method   | Test                                        | Time        |
| -------- | ------------------------------------------- | ----------- |
| sync.Map | random_read (10 concurrents) (100000 iters) | 25.910149ms |
| RWMutex  | random_read (10 concurrents) (100000 iters) | 28.786763ms |
| Mutex    | random_read (10 concurrents) (100000 iters) | 68.581472ms |

### 10 concurrents 100 range read all

| Method   | Test                                    | Time         |
| -------- | --------------------------------------- | ------------ |
| RWMutex  | range_read (10 concurrents) (100 iters) | 128.789501ms |
| sync.Map | range_read (10 concurrents) (100 iters) | 656.390123ms |
| Mutex    | range_read (10 concurrents) (100 iters) | 851.01652ms  |

### 100 concurrents 10k write

| Method   | Test                                   | Time         |
| -------- | -------------------------------------- | ------------ |
| sync.Map | write (100 concurrents) (100000 iters) | 208.484083ms |
| Mutex    | write (100 concurrents) (100000 iters) | 1.303995562s |
| RWMutex  | write (100 concurrents) (100000 iters) | 1.451269116s |

### 100 concurrents 10k random read

| Method   | Test                                         | Time         |
| -------- | -------------------------------------------- | ------------ |
| sync.Map | random_read (100 concurrents) (100000 iters) | 77.787739ms  |
| RWMutex  | random_read (100 concurrents) (100000 iters) | 267.869797ms |
| Mutex    | random_read (100 concurrents) (100000 iters) | 1.150969105s |

### 100 concurrents 10 range read all

| Method   | Test                                    | Time         |
| -------- | --------------------------------------- | ------------ |
| RWMutex  | range_read (100 concurrents) (10 iters) | 96.962034ms  |
| sync.Map | range_read (100 concurrents) (10 iters) | 526.186879ms |
| Mutex    | range_read (100 concurrents) (10 iters) | 842.53781ms  |
