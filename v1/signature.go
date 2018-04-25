package v1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func calculateSignature(string2Sign string, secret string) string {
	hm := hmac.New(sha1.New, []byte(secret))
	hm.Write([]byte(string2Sign))
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}

func Signature(date string, data string, secret string) string {
	string2Sign := date + "\n" + data + "\n"
	return calculateSignature(string2Sign, secret)
}
