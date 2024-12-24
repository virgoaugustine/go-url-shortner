package main

import "math/rand"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// generateShortURL creates a random string of 6 characters to be assigned as the short form of a long URL
func generateShortURL() string {
	res := make([]byte, 6)

	for i := range res {
		res[i] = charset[rand.Intn(len(charset))]
	}

	return string(res)
}
