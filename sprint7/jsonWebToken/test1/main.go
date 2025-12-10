package main

import "crypto/sha256"

func SignDocument(document []byte, secretKey []byte) [32]byte {
	return sha256.Sum256(append(document, secretKey...))
}

func CheckDocument(document []byte, secretKey []byte, hash [32]byte) bool {
	documentHash := SignDocument(document, secretKey)
	return documentHash == hash
}
