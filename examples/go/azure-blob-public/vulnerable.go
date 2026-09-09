// vulnerable.go - CWE-284: Improper Access Control
// Ce code cree un conteneur Azure Blob Storage avec un niveau d'acces public,
// permettant a quiconque de lire (et lister) les blobs sans authentification.
package storage

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

const (
	accountURL    = "https://REPLACE_WITH_YOUR_ACCOUNT.blob.core.windows.net/"
	containerName = "myfiles"
)

// CreatePublicContainer cree un conteneur avec un acces public en lecture au niveau "container".
func CreatePublicContainer(ctx context.Context, connectionString string) error {
	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		return fmt.Errorf("creation client azblob: %w", err)
	}

	// VULNERABLE: niveau d'acces public "Container" -> lecture anonyme des blobs ET listing
	_, err = client.CreateContainer(ctx, containerName, &azblob.CreateContainerOptions{
		Access: to.Ptr(container.PublicAccessTypeContainer),
	})
	if err != nil {
		return fmt.Errorf("creation conteneur public: %w", err)
	}

	return nil
}

// UploadSensitiveFile televerse un fichier dans le conteneur public, l'exposant de fait publiquement.
func UploadSensitiveFile(ctx context.Context, client *azblob.Client, blobName string, data []byte) error {
	// VULNERABLE: aucun controle d'acces supplementaire, le blob herite du conteneur public
	_, err := client.UploadBuffer(ctx, containerName, blobName, data, nil)
	if err != nil {
		return fmt.Errorf("upload blob: %w", err)
	}
	return nil
}
