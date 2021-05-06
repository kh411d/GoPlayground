package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {
	currentTime := time.Now()

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	clientID := "afe4ffe3-3004-4ece-8e31-e687fb384eaf"
	clientSecret := "22b51175-d5e9-4625-9490-62adf9477faa"
	client := &http.Client{Transport: transCfg}
	req, err := http.NewRequest("GET", "https://172.24.1.142:5900/rest/mdr/api/auth/jwtToken", nil)
	// req.Header.Set("X-TIMESTAMP", currentTime.Format("2006-01-02 15:04:05"))
	// req.Header.Set("X-Mandiri-Key", clientID)
	// req.Header.Set("X-SIGNATURE", generateSignature(clientID, clientSecret, currentTime.Format("2006-01-02 15:04:05")))
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("User-Agent", "linkaja-net/http")
	// req.Header.Set("cache-control", "no-cache")

	req.Header = http.Header{
		"X-TIMESTAMP":   []string{currentTime.Format("2006-01-02 15:04:05")},
		"Content-Type":  []string{"application/json"},
		"X-SIGNATURE":   []string{generateSignature(clientID, clientSecret, currentTime.Format("2006-01-02 15:04:05"))},
		"X-Mandiri-Key": []string{clientID},
	}

	res, err := client.Do(req)

	dumpr, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%q\n\n", dumpr)

	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", dump)

	/*fmt.Printf("Request %#v\n", req)
	fmt.Printf("Response %#v\n", res)
	fmt.Println(err)

	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(err)
	fmt.Println(string(body))*/

}

func generateSignature(clientID, secretKey, timestamp string) string {
	payload := fmt.Sprintf("%s|%s", clientID, timestamp)

	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(payload))
	return fmt.Sprintf("%x", h.Sum(nil))
}
