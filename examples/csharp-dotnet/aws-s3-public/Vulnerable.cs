// Vulnérable : Bucket S3 public (CWE-284)
// Le bucket est créé avec une ACL "public-read" et une bucket policy
// accordant un accès en lecture/liste à n'importe quel principal ("*"),
// exposant tous les objets sans authentification.
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
        // Création du bucket avec une ACL publique en lecture.
        await _s3Client.PutBucketAsync(new PutBucketRequest
        {
            BucketName = bucketName,
            CannedACL = S3CannedACL.PublicRead
        });

        // Bucket policy autorisant tout le monde ("*") à lire et lister les objets.
        var publicPolicy = $$"""
        {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Principal": "*",
            "Action": ["s3:GetObject", "s3:ListBucket"],
            "Resource": ["arn:aws:s3:::{{bucketName}}", "arn:aws:s3:::{{bucketName}}/*"]
          }]
        }
        """;

        await _s3Client.PutBucketPolicyAsync(new PutBucketPolicyRequest
        {
            BucketName = bucketName,
            Policy = publicPolicy
        });

        return Ok(new { bucketName, publicUrl = $"https://{bucketName}.s3.amazonaws.com/" });
    }
}
