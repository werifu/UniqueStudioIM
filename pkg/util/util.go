package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
)

func res2token(body []byte) string {
	bodyString := string(body)
	return strings.Split(strings.Split(bodyString, "&")[0], "=")[1]
}

func StringSha256(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}