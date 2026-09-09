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

// CORRIGÉ : remplacement de DES par AES-256 en mode GCM (chiffrement authentifié).
// La clé est chargée depuis l'environnement (jamais codée en dur), et un nonce
// unique et aléatoire est généré à chaque chiffrement — corrige CWE-327.
func chiffrerDonneesSensibles(donnees []byte, cle []byte) ([]byte, error) {
	if len(cle) != 32 {
		return nil, fmt.Errorf("la clé AES-256 doit faire 32 octets")
	}

	block, err := aes.NewCipher(cle)
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

	// Le nonce est préfixé au texte chiffré ; le tag d'authentification est
	// intégré automatiquement par Seal.
	chiffre := gcm.Seal(nonce, nonce, donnees, nil)
	return chiffre, nil
}

func chargerCleDepuisEnv() ([]byte, error) {
	hexCle := os.Getenv("APP_ENCRYPTION_KEY") // 64 caractères hex = 32 octets
	if hexCle == "" {
		return nil, fmt.Errorf("APP_ENCRYPTION_KEY manquant dans l'environnement")
	}
	cle := make([]byte, 32)
	// Dans un projet réel : hex.Decode ou récupération depuis un gestionnaire
	// de secrets (Vault, AWS KMS, etc.). Simplifié ici pour l'exemple.
	copy(cle, []byte(hexCle))
	return cle, nil
}

func main() {
	cle, err := chargerCleDepuisEnv()
	if err != nil {
		log.Fatal(err)
	}
	chiffre, err := chiffrerDonneesSensibles([]byte("numero_carte_bancaire"), cle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Données chiffrées avec AES-256-GCM : %x\n", chiffre)
}
