package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {

	random := time.Now().Format("20060102150405")
	stringToSign := []byte("hypermart" + random)
	h := hmac.New(sha256.New, []byte("084b13ecac81e1a8caf1775ad02bd5fa40e7219c8956dba11429a497a0e4cd89"))
	h.Write(stringToSign)
	Hmac := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(random)
	fmt.Println(Hmac)
}
