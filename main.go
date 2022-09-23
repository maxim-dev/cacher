package main

import (
	"cacher/cacher"
	"fmt"
	"time"
)

func main() {
	cache := cacher.NewCache(2 * time.Second)

	cache.Set("first", "one")
	cache.Set("second", "two")
	time.Sleep(1 * time.Second)
	f, okF := cache.Get("first")
	fmt.Println("first", f, okF)
	s, okS := cache.Get("second")
	fmt.Println("second1", s, okS)

	// При повторном запросе ключа 'second', он не вернется, так как ttl истекло
	time.Sleep(3 * time.Second)
	s, okS = cache.Get("second")
	fmt.Println("second2", s, okS)
}
