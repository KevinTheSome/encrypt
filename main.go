package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	if os.Args == nil || len(os.Args) != 4 {
		log.Println("1 argument is encrypt or decrypt, 2 argument is file path 3 with 'your key' key must be 16, 24 ,32 characters")
		log.Println("Must look like this: go run . \"encrypt\" \"file.txt\" \"1234567891011121\"")
		log.Println("OR")
		log.Println("Must look like this: go run . \"decrypt\" \"file.txt\" \"1234567891011121\"")
	}else{
		arg1 := os.Args[1]
		if arg1 == "encrypt" {
			encryptFile()
			log.Println("File Encrypted")
		} else if arg1 == "decrypt" {
			decryptFile()
			log.Println("File Decrypted")
		}
	}
}

func encryptFile() {
	// Reading plaintext file
	plainText, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Reading key
	key := []byte(os.Args[3])
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}
	
	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	
	// Writing ciphertext file
	err = os.WriteFile(os.Args[2], cipherText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

}

func decryptFile() {
	// Reading ciphertext file
	cipherText, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Reading key
	key := []byte(os.Args[3])
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}
	
	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}
	
	// Writing decryption content
	err = os.WriteFile(os.Args[2], plainText, 0777)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}
}