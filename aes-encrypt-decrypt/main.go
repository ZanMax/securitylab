package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) string {
	dec, err := base64.StdEncoding.DecodeString(str)
	checkError(err)
	return string(dec)
}

func encryptMsg(msg, secretKey string) string {
	text := []byte(msg)
	key := []byte(getMD5Hash(secretKey))
	c, err := aes.NewCipher(key)
	checkError(err)
	gcm, err := cipher.NewGCM(c)
	checkError(err)
	nonce := make([]byte, gcm.NonceSize())
	encMsg := string(gcm.Seal(nonce, nonce, text, nil))
	bs64Msg := base64Encode(encMsg)
	return bs64Msg
}

func decryptMsg(encMsg, secretKey string) string {
	encMsg = base64Decode(encMsg)
	key := []byte(getMD5Hash(secretKey))
	c, err := aes.NewCipher(key)
	checkError(err)
	gcm, err := cipher.NewGCM(c)
	checkError(err)
	nonceSize := gcm.NonceSize()

	encMsgBytes := []byte(encMsg)

	nonce, encMsgBytes := encMsgBytes[:nonceSize], encMsgBytes[nonceSize:]
	decMsg, err := gcm.Open(nil, nonce, encMsgBytes, nil)
	decMsgStr := string(decMsg)
	return decMsgStr
}

func encryptFile(filename, secretKey string) {
	key := []byte(getMD5Hash(secretKey))
	content, err := ioutil.ReadFile(filename)
	checkError(err)
	c, err := aes.NewCipher(key)
	checkError(err)
	gcm, err := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	encMsg := gcm.Seal(nonce, nonce, content, nil)
	err = ioutil.WriteFile("enc_"+filename, encMsg, 0777)
	checkError(err)
}

func decryptFile(filename, secretKey string) {
	key := []byte(getMD5Hash(secretKey))
	contentEnc, err := ioutil.ReadFile(filename)
	checkError(err)
	c, err := aes.NewCipher(key)
	checkError(err)
	gcm, err := cipher.NewGCM(c)
	nonce := contentEnc[:gcm.NonceSize()]
	contentEnc = contentEnc[gcm.NonceSize():]
	contentDec, err := gcm.Open(nil, nonce, contentEnc, nil)
	checkError(err)
	err = ioutil.WriteFile("dec_"+filename, contentDec, 0777)
	checkError(err)
}

func main() {
	fmt.Println()
	fmt.Println("Encrypt/Decrypt messages and files v.0.1")
	fmt.Println()

	if len(os.Args) != 5 {
		fmt.Println("--> Encrypt/Decrypt messages")
		fmt.Println("Usage: ./main <enc | dec> msg <message> <secretKey>")
		fmt.Println("--> Encrypt/Decrypt files")
		fmt.Println("Usage: ./main <enc | dec> file <filename> <secretKey>")
		fmt.Println()
		return
	} else {
		if os.Args[1] == "enc" {
			if os.Args[2] == "msg" {
				fmt.Println("Encrypting message...")
				encMsg := encryptMsg(os.Args[3], os.Args[4])
				fmt.Println(encMsg)
			} else if os.Args[2] == "file" {
				fmt.Println("Encrypting file...")
				encryptFile(os.Args[3], os.Args[4])
			}
		} else if os.Args[1] == "dec" {
			if os.Args[2] == "msg" {
				fmt.Println("Decrypting message...")
				decMsg := decryptMsg(os.Args[3], os.Args[4])
				fmt.Println(decMsg)
			} else if os.Args[2] == "file" {
				fmt.Println("Decrypting file...")
				decryptFile(os.Args[3], os.Args[4])
			}
		}
	}
}
