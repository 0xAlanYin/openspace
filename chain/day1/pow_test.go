package day1

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
	"time"
)

// 实践 POW， 编写程序（编程语言不限）用自己的昵称 + nonce，不断进行 sha256 Hash 运算：
// 1.直到满足 4 个 0 开头的哈希值，打印出花费的时间。
// 2.再次运算直到满足 5 个 0 开头的哈希值，打印出花费的时间。
func TestPow(t *testing.T) {
	nickname := "Alan"
	targetPrefix := "0000" // 目标哈希值的前缀

	// Case1
	startTime := time.Now()
	nonce := ""
	hash := ""
	for !strings.HasPrefix(hash, targetPrefix) {
		nonce = generateNonce()
		data := fmt.Sprintf("%s%s", nickname, nonce)
		hash = calculateHash(data)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("find %s begin hash：\n", targetPrefix)
	fmt.Printf("hashValue: %s\n", hash)
	fmt.Printf("cost: %s\n\n", elapsedTime)

	// Case2
	targetPrefix = "00000" // 新的目标哈希值前缀
	startTime = time.Now()
	nonce = ""
	hash = ""
	for !strings.HasPrefix(hash, targetPrefix) {
		nonce = generateNonce()
		data := fmt.Sprintf("%s%s", nickname, nonce)
		hash = calculateHash(data)
	}
	elapsedTime = time.Since(startTime)
	fmt.Printf("find %s begin hash：\n", targetPrefix)
	fmt.Printf("hashValue: %s\n", hash)
	fmt.Printf("cost: %s\n\n", elapsedTime)
}

func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
