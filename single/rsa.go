package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Parse PEM encoded PKCS1 or PKCS8 public key
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("ErrKeyMustBePEMEncoded")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, errors.New("ErrNotRSAPublicKey")
	}

	return pkey, nil
}

func main() {
	var hash crypto.Hash

	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJmb28iOiJiYXIifQ.FhkiHkoESI_cG3NPigFrxEk9Z60_oXrOT2vGm9Pn6RDgYNovYORQmmA0zs1AoAOf09ly2Nx2YAg6ABqAYga1AcMFkJljwxTT5fYphTuqpWdy4BELeSYJx5Ty2gmr8e7RonuUztrdD5WfPqLKMm1Ozp_T6zALpRmwTIW0QPnaBXaQD90FplAg46Iy1UlDKr-Eupy0i5SLch5Q-p2ZpaL_5fnTIUDlxC3pWhJTyx_71qDI-mAA_5lE_VdroOeflG56sSmDxopPEG3bFlSu1eowyBfxtu0_CuVd-M42RU75Zc4Gsj6uV77MBtbMrf4_7M_NUTSgoIF3fRqxrj0NzihIBg"

	parts := strings.Split(token, ".")

	// pubk := []byte(`-----BEGIN PUBLIC KEY-----
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
	// fGnJm6gOdrj8ym3rFkEU/wT8RDtnSgFEZOQpHEgQ7JL38xUfU0Y3g6aYw9QT0hJ7
	// mCpz9Er5qLaMXJwZxzHzAahlfA0icqabvJOMvQtzD6uQv6wPEyZtDTWiQi9AXwBp
	// HssPnpYGIn20ZZuNlX2BrClciHhCPUIIZOQn/MmqTD31jSyjoQoV7MhhMTATKJx2
	// XrHhR+1DcKJzQBSTAGnpYVaqpsARap+nwRipr3nUTuxyGohBTSmjJ2usSeQXHI3b
	// ODIRe1AuTyHceAbewn8b462yEWKARdpd9AjQW5SIVPfdsz5B6GlYQ5LdYKtznTuy
	// 7wIDAQAB
	// -----END PUBLIC KEY-----`)

	keyData, err := ioutil.ReadFile("test/sample_key.pub")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("KeYDATA:\n%s\n", string(keyData))
	rsaPubKey, err := ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		log.Fatal(err)
	}

	hash = crypto.SHA256

	// base64Payload, err := strconv.Unquote(parts[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	payload, err := base64.URLEncoding.DecodeString(parts[1] + "==")
	fmt.Printf("payload: %s\n", payload)
	if err != nil {
		fmt.Printf("%#v\n", parts[1])

		log.Fatal(err)
	}
	//base64.URLEncoding.DecodeString()
	signature, err := base64.URLEncoding.DecodeString(parts[2] + "==")
	fmt.Printf("sig: %s\n", signature)
	if err != nil {
		fmt.Println("signature")
		log.Fatal(err)
	}

	//signature := []byte(parts[2])
	//payload := []byte(parts[1]) //strings.Join(parts[0:2], ".")

	hasher := hash.New()

	// According to documentation, Write() on hash never fails
	_, _ = hasher.Write(payload)
	hashed := hasher.Sum(nil)

	err = rsa.VerifyPKCS1v15(rsaPubKey, hash, hashed, signature)
	fmt.Println(err)
}
