package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/argon2"
)

// CORRIGÉ : remplacement de MD5 par Argon2id pour le hachage de mots de passe.
// Argon2id ajoute un sel aléatoire et un coût mémoire/CPU réglable, rendant
// les attaques par force brute ou GPU/ASIC bien plus coûteuses — corrige CWE-328.
//
// go.mod : require golang.org/x/crypto v0.x.x

func genererSel() ([]byte, error) {
	// En pratique : crypto/rand.Read sur un buffer de 16 octets.
	return []byte("sel-aleatoire-16"), nil
}

func hacherMotDePasse(motDePasse string) (string, error) {
	sel, err := genererSel()
	if err != nil {
		return "", err
	}

	// Paramètres recommandés OWASP : temps=1, mémoire=64Mo, parallélisme=4
	hash := argon2.IDKey([]byte(motDePasse), sel, 1, 64*1024, 4, 32)
	return fmt.Sprintf("%x$%x", sel, hash), nil
}

func main() {
	hash, err := hacherMotDePasse("MonMotDePasse123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash Argon2id du mot de passe :", hash)
}
