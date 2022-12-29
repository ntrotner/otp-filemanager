package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"os"
)

var Block cipher.Block

func InitializeSecurity(input_key *string) {
	if *input_key == "" {
		log.Println("DISABLED ENCRYPTION DUE TO INVALID KEY PROVIDED")
		return
	}

	key := make([]byte, 32)
	input_key_length := len(*input_key)

	for i := 0; i < 32; i++ {
		if input_key_length <= i {
			log.Println("Ignoring", i, "character for encryption and placing 0 instead")
			key[i] = byte(0)
		} else {
			key[i] = byte((*input_key)[i])
		}
	}

	var err error
	Block, err = aes.NewCipher(key)

	if err != nil {
		log.Println("Failed to create Cipher Block")
		os.Exit(1)
	}
}

func Encrypt(toEncrypt *[]byte) (*[]byte, error) {
	if Block == nil {
		return toEncrypt, nil
	}

	ciphertext := make([]byte, aes.BlockSize+len(*toEncrypt))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(Block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], *toEncrypt)

	return &ciphertext, nil
}

func Decrypt(toDecrypt *[]byte) (*[]byte, error) {
	if Block == nil {
		return toDecrypt, nil
	}

	if len(*toDecrypt) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := (*toDecrypt)[:aes.BlockSize]
	*toDecrypt = (*toDecrypt)[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(Block, iv)
	stream.XORKeyStream(*toDecrypt, *toDecrypt)

	return toDecrypt, nil
}
