// Corrigé : Conteneur Azure Blob Storage public (CWE-284)
// Le conteneur est créé en accès privé (aucun accès anonyme), et le partage
// ponctuel de blobs passe par un SAS en lecture seule, permissions minimales
// et expiration courte, plutôt que par un accès public permanent.
using Azure;
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
using Azure.Storage.Sas;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/storage")]
public class ContainerProvisioningController : ControllerBase
{
    private readonly BlobServiceClient _blobServiceClient;

    public ContainerProvisioningController(BlobServiceClient blobServiceClient)
    {
        _blobServiceClient = blobServiceClient;
    }

    [HttpPost("containers")]
    public async Task<IActionResult> CreateContainer([FromQuery] string containerName = "myfiles")
    {
        // Création du conteneur sans accès public anonyme (PublicAccessType.None).
        BlobContainerClient containerClient = await _blobServiceClient.CreateBlobContainerAsync(
            containerName,
            PublicAccessType.None);

        return Ok(new { containerName, message = "Conteneur créé en accès privé." });
    }

    [HttpGet("containers/{containerName}/blobs/{blobName}/share-link")]
    public IActionResult GetSasUrl(string containerName, string blobName)
    {
        var containerClient = _blobServiceClient.GetBlobContainerClient(containerName);
        var blobClient = containerClient.GetBlobClient(blobName);

        if (!blobClient.CanGenerateSasUri)
            return BadRequest("Impossible de générer un SAS avec ces identifiants.");

        // Partage ponctuel sécurisé via un SAS en lecture seule expirant dans 1 heure.
        var sasBuilder = new BlobSasBuilder
        {
            BlobContainerName = containerName,
            BlobName = blobName,
            Resource = "b",
            ExpiresOn = DateTimeOffset.UtcNow.AddHours(1),
            Protocol = SasProtocol.Https
        };
        sasBuilder.SetPermissions(BlobSasPermissions.Read);

        Uri sasUri = blobClient.GenerateSasUri(sasBuilder);
        return Ok(new { url = sasUri.ToString(), expiresInHours = 1 });
    }
}
