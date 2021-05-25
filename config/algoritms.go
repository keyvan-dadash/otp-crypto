package config

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type HmacAlgo int

const (
	HMAC_SHA1 HmacAlgo = iota + 1
	HMAC_SHA256
	HMAC_SHA512
)

func (algo HmacAlgo) Hash() hash.Hash {

	switch algo {
	case HMAC_SHA1:
		return sha1.New()
	case HMAC_SHA256:
		return sha256.New()
	case HMAC_SHA512:
		return sha512.New()
	default:
		panic("this HMAC algoritm not supported yet")
	}
}

func (algo HmacAlgo) GetHashNameString() hash.Hash {

	switch algo {
	case HMAC_SHA1:
		return sha1.New()
	case HMAC_SHA256:
		return sha256.New()
	case HMAC_SHA512:
		return sha512.New()
	default:
		panic("this HMAC algoritm not supported yet")
	}
}
