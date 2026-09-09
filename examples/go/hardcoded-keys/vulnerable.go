package main

import (
	"fmt"
	"net/http"
)

// VULNÉRABLE : CWE-798 — Use of Hard-coded Credentials
//
// Les clés API et secrets sont écrits en clair dans le code source. Quiconque
// a accès au dépôt (y compris via l'historique git ou une fuite de build)
// récupère des identifiants valides et permanents, sans rotation possible
// sans redéploiement.
const (
	stripeSecretKey = "sk_live_EXAMPLE_NOT_A_REAL_KEY" // clé de paiement en dur
	awsAccessKey    = "AKIAIOSFODNN7EXAMPLE"
	awsSecretKey    = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	dbPassword      = "SuperSecretP@ssw0rd!" // mot de passe base de données en dur
)

func appelerAPIStripe(montant int) error {
	req, err := http.NewRequest("POST", "https://api.stripe.com/v1/charges", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+stripeSecretKey)
	_, err = http.DefaultClient.Do(req)
	return err
}

func main() {
	fmt.Println("Connexion AWS avec clé en dur :", awsAccessKey)
	fmt.Println("Mot de passe DB en dur :", dbPassword)
	if err := appelerAPIStripe(1000); err != nil {
		fmt.Println("erreur:", err)
	}
	_ = awsSecretKey
}
