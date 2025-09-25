package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP generates a secure 6-digit numeric OTP as a string.
func GenerateOTP() (string, error) {
	max := big.NewInt(1000000) // 6 digits: 000000 - 999999
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
