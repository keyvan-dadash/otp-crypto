package auth

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
)

const (
	QRSize = 256
)

type LabelInfo struct {
	Account string
	Issuer  string
}

func (li *LabelInfo) String() string {
	return fmt.Sprintf("%s:%s", li.Issuer, li.Account)
}

type KeyUri struct {
	Type       string
	Label      LabelInfo
	Parameters Params
}

type Params interface {
	GetString() string
}

func (ku *KeyUri) GetURL() (*url.URL, error) {
	rawUri := fmt.Sprintf("otpauth://%s/%s?%s", ku.Type, ku.Label.String(), ku.Parameters.GetString())
	uri, err := url.Parse(rawUri)
	if err != nil {
		logrus.Errorf("Cannot parse rawUri to URL structure. error: %v", err)
		return nil, err
	}
	return uri, nil
}

func (ku *KeyUri) String() string {
	uri, err := ku.GetURL()
	if err != nil {
		logrus.Errorf("Cannot convert url to string. error: %v", err)
		return ""
	}
	return uri.String()
}

func (ku *KeyUri) QRCodeString() string {
	uri, err := ku.GetURL()
	if err != nil {
		logrus.Errorf("Cannot convert url to QRCodeString. error: %v", err)
		return ""
	}

	qrCodeByte, err := qrcode.Encode(uri.String(), qrcode.Medium, QRSize)
	if err != nil {
		logrus.Errorf("Cannot convert url string QRCode byte. error: %v", err)
		return ""
	}

	return "data:image/png;base64," + base64.RawURLEncoding.EncodeToString(qrCodeByte)
}

func (ku *KeyUri) QRCodeImage(filename string) error {
	uri, err := ku.GetURL()
	if err != nil {
		logrus.Errorf("Cannot convert url to QRCodeString. error: %v", err)
		return err
	}

	err = qrcode.WriteFile(uri.String(), qrcode.Medium, 256, filename)
	if err != nil {
		logrus.Errorf("Cannot create QRCode image from uri string. error: %v", err)
		return err
	}

	return nil
}
