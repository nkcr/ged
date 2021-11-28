// Copyright (c) 2021 Noémien Kocher

// Package alphabet defines character sets for encoding/decoding.
package alphabet

import (
	"fmt"
	"math"
)

var (
	Base58Bitcoin = MustCreate("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	Base64        = MustCreate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	Base64URL     = MustCreate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
	Base32        = MustCreate("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	Base32Hex     = MustCreate("0123456789ABCDEFGHIJKLMNOPQRSTUV")
	Hex           = MustCreate("0123456789abcdef")
	HexCapital    = MustCreate("0123456789ABCDEF")
	DNA           = MustCreate("ATCG")
)

// MustCreate returns a new initialized alphabet. The function panics if the
// alphabet's length is greater than 256 or lower than 2. The charset should not
// contain duplicate characters.
func MustCreate(charset string) Alphabet {
	if len(charset) > math.MaxUint8 {
		panic(fmt.Sprintf("alphabet's length too big: %d > %d",
			len(charset), math.MaxUint8))
	}

	if len(charset) <= 1 {
		panic(fmt.Sprintf("alphabet's length too small: %d <= 1", len(charset)))
	}

	decoding := make(map[rune]uint8)

	for i, c := range charset {
		decoding[c] = uint8(i)
	}

	return Alphabet{
		Charset:  charset,
		BaseTo:   uint(len(charset)),
		Decoding: decoding,
	}
}

type Alphabet struct {
	// Charset contains the list of runes and define their values, which
	// correspond to their indexes from the list. There shouldn't be duplicate,
	// even if allowed.
	Charset string

	// BaseTo is the encoding base target.
	BaseTo uint

	// Decoding maps a rune to its corresponding value.
	Decoding map[rune]uint8
}
