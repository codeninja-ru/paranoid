package main

import (
	"crypto/rand"
	"fmt"
	//"math/rand"
)

// inspirited by https://golang.org/src/crypto/rand/util.go
// https://gist.github.com/joepie91/7105003c3b26e65efcea63f3db82dfba
// https://github.com/ai/nanoid
// https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/
// https://github.com/matoous/go-nanoid
func Goid64(length int) (string, error) {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~" // the length must be 64 bytes

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, _ := range bytes {
		bytes[i] = alphabet[bytes[i]&63]
	}

	return string(bytes[:]), nil

}

func Govnoid() (string, error) {
	const size = 21
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	var bitLen byte = 0
	var aLen byte = byte(len(alphabet))
	for b := len(alphabet) - 1; b > 0; bitLen++ {
		b = b >> 1
	}
	maxNumber := ^(^byte(0) << bitLen)
	diff := maxNumber - aLen + 1                                             // including zero
	steps := byte(float32(aLen)*(((1/float32(bitLen))*float32(diff))+1)) + 1 // plas 1 to not ceil
	//fmt.Printf("len = %d,bitLen = %d, max = %d, diff = %d, steps = %v\n", aLen, bitLen, maxNumber, diff, steps)

	bytes := make([]byte, steps)
	chars := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	j := byte(0)
	for i, _ := range bytes {
		bytes[i] &= 1<<bitLen - 1
		if bytes[i] < aLen-1 {
			chars[j] = alphabet[bytes[i]]
			j++
		}
		if j >= size {
			break
		}
	}

	return string(chars[:j]), nil

}

func main() {
	m := make(map[rune]int)

	for i := 0; i < 1000000; i++ {
		uuid, _ := Goid64(21)
		for _, val := range uuid {
			m[val] = m[val] + 1
		}
		fmt.Print(uuid)
	}
	for idx, val := range m {
		fmt.Printf("%#U - %d\n", idx, val)
	}
	//for {
	//	uuid, _ := Govnoid()
	//	fmt.Println(uuid)
	//	if len(uuid) < 21 {
	//		break
	//	}
	//}
}
