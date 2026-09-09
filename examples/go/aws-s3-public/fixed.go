// fixed.go - Correctif CWE-284: Improper Access Control
// Le bucket est cree en prive, le "Block Public Access" est active integralement,
// la bucket policy restreint l'acces a un role applicatif precis, et le partage
// ponctuel se fait via une URL pre-signee a duree limitee.
package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const (
	bucketName    = "REPLACE_WITH_YOUR_BUCKET"
	appRoleArn    = "arn:aws:iam::REPLACE_WITH_ACCOUNT_ID:role/app-read-role"
)

// CreatePrivateBucket cree un bucket prive et le securise selon le moindre privilege.
func CreatePrivateBucket(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return fmt.Errorf("chargement config AWS: %w", err)
	}
	client := s3.NewFromConfig(cfg)

	// SECURISE: creation du bucket sans ACL publique (ACL par defaut "private")
	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		ACL:    types.BucketCannedACLPrivate,
	})
	if err != nil {
		return fmt.Errorf("creation bucket: %w", err)
	}

	// SECURISE: activation integrale du "Block Public Access"
	_, err = client.PutPublicAccessBlock(ctx, &s3.PutPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       aws.Bool(true),
			IgnorePublicAcls:      aws.Bool(true),
			BlockPublicPolicy:     aws.Bool(true),
			RestrictPublicBuckets: aws.Bool(true),
		},
	})
	if err != nil {
		return fmt.Errorf("configuration public access block: %w", err)
	}

	// SECURISE: chiffrement au repos via SSE-KMS
	_, err = client.PutBucketEncryption(ctx, &s3.PutBucketEncryptionInput{
		Bucket: aws.String(bucketName),
		ServerSideEncryptionConfiguration: &types.ServerSideEncryptionConfiguration{
			Rules: []types.ServerSideEncryptionRule{
				{
					ApplyServerSideEncryptionByDefault: &types.ServerSideEncryptionByDefault{
						SSEAlgorithm: types.ServerSideEncryptionAwsKms,
					},
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("configuration chiffrement: %w", err)
	}

	// SECURISE: bucket policy restreinte a un role applicatif precis, avec TLS obligatoire
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [{
			"Effect": "Allow",
			"Principal": {"AWS": "%s"},
			"Action": ["s3:GetObject"],
			"Resource": ["arn:aws:s3:::%s/*"],
			"Condition": {"Bool": {"aws:SecureTransport": "true"}}
		}]
	}`, appRoleArn, bucketName)
	_, err = client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(policy),
	})
	if err != nil {
		return fmt.Errorf("application bucket policy: %w", err)
	}

	log.Println("bucket prive cree et securise")
	return nil
}

// GeneratePresignedURL genere une URL de partage temporaire au lieu d'un acces public permanent.
func GeneratePresignedURL(ctx context.Context, objectKey string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		return "", fmt.Errorf("chargement config AWS: %w", err)
	}
	client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(client)

	// SECURISE: URL valable 15 minutes seulement, en lecture seule
	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(15*time.Minute))
	if err != nil {
		return "", fmt.Errorf("generation URL pre-signee: %w", err)
	}
	return req.URL, nil
}
