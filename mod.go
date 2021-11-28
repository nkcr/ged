// Copyright 2021 Noémien Kocher under the MIT license.

// Package ged implements a global-purpose encoding/decoding library.
package ged

import (
	"ged/alphabet"
	"math"
	"strings"

	"golang.org/x/xerrors"
)

// EncodeHex encodes data to hexadecimal. Uses the lower case letter form.
func EncodeHex(data []byte) string {
	return EncodeString(data, alphabet.Hex)
}

// DecodeHex decodes encoded hex string.
func DecodeHex(encoded string) ([]byte, error) {
	// lower 'encoded' to ensure compatibility with other encoding that use
	// capital letters.
	return DecodeString(strings.ToLower(encoded), alphabet.Hex)
}

// EncodeBase58 encodes data to a base58 representation using the Bitcoin
// alphabet.
func Encode58(data []byte) string {
	return EncodeString(data, alphabet.Base58Bitcoin)
}

// DecodeBase58 decodes string from a base58 representation using the Bitcoin
// alphabet.
func Decode58(encoded string) ([]byte, error) {
	return DecodeString(encoded, alphabet.Base58Bitcoin)
}

// EncodeString encodes data using a provided alphabet.
func EncodeString(data []byte, alphabet alphabet.Alphabet) string {
	encoded := Encode(data, alphabet.BaseTo)

	out := new(strings.Builder)

	for _, e := range encoded {
		out.WriteByte(alphabet.Charset[e])
	}

	return out.String()
}

// DecodeString decodes string using the provided alphabet.
func DecodeString(encoded string, alphabet alphabet.Alphabet) ([]byte, error) {
	buf := make([]byte, len(encoded))

	for i, c := range encoded {
		val, ok := alphabet.Decoding[c]
		if !ok {
			return nil, xerrors.Errorf("invalid character: '%s'", c)
		}

		buf[i] = val
	}

	return Decode(buf, alphabet.BaseTo)
}

// Encode encodes data to a given base. baseTo SHOULD NOT be greater than 256.
func Encode(data []byte, baseTo uint) []byte {
	// there can't be an error because baseFrom is set to 256
	encoded, _ := Transform(data, 256, baseTo)
	return encoded
}

// Decode decodes data from a given base. baseTo SHOULD NOT be greater than 256.
func Decode(data []byte, baseFrom uint) ([]byte, error) {
	return Transform(data, baseFrom, 256)
}

// Transform is a low-level generic function to transform bytes from a given
// base to a target base. Returns an error if an element in data is greater or
// equal than the given base: if the baseFrom is 128, then all elements from
// data must be lower than 128. baseFrom and baseTo SHOULD NOT be greater than
// 256 and MUST NOT be lower than 2.
func Transform(data []byte, baseFrom, baseTo uint) ([]byte, error) {
	if baseFrom <= 1 {
		return nil, xerrors.Errorf("invalid baseFrom: %d <= 1", baseFrom)
	}

	if baseTo <= 1 {
		return nil, xerrors.Errorf("invalid baseTo: %d <= 1", baseTo)
	}

	prefixZeros := 0

	for _, e := range data {
		if e != 0 {
			break
		}

		prefixZeros++
	}

	factor := math.Log(float64(baseFrom)) / math.Log(float64(baseTo))
	cap := int(float64(len(data)-prefixZeros)*factor + 1)

	result := make([]byte, cap)

	reverseEnd := len(result) - 1
	var reverseIndex int

	for _, e := range data[prefixZeros:] {
		carry := uint(e)

		if carry >= baseFrom {
			return nil, xerrors.Errorf("invalid data: %d >= %d", e, baseFrom)
		}

		reverseIndex = len(result) - 1
		for ; reverseIndex > reverseEnd || carry != 0; reverseIndex-- {
			// we populate in reverse order
			carry = carry + uint(result[reverseIndex])*baseFrom

			result[reverseIndex] = byte(carry % baseTo)
			carry = carry / baseTo
		}
		// keep track of the last encoded index
		reverseEnd = reverseIndex
	}

	return append(make([]byte, prefixZeros), result...)[reverseEnd+1:], nil
}
