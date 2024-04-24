package hash

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"io"
	"regexp"
	"strconv"
	"time"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func GenerateHash(ip string) string {
	md5HashBytes := generateMd5Hash(ip)
	base64Str := b64.StdEncoding.EncodeToString(md5HashBytes)
	alphanumericStr := nonAlphanumericRegex.ReplaceAllString(base64Str, "")
	return alphanumericStr[0:8]
}

func generateMd5Hash(ip string) []byte {
	sha256Hash := sha256.New()
	_, _ = io.WriteString(sha256Hash, ip)
	_, _ = io.WriteString(sha256Hash, strconv.FormatInt(time.Now().UnixNano(), 10))
	return sha256Hash.Sum(nil)
}
