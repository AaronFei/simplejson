# BuffPool

BuffPool is an efficient, thread-safe buffer pool management library for Go. It uses generics to support data of any type and provides a simple API for managing and reusing fixed-size buffers.

## Features

- Supports data of any type (using Go generics)
- Thread-safe buffer management
- Efficient buffer reuse mechanism
- Clean and simple API design
- Supports non-blocking operations
- Built-in error handling and logging

## Installation

Install BuffPool using Go modules:

```bash
go get github.com/yourusername/buffpool

Quick Start
Here's a simple example of how to use BuffPool:
go
package main

import (
    "fmt"
    "sync"
    "github.com/yourusername/buffpool"
)

func producer(pool *buffpool.Pool[int], bufChan chan<- *buffpool.Buffer[int], count int) {
    for i := 0; i < count; i++ {
        buf, ok := pool.Acquire()
        if !ok {
            fmt.Println("Failed to acquire buffer")
            continue
        }

        // Fill data and set length
        data := buf.GetFullData()
        for j := 0; j < 10; j++ {
            data[j] = i * 10 + j
        }
        buf.SetLength(10)

        // Send buffer through channel
        bufChan <- buf
    }
    close(bufChan)
}

func consumer(id int, bufChan <-chan *buffpool.Buffer[int], wg *sync.WaitGroup) {
    defer wg.Done()
    for buf := range bufChan {
        // Use GetValidData to get the valid slice
        validData := buf.GetValidData()
        fmt.Printf("Consumer %d received: %v\n", id, validData)

        // Release buffer after use
        buf.Release()
    }
}

func main() {
    // Create a buffer pool for int type
    pool := buffpool.NewPool[int]()
    err := pool.Init(5, 100)
    if err != nil {
        panic(err)
    }

    bufChan := make(chan *buffpool.Buffer[int], 5)

    // Start producer
    go producer(pool, bufChan, 10)

    // Start consumers
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go consumer(i, bufChan, &wg)
    }

    // Wait for all consumers to finish
    wg.Wait()

    fmt.Println("Available buffers:", pool.Available())
}

API Overview
Pool
NewPool[T any](): Create a new Pool
Init(bufCount, bufSize int) error: Initialize the Pool
Acquire() (*Buffer[T], bool): Acquire an available buffer
BufferChan() <-chan *Buffer[T]: Return a read-only channel for acquiring buffers
Available() int: Return the number of currently available buffers
Reset(): Reset the Pool, clear and refill the pool
Buffer
SetLength(len int): Set the effective length of the buffer
GetLength() int: Get the effective length of the buffer
GetValidData() []T: Get the valid data in the buffer
GetFullData() []T: Get all data in the buffer
Release(): Release the buffer back to the pool
Performance Considerations
BuffPool is designed to reduce memory allocation and garbage collection pressure, making it particularly suitable for scenarios that require frequent creation and destruction of temporary buffers. By reusing pre-allocated buffers, it can significantly improve application performance, especially in high-concurrency environments.
Thread Safety
BuffPool uses synchronization primitives to ensure that all operations are thread-safe, making it suitable for use in multi-goroutine concurrent environments.
Contributing
Contributions are welcome! Please feel free to submit issues and pull requests to help improve this project.
