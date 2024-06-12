package util

import "math/rand"

func RandomType() string {
	types := []string{"Debit", "Credit"}
	n := len(types)
	return types[rand.Intn(n)]
}
