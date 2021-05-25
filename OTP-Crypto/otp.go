package otp

import (
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"otp-crypto/config"
	"strings"

	"github.com/sirupsen/logrus"
)

func generateOTP(hmacAlgo config.HmacAlgo, secretKey string, counter int, lenght config.Lenght) string {

	secretKey = strings.ToUpper(secretKey)

	secretKey = strings.TrimRight(secretKey, string(base32.StdPadding))

	generatedHMAC := hmac.New(hmacAlgo.Hash, []byte(secretKey))

	var hmacByte []byte

	n, err := generatedHMAC.Write(hmacByte)
	if err != nil {
		logrus.Error("cannot write generated hmac on byte array. error: %v", err)
		return ""
	}

	last4bit := hmacByte[n-4 : n]

	rawhotp := extract31(hmacByte, last4bit)

	return string(lenght.Truncate(int(binary.BigEndian.Uint32(rawhotp))))
}

func extract31(hmacByteArray []byte, last4bit []byte) []byte {
	index := binary.BigEndian.Uint16(last4bit)

	return hmacByteArray[index*8 : index*8+(4*8)-1]
}
