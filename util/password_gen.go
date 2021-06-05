package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// PossibleSymbols we allow for when we generate passwords
const PossibleSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&/?<>*"

var maxlen = big.NewInt(int64(len(PossibleSymbols)))

func GeneratePassword(length int) (string, error) {
	output := make([]byte, 0, length) // treat slice as empty incase some symbols are > 1 byte

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maxlen)
		if err != nil {
			return "", fmt.Errorf("generating random symbol: %w", err)
		}

		output = append(output, PossibleSymbols[n.Int64()])
	}

	return string(output), nil
}
