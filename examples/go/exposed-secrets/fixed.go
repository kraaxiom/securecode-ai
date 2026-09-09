// fixed.go - Correctif CWE-798: Use of Hard-coded Credentials
// Les secrets ne sont plus codes en dur : ils sont recuperes a l'execution
// depuis un gestionnaire de secrets dedie (AWS Secrets Manager), jamais journalises,
// et jamais stockes en clair dans le code source ou un fichier versionne.
package payment

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

// secretRef designe une reference vers un secret stocke dans le gestionnaire dedie,
// jamais la valeur elle-meme.
const (
	stripeSecretName = "prod/stripe/secret_key"
	dbSecretName      = "prod/db/credentials"
)

// getSecret recupere la valeur d'un secret depuis AWS Secrets Manager a l'execution.
func getSecret(ctx context.Context, secretName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("chargement config AWS: %w", err)
	}
	client := secretsmanager.NewFromConfig(cfg)

	out, err := client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return "", fmt.Errorf("recuperation secret %s: %w", secretName, err)
	}
	return aws.ToString(out.SecretString), nil
}

func InitPaymentClient(ctx context.Context) (string, error) {
	stripeSecretKey, err := getSecret(ctx, stripeSecretName)
	if err != nil {
		return "", err
	}
	// SECURISE: on ne journalise jamais la valeur du secret, seulement une confirmation
	log.Println("client Stripe initialise depuis le gestionnaire de secrets")
	return stripeSecretKey, nil
}

func ConnectDatabase(ctx context.Context) (*sql.DB, error) {
	// SECURISE: les identifiants sont recuperes depuis le gestionnaire de secrets,
	// jamais codes en dur dans le code source
	credsJSON, err := getSecret(ctx, dbSecretName)
	if err != nil {
		return nil, err
	}
	// credsJSON contient typiquement {"host":..., "user":..., "password":...}
	// a parser avec encoding/json avant construction du DSN.
	dsn := buildDSNFromSecretJSON(credsJSON)
	return sql.Open("mysql", dsn)
}

// buildDSNFromSecretJSON construit le DSN a partir du JSON de secret recupere.
// L'implementation complete du parsing JSON est omise ici par souci de concision.
func buildDSNFromSecretJSON(credsJSON string) string {
	// SECURISE: aucune valeur sensible en dur dans le code, tout provient du coffre.
	return credsJSON // placeholder illustratif : voir encoding/json en production
}
