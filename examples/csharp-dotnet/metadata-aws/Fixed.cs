using System.Net;
using System.Net.Http;
using System.Net.Sockets;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.MetadataAws;

/// <summary>
/// Exemple CORRIGÉ — SSRF vers le service de métadonnées AWS/IMDS (CWE-918).
/// Filtre explicitement les plages privées et link-local (dont
/// 169.254.169.254) ; à combiner côté infrastructure avec une migration vers
/// IMDSv2 (jeton de session requis) et un rôle IAM appliquant le moindre
/// privilège.
/// </summary>
[ApiController]
[Route("api/images")]
public class RemoteImageController : ControllerBase
{
    private static readonly HashSet<string> AllowedHosts = new(StringComparer.OrdinalIgnoreCase)
    {
        "cdn.example.com",
    };

    private readonly HttpClient _httpClient;

    public RemoteImageController(IHttpClientFactory httpClientFactory)
    {
        _httpClient = httpClientFactory.CreateClient("SafeOutboundClient");
    }

    [HttpGet("fetch")]
    public async Task<IActionResult> FetchImage([FromQuery] string source)
    {
        if (!TryValidateOutboundUrl(source, out var validatedUri, out var error))
        {
            return BadRequest(new { error });
        }

        using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(5));
        var response = await _httpClient.GetAsync(validatedUri, cts.Token);
        var bytes = await response.Content.ReadAsByteArrayAsync(cts.Token);
        return File(bytes, "application/octet-stream");
    }

    private static bool TryValidateOutboundUrl(string? rawUrl, out Uri? validatedUri, out string? error)
    {
        validatedUri = null;
        error = null;

        if (string.IsNullOrWhiteSpace(rawUrl) || !Uri.TryCreate(rawUrl, UriKind.Absolute, out var uri))
        {
            error = "URL invalide.";
            return false;
        }

        if (uri.Scheme != Uri.UriSchemeHttps || !AllowedHosts.Contains(uri.Host))
        {
            error = "URL non autorisée.";
            return false;
        }

        IPAddress[] addresses;
        try
        {
            addresses = Dns.GetHostAddresses(uri.Host);
        }
        catch (SocketException)
        {
            error = "Résolution DNS impossible.";
            return false;
        }

        // Rejet explicite de la plage link-local 169.254.0.0/16, qui héberge
        // le service de métadonnées AWS (169.254.169.254), et des autres
        // plages privées/réservées.
        if (addresses.Length == 0 || addresses.Any(IsPrivateOrReserved))
        {
            error = "Destination interdite (adresse interne/réservée, y compris IMDS).";
            return false;
        }

        validatedUri = uri;
        return true;
    }

    private static bool IsPrivateOrReserved(IPAddress ip)
    {
        if (IPAddress.IsLoopback(ip)) return true;
        var bytes = ip.GetAddressBytes();
        if (ip.AddressFamily != AddressFamily.InterNetwork) return false;
        return bytes[0] == 10
            || (bytes[0] == 172 && bytes[1] is >= 16 and <= 31)
            || (bytes[0] == 192 && bytes[1] == 168)
            || (bytes[0] == 169 && bytes[1] == 254); // couvre 169.254.169.254 (IMDS)
    }
}

// Rappel infrastructure (hors périmètre de ce fichier applicatif) :
// - Migrer vers IMDSv2 (jeton de session requis via requête PUT), ce qui
//   neutralise la plupart des SSRF classiques basées sur de simples GET.
// - Appliquer le principe du moindre privilège sur le rôle IAM attaché
//   à l'instance/tâche.
