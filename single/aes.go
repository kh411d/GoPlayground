package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "fmt"
    "io"
)

func Pad(buf []byte, size int) ([]byte, error) {
    bufLen := len(buf)
    padLen := size - bufLen%size
    padded := make([]byte, bufLen+padLen)
    copy(padded, buf)
    for i := 0; i < padLen; i++ {
        padded[bufLen+i] = byte(padLen)
    }
    return padded, nil
}

func Unpad(padded []byte, size int) ([]byte, error) {
    if len(padded)%size != 0 {
        return nil, errors.New("pkcs7: Padded value wasn't in correct size.")
    }

    bufLen := len(padded) - int(padded[len(padded)-1])
    buf := make([]byte, bufLen)
    copy(buf, padded[:bufLen])
    return buf, nil
}

func main() {
    var ciphertext, plaintext []byte
    var err error

    // The key length can be 32, 24, 16  bytes (OR in bits: 128, 192 or 256)
    key := []byte("32423423n432423423n432423423n412")

    plaintext = []byte("P|1fd94d876c3c8c55016c747cc647000a|14840627")
    fmt.Printf("Orig  plaintext : %#v\n", []byte(plaintext))
    fmt.Printf("Orig Len plaintext : %d\n", len(plaintext))

    plaintext, err = Pad(plaintext, aes.BlockSize)
    if err != nil {
        fmt.Errorf(`plainText: "%s" has error`, plaintext)
    }
    fmt.Printf("Padded plaintext: %#v\n", []byte(plaintext))

    if ciphertext, err = encryptCBC(key, plaintext); err != nil {
        panic(err)
    }

    //ciphertext, _ = base64.StdEncoding.DecodeString("fWJ3Z97ffg3vjkQYqy2gORFh9zXfvm7gNvmSTCCIswEx6X4USj3FejvdOPi9fAvf")

    cleartext := base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:])
    fmt.Printf("CBC: %s\n", cleartext)

    if plaintext, err = decryptCBC(key, ciphertext); err != nil {
        panic(err)
    }
    fmt.Printf("Clear from CBC: %s\n", plaintext)

}

// CBC
func encryptCBC(key, plaintext []byte) (ciphertext []byte, err error) {
    if len(plaintext)%aes.BlockSize != 0 {
        panic("plaintext is not a multiple of the block size")
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Lengh plaintext: %d\n", len(plaintext))

    ciphertext = make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]

    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    // x, _ := hex.DecodeString("31353636383038383337383539000000")
    // re := strings.NewReader(string(x))
    // if _, err := io.ReadFull(re, iv); err != nil {
    //     panic(err)
    // }

    fmt.Printf("CBC Key: %s\n", hex.EncodeToString(key))
    fmt.Printf("CBC IV: %s\n", hex.EncodeToString(iv))

    cbc := cipher.NewCBCEncrypter(block, iv)
    cbc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

    return
}
func decryptCBC(key, ciphertext []byte) (plaintext []byte, err error) {
    var block cipher.Block

    if block, err = aes.NewCipher(key); err != nil {
        return
    }

    if len(ciphertext) < aes.BlockSize {
        fmt.Printf("ciphertext too short")
        return
    }

    iv := ciphertext[:aes.BlockSize]

    ciphertext = ciphertext[aes.BlockSize:]

    fmt.Printf("decrypt chipertext : %s\n", hex.EncodeToString(ciphertext))
    fmt.Printf("decrypt iv: %s\n", hex.EncodeToString(iv))

    cbc := cipher.NewCBCDecrypter(block, iv)
    cbc.CryptBlocks(ciphertext, ciphertext)

    plaintext = ciphertext

    return
}
