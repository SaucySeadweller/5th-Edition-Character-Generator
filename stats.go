package main

import (
  "time"
  "math/rand"
  "fmt"
)
func statGenerator() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(17) + 1)
}
