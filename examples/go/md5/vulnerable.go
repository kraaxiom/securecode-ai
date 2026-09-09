package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// VULNÉRABLE : CWE-328 — Use of Weak Hash
//
// MD5 est cryptographiquement cassé (collisions praticables) et surtout bien
// trop rapide à calculer : un attaquant peut tester des milliards de mots de
// passe par seconde sur GPU pour retrouver le mot de passe original à partir
// du hash (attaque par force brute / rainbow tables), aucun sel n'étant utilisé.
func hacherMotDePasse(motDePasse string) string {
	somme := md5.Sum([]byte(motDePasse))
	return hex.EncodeToString(somme[:])
}

func main() {
	hash := hacherMotDePasse("MonMotDePasse123")
	fmt.Println("Hash MD5 du mot de passe (faible) :", hash)
}
