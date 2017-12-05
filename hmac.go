package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "fmt"
)

func main() {
    stringToSign := []byte("merchantkey" + "random")
    h := hmac.New(sha256.New, []byte("apikey"))

    h.Write(stringToSign)

    w := h.Sum(nil)

    fmt.Printf("%x\n", w)
}
