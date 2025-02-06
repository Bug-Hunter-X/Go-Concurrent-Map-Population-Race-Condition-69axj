# Go Concurrent Map Population Race Condition

This repository demonstrates a common race condition that can occur when concurrently populating a map in Go. The program uses goroutines to populate a map, but without proper synchronization, the final map size may be less than expected.

## Bug Description
The `bug.go` file contains a program that uses goroutines to populate a map.  Each goroutine adds a key-value pair to the map.  However, without proper synchronization mechanisms, there's a race condition when multiple goroutines try to access and modify the map concurrently.  This can lead to data loss or inconsistency.

## Solution
The `bugSolution.go` file provides a solution using a `sync.Mutex` to protect the map from concurrent access.  The mutex ensures that only one goroutine can access the map at a time, preventing data races and ensuring the correct map size.
