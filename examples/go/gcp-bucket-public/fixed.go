// fixed.go - Correctif CWE-284: Improper Access Control
// Le binding IAM public est retire, l'uniform bucket-level access est active,
// et le partage ponctuel utilise une URL signee a duree limitee.
package storage

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

const (
	bucketName        = "REPLACE_WITH_YOUR_BUCKET"
	appReaderAccount  = "app-reader@REPLACE_WITH_YOUR_PROJECT.iam.gserviceaccount.com"
)

// SecureBucketAccess retire l'acces public et applique le moindre privilege.
func SecureBucketAccess(ctx context.Context, client *storage.Client) error {
	bucket := client.Bucket(bucketName)

	policy, err := bucket.IAM().Policy(ctx)
	if err != nil {
		return fmt.Errorf("recuperation policy IAM: %w", err)
	}

	// SECURISE: retrait du binding public
	policy.Remove("allUsers", "roles/storage.objectViewer")
	policy.Remove("allAuthenticatedUsers", "roles/storage.objectViewer")

	// SECURISE: acces scope a un service account applicatif precis
	policy.Add("serviceAccount:"+appReaderAccount, "roles/storage.objectViewer")

	if err := bucket.IAM().SetPolicy(ctx, policy); err != nil {
		return fmt.Errorf("application policy IAM: %w", err)
	}

	// SECURISE: uniform bucket-level access active pour eviter des ACL d'objet incoherentes
	_, err = bucket.Update(ctx, storage.BucketAttrsToUpdate{
		UniformBucketLevelAccess: &storage.UniformBucketLevelAccess{Enabled: true},
		PublicAccessPrevention:   storage.PublicAccessPreventionEnforced,
	})
	if err != nil {
		return fmt.Errorf("mise a jour attributs bucket: %w", err)
	}

	return nil
}

// GenerateSignedURL genere une URL signee a duree limitee pour un partage ponctuel
// au lieu d'un acces public permanent.
func GenerateSignedURL(bucket *storage.BucketHandle, objectName string, opts *storage.SignedURLOptions) (string, error) {
	// SECURISE: URL valable 30 minutes, methode GET uniquement
	opts.Method = "GET"
	opts.Expires = time.Now().Add(30 * time.Minute)

	url, err := bucket.SignedURL(objectName, opts)
	if err != nil {
		return "", fmt.Errorf("generation URL signee: %w", err)
	}
	return url, nil
}
