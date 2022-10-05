package main

import "fmt"
// import "time"
import "encoding/base64"
import "crypto/hmac"
import "crypto/sha256"

func main() {
	sec := "LhZPfECd7XVYFDegLLCOnb"
	// timestamp := time.Now().Unix()
	timestamp := int64(1664944271)
	sig, _ := GenSign(sec, timestamp)
	fmt.Printf("%v sign: %s\n",  timestamp, sig)
}

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		 return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}