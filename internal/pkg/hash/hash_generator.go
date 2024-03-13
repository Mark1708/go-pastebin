package hash

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GenerateHash(ip string) string {
	md5HashBytes := generateMd5Hash(ip)
	base64Str := b64.StdEncoding.EncodeToString(md5HashBytes)

	return base64Str[0:8]
}

func generateMd5Hash(ip string) []byte {
	sha256Hash := sha256.New()
	_, _ = io.WriteString(sha256Hash, ip)
	_, _ = io.WriteString(sha256Hash, strconv.FormatInt(time.Now().UnixNano(), 10))
	return sha256Hash.Sum(nil)
}

func GetUserIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}
