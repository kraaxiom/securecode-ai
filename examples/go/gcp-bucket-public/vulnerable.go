// vulnerable.go - CWE-284: Improper Access Control
// Ce code accorde le role de lecture "storage.objectViewer" au principal "allUsers"
// sur un bucket GCS, rendant tous les objets du bucket lisibles publiquement.
package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"cloud.google.com/go/storage/control/apiv2/controlpb"
)

const bucketName = "REPLACE_WITH_YOUR_BUCKET"

// MakeBucketPublic accorde un acces public en lecture a tous les objets du bucket.
func MakeBucketPublic(ctx context.Context, client *storage.Client) error {
	bucket := client.Bucket(bucketName)

	policy, err := bucket.IAM().Policy(ctx)
	if err != nil {
		return fmt.Errorf("recuperation policy IAM: %w", err)
	}

	// VULNERABLE: allUsers = n'importe qui sur Internet, sans authentification
	policy.Add("allUsers", "roles/storage.objectViewer")

	if err := bucket.IAM().SetPolicy(ctx, policy); err != nil {
		return fmt.Errorf("application policy IAM: %w", err)
	}

	// VULNERABLE: uniform bucket-level access desactive, les ACL d'objet
	// individuelles peuvent rester plus permissives encore que la policy du bucket
	_, err = bucket.Update(ctx, storage.BucketAttrsToUpdate{
		UniformBucketLevelAccess: &storage.UniformBucketLevelAccess{Enabled: false},
	})
	if err != nil {
		return fmt.Errorf("mise a jour attributs bucket: %w", err)
	}

	_ = controlpb.Bucket{} // reference illustrative au package control (non utilise ici)
	return nil
}
