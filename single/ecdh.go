package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/aead/ecdh"
)

func main() {
	p256 := ecdh.Generic(elliptic.P256())

	privateAlice, publicAlice, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Alice's private/public key pair: %s\n", err)
	}

	/*privateBob, publicBob, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
	}

	priv := new(big.Int)
	priv.SetString("959002911870247915586476816117910306058956722294361644472722635894727245699", 10)
	privateBob = priv.Bytes()
	*/
	x := new(big.Int)
	x.SetString("52954863476421554369917303459086111141944620376294749798322602614812281075473", 10)
	y := new(big.Int)
	y.SetString("78752030502886145285636836736474826406308298965952922287520199873579721619311", 10)
	publicBob := ecdh.Point{
		X: x,
		Y: y,
	}

	fmt.Println("Alice")
	fmt.Printf("Priv: %x\n", privateAlice)
	fmt.Printf("Pub: %#v\n\n", publicAlice)
	fmt.Printf("PubX: %s\n\n", publicAlice.(ecdh.Point).X.String())

	/*	fmt.Println("Bob")
		fmt.Printf("Priv: %x\n", privateBob)
		fmt.Printf("Pub: %#v\n\n", publicBob)
	*/
	if err := p256.Check(publicBob); err != nil {
		fmt.Printf("Bob's public key is not on the curve: %s\n", err)
	}
	secretAlice := p256.ComputeSecret(privateAlice, publicBob)

	fmt.Println("Alice")
	fmt.Printf("Secret: %x\n", secretAlice)

	/*if err := p256.Check(publicAlice); err != nil {
		fmt.Printf("Alice's public key is not on the curve: %s\n", err)
	}
	secretBob := p256.ComputeSecret(privateBob, publicAlice)

	fmt.Println("Bob")
	fmt.Printf("Secret: %x\n", secretBob)

	if !bytes.Equal(secretAlice, secretBob) {
		fmt.Printf("key exchange failed - secret X coordinates not equal\n")
	}*/
}
