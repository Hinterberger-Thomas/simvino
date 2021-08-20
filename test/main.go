package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func main() {

	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	hash, err := generateFromPassword("password123", p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)

}

func generateFromPassword(password string, p *params) (hash []byte, err error) {
	// Generate a cryptographically secure random salt.
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte("some password"), salt, 1, 64*1024, 4, 32)

	return hash, nil
}
func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
