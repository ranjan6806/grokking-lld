package encoder

import (
	"crypto/sha256"
	"encoding/base64"
)

type Encoder interface {
	Encode(url string) string
}

type Base66Encoder struct {
	charset string
}

func NewBase64Encoder(charset string) Encoder {
	return &Base66Encoder{charset}
}

func (e *Base66Encoder) Encode(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:])[:8]
}
