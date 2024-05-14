package util

import "math/rand"

func GenerateRandomNumber(upper int, lower int) int {
	return upper + rand.Intn(lower-upper+1)
}
