package main

import "local/uber-test/cache"

func main() {
	cache := cache.New()

	cache.Add("test", 25)
}
