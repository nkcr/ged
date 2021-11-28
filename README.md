# GED - Global-purpose Encoding / Decoding library

This library lets you use common encoding/decoding schemes and allows you to use
custom ones.

Use standard hex encoding/decoding:

```go
encoded := ged.EncodeHex([]byte("Hello World"))
// => 48656c6c6f20576f726c64
decoded, err := ged.DecodeHex(encoded)
// => Hello World
```

Uses standard base 58 encoding/decoding:

```go
encoded := ged.Encode58([]byte("Hello World"))
// => JxF12TrwUP45BMd
decoded, err := ged.Decode58(encoded)
// => Hello World
```

Use a custom encoding:

```go
alphabet := alphabet.MustCreate("ABCD1234")

encoded := ged.EncodeString([]byte("Hello World"), alphabet)
// => 11B12DDA33B24BAA2D224D1133B11
decoded, err := ged.DecodeString(encoded, alphabet)
// => Hello World
```