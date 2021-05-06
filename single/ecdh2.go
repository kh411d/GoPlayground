package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/aead/ecdh"
)

type JsonResponse struct {
	Message string `json:"message"`
}

type JsonRequest struct {
	Message string `json:"message"`
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t JsonRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	alicePubX := r.Header.Get("Pub-X")
	alicePubY := r.Header.Get("Pub-Y")

	fmt.Printf("alicePubX %s\n", alicePubX)
	fmt.Printf("alicePubY %s\n", alicePubY)

	secret, pubx, puby := getKey(alicePubX, alicePubY)

	fmt.Printf("bobPubX %s\n", pubx)
	fmt.Printf("bobPubY %s\n", puby)

	//fmt.Printf("Encrypted Message Request: %s \n", t.Message)
	//fmt.Printf("Decrypted Message Request: %s \n", string(decrypt([]byte(t.Message), string(secret))))
	fmt.Printf("Secret Bob: %x \n", secret)

	data := JsonResponse{
		Message: "test", //string(encrypt([]byte("0818228467"), string(secret))),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Pub-X", pubx)
	w.Header().Set("Pub-Y", puby)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getKey(pubX, pubY string) ([]byte, string, string) {
	p256 := ecdh.Generic(elliptic.P256())

	privateBob, publicBob, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
	}

	// priv := new(big.Int)
	// priv.SetString("95415940569486884419632158163156959896840196456699243971802766620573781555324", 10)
	// privateBob = priv.Bytes()

	x := new(big.Int)
	x.SetString(pubX, 10)

	y := new(big.Int)
	y.SetString(pubY, 10)

	publicAlice := ecdh.Point{
		X: x,
		Y: y,
	}

	if err := p256.Check(publicAlice); err != nil {
		fmt.Printf("Alice's public key is not on the curve: %s\n", err)
	}
	secretBob := p256.ComputeSecret(privateBob, publicAlice)

	return secretBob, publicBob.(ecdh.Point).X.String(), publicBob.(ecdh.Point).Y.String()
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {

	http.HandleFunc("/", sendMessage)        // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
