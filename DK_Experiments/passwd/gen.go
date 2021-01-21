package main

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
)

const (
	PW_SALT_BYTES = 8
	PW_HASH_BYTES = 32
	PASS_WORD     = "hello scrypt"
)

func main() {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("salt=%v\n", salt)
		fmt.Printf("salt=%x\n", salt)
	}

	hash, err := scrypt.Key([]byte(PASS_WORD), salt, 1<<15, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("encrypt=%x\n", hash)
}
