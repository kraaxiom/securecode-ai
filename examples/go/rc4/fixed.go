package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

// CORRIGÉ : remplacement de RC4 par AES-256-GCM, un chiffrement authentifié
// moderne qui garantit à la fois confidentialité et intégrité, sans les
// faiblesses statistiques de RC4 — corrige CWE-327.
func chiffrerFluxDonnees(donnees []byte, cle []byte) ([]byte, error) {
	block, err := aes.NewCipher(cle) // cle de 32 octets = AES-256
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, donnees, nil), nil
}

func main() {
	hexCle := os.Getenv("APP_ENCRYPTION_KEY")
	if hexCle == "" || len(hexCle) < 32 {
		log.Fatal("APP_ENCRYPTION_KEY manquant ou trop court dans l'environnement")
	}
	cle := []byte(hexCle)[:32]

	chiffre, err := chiffrerFluxDonnees([]byte("donnees confidentielles"), cle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Données chiffrées avec AES-256-GCM : %x\n", chiffre)
}
