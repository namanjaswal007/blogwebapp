package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt text using AES-GCM
func Encrypt(plainText *string) (err error) {
	block, err := aes.NewCipher([]byte(CypherKey))
	if err != nil {
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(*plainText), nil)
	*plainText = base64.URLEncoding.EncodeToString(cipherText)
	return
}

// Decrypt ciphertext using AES-GCM
func Decrypt(encryptedText string) (string, error) {
	cipherData, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(CypherKey))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherData) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherText := cipherData[:nonceSize], cipherData[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
