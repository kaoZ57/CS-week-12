package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}
