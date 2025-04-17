package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

var AESKey []byte

// SetKey gán khóa mã hóa AES (32 bytes cho AES-256)
func SetKey(secret string) error {
	if len(secret) != 32 {
		return fmt.Errorf("🔐 SESSION_SECRET_KEY phải đúng 32 ký tự (AES-256)")
	}
	AESKey = []byte(secret)
	return nil
}

// EncryptMap mã hóa map[string]string và trả về chuỗi base64
func EncryptMap(data map[string]string) (string, error) {
	plainBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesgcm.Seal(nil, nonce, plainBytes, nil)
	final := append(nonce, cipherText...)
	return base64.StdEncoding.EncodeToString(final), nil
}

// DecryptMap giải mã chuỗi base64 thành map[string]string
func DecryptMap(encoded string) (map[string]string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	nonce := decoded[:12]
	cipherText := decoded[12:]

	block, err := aes.NewCipher(AESKey)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]string
	err = json.Unmarshal(plainText, &result)
	return result, err
}
