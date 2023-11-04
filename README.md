# in-memory-cache

# Go Cache

This package represents a basic implementation of a cache in the Go programming language. The cache provides the ability to store and access data by key.

## Installation

```bash
go get -u github.com/Vhoode/go-cache

package main

import (
    "fmt"
    "github.com/Vhoode/go-cache/cache"
)

func main() {
    cache := cache.New()

    cache.Set("userId", 42)
    cache.Set("userId2", 42)
    cache.Set("userId3", 42)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId = cache.Get("userId")
    fmt.Println(userId)

    fmt.Println(cache.GetMap())
}

API Description
Methods
New() *Cache
Initializes a new Cache object and returns it.

Get(key string) interface{}
Returns data associated with the given key from the cache.

GetMap() interface{}
Returns all the data stored in the cache.

Set(key string, value interface{})
Stores the value in the cache under the specified key.

Delete(key string)
Deletes the data from the cache based on the key.

Notes
The package provides basic functionalities to manage the cache.
Import github.com/your-profile/go-cache/cache to use the package.
Contributions
Any questions, suggestions, or improvements are welcomed! Please create issues or pull requests in the repository.
