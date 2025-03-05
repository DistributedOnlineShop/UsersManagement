package util

import (
	"golang.org/x/crypto/argon2"
)

func Hash(password, salt string) ([]byte, error) {
	Config, err := LoadConfig("../")
	if err != nil {
		return []byte{}, err
	}

	hash := argon2.Key([]byte(password), []byte(salt), Config.TimeCost, Config.MemoryCost, Config.Parallelism, Config.KeyLength)
	return hash, nil
}

func VerifyHashPassword(password, salt string, hash []byte) bool {
	Config, err := LoadConfig("../")
	if err != nil {
		return false
	}

	inputPasswordHash := argon2.Key([]byte(password), []byte(salt), Config.TimeCost, Config.MemoryCost, Config.Parallelism, Config.KeyLength)

	return string(inputPasswordHash) == string(hash)
}
