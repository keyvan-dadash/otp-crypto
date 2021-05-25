package main

import (
	"fmt"
	"otp-crypto/config"
	"otp-crypto/otp"
)

func main() {

	fmt.Print(otp.GenerateOTP(config.HMAC_SHA1, "key", 33, config.Lenght_6))
}
