package utils

import (
    "math/rand"
    "time"
    "strconv"
    "crypto/sha512"
	"encoding/hex"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CreateRandomKey(size int) string {
    rand.Seed(time.Now().UnixNano())
    key := make([]byte, size)
    for i := range key {
        key[i] = charset[rand.Intn(len(charset))]
    }
    return string(key)
}

const charsetnum = "0123456789"

func CreateRandomKeyNumber(size int) int {
    rand.Seed(time.Now().UnixNano())
    key := make([]byte, size)
    for i := range key {
        key[i] = charsetnum[rand.Intn(len(charsetnum))]
    }
    keyInt, err := strconv.Atoi(string(key))
    if err != nil {
        return 0
    }
    return keyInt
}

func CalculateSHA512Hash(data ...string) (string, error) {
	dataToHash := ""
	for _, s := range data {
		dataToHash += s
	}

	hasher := sha512.New()
	hasher.Write([]byte(dataToHash))
	hashedData := hasher.Sum(nil)

	hashedDataHex := hex.EncodeToString(hashedData)

	return hashedDataHex, nil
}

