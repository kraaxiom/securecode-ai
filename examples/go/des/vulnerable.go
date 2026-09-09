package main

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"log"
)

// VULNÉRABLE : CWE-327 — Use of a Broken or Risky Cryptographic Algorithm
//
// DES (et 3DES) utilisent un bloc/clé trop faibles (56 bits effectifs pour DES,
// espace de clé réduit et vulnérable à Sweet32 pour 3DES en mode CBC sur de gros
// volumes de données). Un attaquant disposant de ressources de calcul modestes
// peut casser la clé par force brute, ou exploiter les collisions de bloc de 64
// bits (Sweet32) pour retrouver du texte clair dans un flux chiffré volumineux.
func chiffrerDonneesSensibles(donnees []byte, cle []byte) ([]byte, error) {
	// cle doit faire exactement 8 octets pour DES — espace de clé bien trop petit
	block, err := des.NewCipher(cle)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, des.BlockSize)
	// IV nul et algorithme DES : double faiblesse
	mode := cipher.NewCBCEncrypter(block, iv)

	// Padding simplifié pour l'exemple
	padded := padPKCS7(donnees, des.BlockSize)
	chiffre := make([]byte, len(padded))
	mode.CryptBlocks(chiffre, padded)

	return chiffre, nil
}

func padPKCS7(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := make([]byte, padLen)
	for i := range padding {
		padding[i] = byte(padLen)
	}
	return append(data, padding...)
}

func main() {
	cle := []byte("8octets!") // clé DES codée en dur, 64 bits (56 bits effectifs)
	chiffre, err := chiffrerDonneesSensibles([]byte("numero_carte_bancaire"), cle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Données chiffrées avec DES (faible) : %x\n", chiffre)
}
