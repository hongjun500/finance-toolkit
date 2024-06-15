// @author hongjun500
// @date 2024-06-15 22:07
// @tool ThinkPadX1 隐士
// Created with GoLand 2022.2
// Description: crypt_common_test.go

package common

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := "thisis32bitlongpassphraseimusing" // Key should be 16, 24, or 32 bytes long
	plaintext := "Hello, World!"
	fmt.Printf("Original text: %s\n", plaintext)
	encrypted, err := encrypt(key, plaintext)
	if err != nil {
		t.Errorf("Encryption failed: %v\n", err)
	}
	println("Encrypted text:", encrypted)
	decrypted, err := decrypt(key, encrypted)
	if err != nil {
		t.Errorf("Decryption failed: %v\n", err)
	}
	t.Log("Decrypted text:", decrypted)
	if decrypted != plaintext {
		t.Errorf("Expected %s, got %s\n", plaintext, decrypted)
	}

}
