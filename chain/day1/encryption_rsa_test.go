package day1

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
)

// 实践非对称加密 RSA（编程语言不限）：
//
// 1.先生成一个公私钥对
// 2.用私钥对符合 POW 4个开头的哈希值的 “昵称 + nonce” 进行私钥签名
// 3.用公钥验证
func Test(t *testing.T) {
	nickname := "Alan"
	targetPrefix := "0000" // 目标哈希值的前缀

	hash := ""
	nonce := generateNonce()
	for strings.HasPrefix(hash, targetPrefix) {
		nonce = generateNonce()
		data := fmt.Sprintf("%s%s", nickname, nonce)
		hash = calculateHash(data)
	}

	privateKey, publicKey := generateKeyPair()
	hashed := sha256.Sum256([]byte(nickname + nonce))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("Failed to sign the message:", err)
		return
	}
	fmt.Println("Signature:", signature)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		t.Errorf("Signature verification failed: %v", err)
	}

	fmt.Println("Signature verification successful")
}

func generateNonce() string {
	nonceBytes := make([]byte, 8)
	// ignore error
	_, _ = rand.Read(nonceBytes)
	return fmt.Sprintf("%x", nonceBytes)
}

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return nil, nil
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey
}
