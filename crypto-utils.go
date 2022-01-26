package simpleutils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
)

//GenerateRandomBytes returns a byte array of cryptographic (true)
//random numbers. The byte array length is specified by count.
func GenerateRandomBytes(count int) ([]byte, error) {
	bytes := make([]byte, count)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

//GenerateRandomString returns a base-64 encoding of a cryptographic
//(true) random number byte array. The byte array length is specified
//by count.
func GenerateRandomString(count int) (string, error) {
	bytes, err := GenerateRandomBytes(count)
	return base64.URLEncoding.EncodeToString(bytes), err
}

// HashedText is a typed alias for string
type HashedText string

//HashText returns a sha1 hash of text in base64 encoding.
func HashText(text string) HashedText {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := HashedText(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
	return sha
}
