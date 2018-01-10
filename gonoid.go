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
func Goid64() (string, error) {
	const size = 21
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~" // the length must be 64 bytes
	bitLen := byte(6)                                                                   //byte(math.Floor(math.Log2(float64(len(alphabet)-1)) + 1))

	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, _ := range bytes {
		bytes[i] &= 1<<bitLen - 1
		bytes[i] = alphabet[bytes[i]]
	}

	return string(bytes[:]), nil

}

func Govnoid() (string, error) {
	const size = 21
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	var bitLen byte = 0
	var length byte = byte(len(alphabet))
	for b := len(alphabet) - 1; b > 0; bitLen++ {
		b = b >> 1
	}
	maxNumber := ^(^byte(0) << bitLen)
	diff := maxNumber - length + 1 // including zero
	//bitLen := byte(6) //byte(math.Floor(math.Log2(float64(len(alphabet)-1)) + 1))
	steps := byte(float32(length)*(((1/float32(bitLen))*float32(diff))+1)) + 1 // plas 1 to not ceil
	//fmt.Printf("len = %d,bitLen = %d, max = %d, diff = %d, steps = %v\n", length, bitLen, maxNumber, diff, steps)

	bytes := make([]byte, steps)
	chars := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	j := byte(0)
	for i, _ := range bytes {
		bytes[i] &= 1<<bitLen - 1
		//fmt.Printf("v = %v, i = %d, j = %d\n", bytes[i], i, j)
		if bytes[i] < length-1 {
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
	for i := 0; i < 100000; i++ {
		uuid, _ := Govnoid()
		fmt.Println(uuid)
	}
	//for {
	//	uuid, _ := Govnoid()
	//	fmt.Println(uuid)
	//	if len(uuid) < 21 {
	//		break
	//	}
	//}
}
