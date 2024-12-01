package main

import (
	"crypto/sha1"
	"fmt"
	"hash"
)

func CreateHash(byteStr []byte) []byte {
	var hashVal hash.Hash
	hashVal = sha1.New()
	hashVal.Write(byteStr)

	var bytes []byte
	bytes = hashVal.Sum(nil)
	return bytes
}

func CreateHashMultiply(byteStr1 []byte, byteStr2 []byte) []byte {
	return xor(CreateHash(byteStr1), CreateHash(byteStr2))
}

func xor(byteStr1 []byte, byteStr2 []byte) []byte {
	var xorbytes []byte
	xorbytes = make([]byte, len(byteStr1))
	var i int
	for i = 0; i < len(byteStr1); i++ {
		xorbytes[i] = byteStr1[i] ^ byteStr2[i]
	}
	return xorbytes
}

func main() {
	var bytes []byte
	bytes = CreateHashMultiply([]byte("Check"), []byte("Hash"))
	fmt.Printf("%x\n", bytes)
}
