// fixed.go - Correctif CWE-269: Improper Privilege Management
// La policy IAM est remplacee par une policy granulaire respectant le moindre
// privilege, et la trust policy exige un compte precis, un ExternalId et le MFA.
package iamsetup

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const (
	roleName      = "app-service-role"
	trustedAcctID = "REPLACE_WITH_TRUSTED_ACCOUNT_ID"
	externalID    = "REPLACE_WITH_UNIQUE_EXTERNAL_ID"
)

type policyDocument struct {
	Version   string            `json:"Version"`
	Statement []policyStatement `json:"Statement"`
}

type policyStatement struct {
	Effect    string                 `json:"Effect"`
	Action    interface{}            `json:"Action"`
	Resource  interface{}            `json:"Resource"`
	Condition map[string]interface{} `json:"Condition,omitempty"`
}

// AttachLeastPrivilegePolicy attache une policy IAM granulaire, limitee aux
// actions et ressources reellement necessaires a l'application.
func AttachLeastPrivilegePolicy(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := iam.NewFromConfig(cfg)

	// SECURISE: permissions scopees a l'action et la ressource necessaires,
	// avec une condition supplementaire sur la region
	doc := policyDocument{
		Version: "2012-10-17",
		Statement: []policyStatement{
			{
				Effect:   "Allow",
				Action:   []string{"s3:GetObject", "s3:PutObject"},
				Resource: "arn:aws:s3:::REPLACE_WITH_YOUR_BUCKET/uploads/*",
				Condition: map[string]interface{}{
					"StringEquals": map[string]string{"aws:RequestedRegion": "eu-west-1"},
				},
			},
		},
	}
	policyJSON, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("serialisation policy: %w", err)
	}

	_, err = client.PutRolePolicy(ctx, &iam.PutRolePolicyInput{
		RoleName:       aws.String(roleName),
		PolicyName:     aws.String("least-privilege-policy"),
		PolicyDocument: aws.String(string(policyJSON)),
	})
	if err != nil {
		return fmt.Errorf("attachement policy: %w", err)
	}

	return nil
}

// CreateRestrictedTrustPolicy limite l'assomption du role a un compte precis,
// avec un ExternalId et l'exigence du MFA.
func CreateRestrictedTrustPolicy(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := iam.NewFromConfig(cfg)

	// SECURISE: compte precis, ExternalId unique, MFA obligatoire
	trustPolicy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": {"AWS": "arn:aws:iam::%s:root"},
			"Action": "sts:AssumeRole",
			"Condition": {
				"StringEquals": {"sts:ExternalId": "%s"},
				"Bool": {"aws:MultiFactorAuthPresent": "true"}
			}
		}]
	}`, trustedAcctID, externalID)

	_, err = client.UpdateAssumeRolePolicy(ctx, &iam.UpdateAssumeRolePolicyInput{
		RoleName:       aws.String(roleName),
		PolicyDocument: aws.String(trustPolicy),
	})
	if err != nil {
		return fmt.Errorf("mise a jour trust policy: %w", err)
	}

	return nil
}
