package cusrand

import (
	"strconv"

	"golang.org/x/exp/rand"
)

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomIntAsString(min, max int) string {
	return strconv.Itoa(RandomInt(min, max))
}
