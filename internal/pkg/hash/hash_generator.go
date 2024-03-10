package hash

import (
	"crypto/md5"
	b64 "encoding/base64"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GenerateHash(r *http.Request) string {
	ip := getUserIP(r)

	md5HashBytes := generateMd5Hash(ip)
	base64Str := b64.StdEncoding.EncodeToString(md5HashBytes)

	return base64Str[0:8]
}

func generateMd5Hash(ip string) []byte {
	md5Hash := md5.New()
	_, _ = io.WriteString(md5Hash, ip)
	_, _ = io.WriteString(md5Hash, strconv.FormatInt(time.Now().UnixNano(), 10))
	return md5Hash.Sum(nil)
}

func getUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
