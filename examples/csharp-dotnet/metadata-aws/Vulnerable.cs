using System.Net.Http;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.MetadataAws;

/// <summary>
/// Exemple VULNÉRABLE — SSRF vers le service de métadonnées AWS/IMDS (CWE-918).
/// Application hébergée sur EC2 avec un rôle IAM attaché, exposant un proxy
/// générique de récupération d'images distantes.
/// </summary>
[ApiController]
[Route("api/images")]
public class RemoteImageController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public RemoteImageController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    // Vulnérable : aucune restriction sur la destination, et l'infrastructure
    // sous-jacente utilise potentiellement IMDSv1 (accessible via un simple
    // GET, sans jeton). Un attaquant peut fournir
    // http://169.254.169.254/latest/meta-data/iam/security-credentials/<role>
    // pour récupérer les identifiants temporaires du rôle IAM attaché à
    // l'instance EC2, et usurper ses permissions bien au-delà du périmètre
    // applicatif initial.
    [HttpGet("fetch")]
    public async Task<IActionResult> FetchImage([FromQuery] string source)
    {
        var response = await _httpClient.GetAsync(source);
        var bytes = await response.Content.ReadAsByteArrayAsync();
        return File(bytes, "application/octet-stream");
    }
}
