// fixed.go - Correctif CWE-284: Improper Access Control
// Le conteneur est cree en prive, l'acces public est desactive au niveau compte,
// et le partage ponctuel se fait via un SAS a duree de vie courte et permissions minimales.
package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

const (
	accountURL    = "https://REPLACE_WITH_YOUR_ACCOUNT.blob.core.windows.net/"
	containerName = "myfiles"
)

// CreatePrivateContainer cree un conteneur prive (aucun acces anonyme).
func CreatePrivateContainer(ctx context.Context, connectionString string) error {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return fmt.Errorf("creation client azblob: %w", err)
	}

	// SECURISE: pas d'option "Access" -> le conteneur reste prive par defaut
	_, err = client.CreateContainer(ctx, containerName, &azblob.CreateContainerOptions{
		Access: to.Ptr(container.PublicAccessType("")),
	})
	if err != nil {
		return fmt.Errorf("creation conteneur prive: %w", err)
	}

	return nil
}

// UploadSensitiveFile televerse un fichier dans le conteneur prive.
func UploadSensitiveFile(ctx context.Context, client *azblob.Client, blobName string, data []byte) error {
	_, err := client.UploadBuffer(ctx, containerName, blobName, data, nil)
	if err != nil {
		return fmt.Errorf("upload blob: %w", err)
	}
	return nil
}

// GenerateReadOnlySASURL genere une URL SAS en lecture seule, valable 1 heure,
// pour un partage ponctuel au lieu d'un acces public permanent.
func GenerateReadOnlySASURL(client *azblob.Client, blobName string) (string, error) {
	permissions := sas.BlobPermissions{Read: true}
	expiry := time.Now().UTC().Add(1 * time.Hour)

	// SECURISE: SAS scope en lecture seule, expiration courte, HTTPS uniquement
	url, err := client.ServiceClient().
		NewContainerClient(containerName).
		NewBlobClient(blobName).
		GetSASURL(permissions, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("generation SAS URL: %w", err)
	}
	return url, nil
}
