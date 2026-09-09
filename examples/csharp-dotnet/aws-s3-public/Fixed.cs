// Corrigé : Bucket S3 public (CWE-284)
// Le bucket est créé sans ACL publique, "Block Public Access" est activé
// intégralement, et le partage ponctuel d'objets passe par une URL
// pré-signée à expiration courte plutôt que par un accès public permanent.
using Amazon.S3;
using Amazon.S3.Model;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/storage")]
public class BucketProvisioningController : ControllerBase
{
    private readonly IAmazonS3 _s3Client;

    public BucketProvisioningController(IAmazonS3 s3Client)
    {
        _s3Client = s3Client;
    }

    [HttpPost("buckets")]
    public async Task<IActionResult> CreateBucket([FromQuery] string bucketName = "my-app-bucket")
    {
        // Création du bucket sans ACL publique (ACL par défaut privée).
        await _s3Client.PutBucketAsync(new PutBucketRequest
        {
            BucketName = bucketName,
            CannedACL = S3CannedACL.Private
        });

        // Activation intégrale de "Block Public Access" au niveau du bucket.
        await _s3Client.PutPublicAccessBlockAsync(new PutPublicAccessBlockRequest
        {
            BucketName = bucketName,
            PublicAccessBlockConfiguration = new PublicAccessBlockConfiguration
            {
                BlockPublicAcls = true,
                IgnorePublicAcls = true,
                BlockPublicPolicy = true,
                RestrictPublicBuckets = true
            }
        });

        // Chiffrement au repos activé par défaut sur le bucket.
        await _s3Client.PutBucketEncryptionAsync(new PutBucketEncryptionRequest
        {
            BucketName = bucketName,
            ServerSideEncryptionConfiguration = new ServerSideEncryptionConfiguration
            {
                ServerSideEncryptionRules = new List<ServerSideEncryptionRule>
                {
                    new ServerSideEncryptionRule
                    {
                        ServerSideEncryptionByDefault = new ServerSideEncryptionByDefault
                        {
                            ServerSideEncryptionAlgorithm = ServerSideEncryptionMethod.AWSKMS
                        }
                    }
                }
            }
        });

        return Ok(new { bucketName, message = "Bucket créé en accès privé." });
    }

    [HttpGet("buckets/{bucketName}/objects/{key}/share-link")]
    public IActionResult GetPresignedUrl(string bucketName, string key)
    {
        // Partage ponctuel sécurisé via une URL pré-signée expirant dans 15 minutes.
        var request = new GetPreSignedUrlRequest
        {
            BucketName = bucketName,
            Key = key,
            Verb = HttpVerb.GET,
            Expires = DateTime.UtcNow.AddMinutes(15)
        };

        string url = _s3Client.GetPreSignedURL(request);
        return Ok(new { url, expiresInMinutes = 15 });
    }
}
