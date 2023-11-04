package main

import (
	"fmt"
	"go-mod-ninga/cash"
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
