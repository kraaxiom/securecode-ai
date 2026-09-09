// Vulnérable : Conteneur Azure Blob Storage public (CWE-284)
// Le conteneur est créé avec un niveau d'accès public "Container", permettant
// à quiconque de lister et lire les blobs sans authentification.
using Azure.Storage.Blobs;
using Azure.Storage.Blobs.Models;
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
        // Création du conteneur avec accès public en lecture au niveau conteneur
        // (liste + lecture des blobs accessibles anonymement).
        BlobContainerClient containerClient = await _blobServiceClient.CreateBlobContainerAsync(
            containerName,
            PublicAccessType.BlobContainer);

        var url = containerClient.Uri.ToString();
        return Ok(new { containerName, publicUrl = url });
    }
}
