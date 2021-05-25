package otp

import (
	"crypto/hmac"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"otp-crypto/config"
	"strings"

	"github.com/sirupsen/logrus"
)

func GenerateOTP(hmacAlgo config.HmacAlgo, secretKey string, counter int, lenght config.Lenght) string {

	secretKey = strings.ToUpper(secretKey)

	secretKey = strings.TrimRight(secretKey, string(base32.StdPadding))

	generatedHMAC := hmac.New(hmacAlgo.Hash, []byte(secretKey))

	counterByte := make([]byte, 8)

	binary.BigEndian.PutUint64(counterByte, uint64(counter))

	n, err := generatedHMAC.Write(counterByte)
	if err != nil {
		logrus.Error("cannot write generated hmac on byte array. error: %v", err)
		return ""
	}

	hmacByte := generatedHMAC.Sum([]byte{})

	fmt.Println(hmacByte)

	lastByte := hmacByte[n-1]

	last4bit := lastByte & 15

	fmt.Println(last4bit)

	extractedInt := extract31(hmacByte, int(last4bit))

	return string(lenght.Truncate(extractedInt))
}

func extract31(hmacByteArray []byte, index int) int {
	bitsArray := bytes2bits(hmacByteArray)
	truncatedBits := bitsArray[index*8 : index*8+(4*8)-1]

	var extractedInt int = 0

	for index, v := range truncatedBits {
		extractedInt += v << (len(truncatedBits) - index)
	}
	fmt.Println(extractedInt)

	return extractedInt
}

func bytes2bits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	return dst
}
