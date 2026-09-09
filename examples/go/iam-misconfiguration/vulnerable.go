// vulnerable.go - CWE-269: Improper Privilege Management
// Ce code cree une policy IAM AWS accordant toutes les actions sur toutes les
// ressources ("Action": "*", "Resource": "*") a un role applicatif, violant
// le principe du moindre privilege.
package iamsetup

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

const roleName = "app-service-role"

type policyDocument struct {
	Version   string            `json:"Version"`
	Statement []policyStatement `json:"Statement"`
}

type policyStatement struct {
	Effect   string      `json:"Effect"`
	Action   interface{} `json:"Action"`
	Resource interface{} `json:"Resource"`
}

// AttachOverlyBroadPolicy attache une policy IAM avec des permissions completes.
func AttachOverlyBroadPolicy(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := iam.NewFromConfig(cfg)

	// VULNERABLE: wildcard total sur les actions ET les ressources
	doc := policyDocument{
		Version: "2012-10-17",
		Statement: []policyStatement{
			{Effect: "Allow", Action: "*", Resource: "*"},
		},
	}
	policyJSON, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("serialisation policy: %w", err)
	}

	_, err = client.PutRolePolicy(ctx, &iam.PutRolePolicyInput{
		RoleName:       aws.String(roleName),
		PolicyName:     aws.String("full-access-policy"),
		PolicyDocument: aws.String(string(policyJSON)),
	})
	if err != nil {
		return fmt.Errorf("attachement policy: %w", err)
	}

	return nil
}

// CreateOpenTrustPolicy autorise n'importe quel compte AWS a assumer le role.
func CreateOpenTrustPolicy(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := iam.NewFromConfig(cfg)

	// VULNERABLE: Principal "*" sans condition -> n'importe quel compte AWS
	// peut assumer ce role
	trustPolicy := `{
		"Version": "2012-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": {"AWS": "*"},
			"Action": "sts:AssumeRole"
		}]
	}`

	_, err = client.UpdateAssumeRolePolicy(ctx, &iam.UpdateAssumeRolePolicyInput{
		RoleName:       aws.String(roleName),
		PolicyDocument: aws.String(trustPolicy),
	})
	if err != nil {
		return fmt.Errorf("mise a jour trust policy: %w", err)
	}

	return nil
}
