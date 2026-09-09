// vulnerable.go - CWE-284: Improper Access Control
// Ce code cree/configure un bucket S3 avec une ACL publique et une bucket policy
// autorisant n'importe quel principal a lire les objets, exposant potentiellement
// des donnees sensibles a Internet.
package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const bucketName = "REPLACE_WITH_YOUR_BUCKET"

// CreatePublicBucket cree un bucket et l'expose publiquement.
func CreatePublicBucket(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := s3.NewFromConfig(cfg)

	// VULNERABLE: creation du bucket avec une ACL "public-read"
	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		ACL:    types.BucketCannedACLPublicRead, // n'importe qui peut lister/lire les objets
	})
	if err != nil {
		return fmt.Errorf("creation bucket: %w", err)
	}

	// VULNERABLE: desactivation totale du "Block Public Access"
	_, err = client.PutPublicAccessBlock(ctx, &s3.PutPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       aws.Bool(false),
			IgnorePublicAcls:      aws.Bool(false),
			BlockPublicPolicy:     aws.Bool(false),
			RestrictPublicBuckets: aws.Bool(false),
		},
	})
	if err != nil {
		return fmt.Errorf("configuration public access block: %w", err)
	}

	// VULNERABLE: bucket policy avec Principal "*" sans condition restrictive
	policy := `{
		"Version": "2012-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": "*",
			"Action": ["s3:GetObject", "s3:ListBucket"],
			"Resource": ["arn:aws:s3:::REPLACE_WITH_YOUR_BUCKET", "arn:aws:s3:::REPLACE_WITH_YOUR_BUCKET/*"]
		}]
	}`
	_, err = client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(policy),
	})
	if err != nil {
		return fmt.Errorf("application bucket policy: %w", err)
	}

	log.Println("bucket public cree (non securise)")
	return nil
}
