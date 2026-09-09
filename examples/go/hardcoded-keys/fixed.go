package main

import (
	"fmt"
	"net/http"
	"os"
)

// CORRIGÉ : les secrets sont chargés depuis des variables d'environnement
// (ou, en production, un gestionnaire de secrets comme AWS Secrets Manager /
// HashiCorp Vault). Aucun littéral sensible ne subsiste dans le code source —
// corrige CWE-798.
func chargerSecretRequis(nomVar string) (string, error) {
	valeur := os.Getenv(nomVar)
	if valeur == "" {
		return "", fmt.Errorf("%s manquant dans l'environnement", nomVar)
	}
	return valeur, nil
}

func appelerAPIStripe(montant int) error {
	stripeSecretKey, err := chargerSecretRequis("STRIPE_SECRET_KEY")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.stripe.com/v1/charges", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+stripeSecretKey)
	_, err = http.DefaultClient.Do(req)
	return err
}

func main() {
	awsAccessKey, err := chargerSecretRequis("AWS_ACCESS_KEY_ID")
	if err != nil {
		fmt.Println("erreur de configuration:", err)
		os.Exit(1)
	}
	fmt.Println("Connexion AWS avec clé chargée depuis l'environnement :", awsAccessKey)

	if err := appelerAPIStripe(1000); err != nil {
		fmt.Println("erreur:", err)
	}
	// Rappel : le secret précédemment exposé en dur doit être révoqué et
	// régénéré côté fournisseur, et purgé de l'historique git si committé.
}
