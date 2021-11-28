package ged

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/nkcr/ged/alphabet"

	"github.com/stretchr/testify/require"
)

func Test_Encode_Decode_Hex(t *testing.T) {
	// https://datatracker.ietf.org/doc/html/rfc4648#section-10

	table := [][2]string{
		{"", ""},
		{"f", "66"},
		{"fo", "666f"},
		{"foo", "666f6f"},
		{"foob", "666f6f62"},
		{"fooba", "666f6f6261"},
		{"foobar", "666f6f626172"},
	}

	// Encode
	for _, e := range table {
		require.Equal(t, e[1], EncodeHex([]byte(e[0])))
	}

	// Decode
	for _, e := range table {
		decoded, err := DecodeHex(e[1])
		require.NoError(t, err)

		require.Equal(t, e[0], string(decoded))
	}
}

func TestBla(t *testing.T) {
	alphabet := alphabet.MustCreate("ABCD1234")
	encoded := EncodeString([]byte("Hello World"), alphabet)
	fmt.Println(encoded)
	decoded, _ := DecodeString(encoded, alphabet)
	fmt.Println(string(decoded))
}

func Test_Encode_Decode_Base58(t *testing.T) {
	// https://tools.ietf.org/id/draft-msporny-base58-01.html#rfc.section.5

	table := [][2]string{
		{"", ""},
		{"Hello World!", "2NEpo7TZRRrLZSi2U"},
		{"The quick brown fox jumps over the lazy dog.", "USm3fpXnKG5EUBx2ndxBDMPVciP5hGey2Jh4NDv6gmeo1LkMeiKrLJUUBk6Z"},
		{"\x00\x00\x00(\u007f\xb4\xcd", "111233QC4"},
	}

	// Decode
	for _, e := range table {
		decoded, err := Decode58(e[1])
		require.NoError(t, err)

		require.Equal(t, e[0], string(decoded))
	}

	// Encode
	for _, e := range table {
		require.Equal(t, e[1], Encode58([]byte(e[0])))
	}
}

func Test_Transform(t *testing.T) {
	rand.Seed(1)
	iterations := 500

	for i := 0; i < iterations; i++ {
		buf := make([]byte, i)

		rand.Read(buf)

		encoded, err := Transform(buf, 256, 58)
		require.NoError(t, err)

		decoded, err := Transform(encoded, 58, 256)
		require.NoError(t, err)

		require.Equal(t, buf, decoded, fmt.Sprintf("%x - %x - %x", buf, encoded, decoded))
		fmt.Print(".")
	}
	fmt.Println("done")
}
