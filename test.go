package main

import "ratelimit/lib"

func main() {
	c := lib.NewLRUCache()
	c.Set("name","melt")
	c.Set("age","19")
	c.Set("sex","woman")

	c.Get("name")
	c.RemoveOldest()
	c.Print()
}