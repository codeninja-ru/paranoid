package main

import (
	"crypto/rand"
	"math"
	"strings"
	"testing"
)

// inspirited by https://golang.org/src/crypto/rand/util.go
// https://gist.github.com/joepie91/7105003c3b26e65efcea63f3db82dfba
// https://github.com/ai/nanoid
// https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
// https://github.com/matoous/go-nanoid
func Gonoid() (string, error) {
	const size = 22
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~"
	bitLen := byte(math.Floor(math.Log2(float64(len(alphabet)-1)) + 1))

	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, _ := range bytes {
		bytes[i] &= byte(1<<bitLen - 1)
		bytes[i] = alphabet[bytes[i]]
	}

	return string(bytes[:]), nil

}

type DefaultsType struct {
	Alphabet string
	Size     int
	MaskSize int
}

var defaults = DefaultsType{
	Alphabet: "_~0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", // len=64
	Size:     22,
	MaskSize: 5,
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~"

var urlLength = len(alphabet)

// Test that nanoid generates URL friendly IDs
// it ('generates URL-friendly IDs')
func TestGeneratesURLFriendlyIDs(t *testing.T) {
	for i := 0; i < 10; i++ {
		id, err := Gonoid()
		if err != nil {
			t.Errorf("Nanoid error: %v", err)
		}
		if len(id) != defaults.Size {
			t.Errorf(
				"TestGeneratesURLFriendlyIDs error: length of id %v should be %v, got %v",
				id,
				defaults.Size,
				id,
			)
		}

		runeID := []rune(id)

		for j := 0; j < len(runeID); j++ {
			res := strings.Contains(defaults.Alphabet, string(runeID[j]))
			if !res {
				t.Errorf(
					"GeneratesURLFriendlyIds error: char %v should be contained in %v",
					string(runeID[j]),
					defaults.Alphabet,
				)
			}
		}
	}
}

// Test that nanoid has no collisions
// it ('has no collisions')
func TestHasNoCollisions(t *testing.T) {
	COUNT := 100 * 1000
	used := make(map[string]bool)
	for i := 0; i < COUNT; i++ {
		id, err := Gonoid()
		if err != nil {
			t.Errorf("Nanoid error: %v", err)
		}
		if used[id] {
			t.Errorf("Collision error! Id %v found for test arr %v", id, used)
		}
		used[id] = true
	}
}

// Test that Nanoid has flat distribution
// it ('has flat distribution')
func TestFlatDistribution(t *testing.T) {
	COUNT := 100 * 1000
	instance, err := Gonoid()
	if err != nil {
		t.Errorf("Nanoid error: %v", err)
	}
	LENGTH := len(instance)

	chars := make(map[byte]int)

	for i := 0; i < COUNT; i++ {
		id, _ := Gonoid()
		for j := 0; j < LENGTH; j++ {
			// https://github.com/ai/nanoid/blob/d6ad3412147fa4c2b0d404841ade245a00c2009f/test/index.test.js#L33
			// if (!chars[char]) chars[char] = 0 is useless since it
			// is initialized by default to 0 from Golang
			chars[id[j]]++
		}
	}

	for char, k := range chars {
		distribution := float64(k) * float64(urlLength) / float64(COUNT*LENGTH)
		if !toBeCloseTo(distribution, 1, 1) {
			t.Errorf("Distribution error! Distribution %v found for char %v", distribution, char)
		}
	}
}

// utility that replicates jest.toBeCloseTo
func toBeCloseTo(value, actual, expected float64) bool {
	precision := 2
	// https://github.com/facebook/jest/blob/a397abaf9f08e691f8739899819fc4da41c1e476/packages/expect/src/matchers.js#L83
	pass := math.Abs(expected-actual) < math.Pow10(-precision)/2
	return pass
}

// Benchmark nanoid generator
func BenchmarkNanoid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Gonoid()
	}
}
