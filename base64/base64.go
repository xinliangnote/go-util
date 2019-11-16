package base64

import (
	b64 "encoding/base64"
)

func Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) (string, error) {
	ds, err := b64.StdEncoding.DecodeString(s)
	return string(ds), err
}
