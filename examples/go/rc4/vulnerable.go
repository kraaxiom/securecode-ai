package main

import (
	"crypto/rc4"
	"fmt"
	"log"
)

// VULNÉRABLE : CWE-327 — Use of a Broken or Risky Cryptographic Algorithm
//
// RC4 est un algorithme de chiffrement par flux comportant des biais
// statistiques connus dans son flux de clés (key stream), exploités par des
// attaques pratiques (ex. contre WEP, contre TLS avec RC4). Il ne fournit ni
// intégrité, ni authentification, et est aujourd'hui considéré comme cassé.
func chiffrerFluxDonnees(donnees []byte, cle []byte) ([]byte, error) {
	c, err := rc4.NewCipher(cle)
	if err != nil {
		return nil, err
	}
	chiffre := make([]byte, len(donnees))
	c.XORKeyStream(chiffre, donnees)
	return chiffre, nil
}

func main() {
	cle := []byte("cle-rc4-faible")
	chiffre, err := chiffrerFluxDonnees([]byte("donnees confidentielles"), cle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Données chiffrées avec RC4 (faible) : %x\n", chiffre)
}
