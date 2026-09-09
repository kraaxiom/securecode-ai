using System.Net;
using System.Net.Http;
using System.Net.Sockets;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.SsrfClassique;

/// <summary>
/// Exemple CORRIGÉ — SSRF classique (CWE-918).
/// Applique une whitelist stricte d'hôtes, valide le schéma et l'adresse IP
/// résolue avant d'effectuer la requête sortante, et désactive les redirections
/// automatiques non revalidées.
/// </summary>
[ApiController]
[Route("api/preview")]
public class UrlPreviewController : ControllerBase
{
    // Whitelist explicite des hôtes métier autorisés pour l'aperçu d'URL.
    // Toute autre destination est refusée par défaut (deny by default).
    private static readonly HashSet<string> AllowedHosts = new(StringComparer.OrdinalIgnoreCase)
    {
        "api.partenaire.example.com",
        "cdn.example.com",
    };

    private readonly HttpClient _httpClient;

    public UrlPreviewController(IHttpClientFactory httpClientFactory)
    {
        // Client HTTP dédié, configuré pour ne jamais suivre automatiquement
        // les redirections : chaque redirection doit être revalidée manuellement.
        _httpClient = httpClientFactory.CreateClient("SafeOutboundClient");
    }

    [HttpGet]
    public async Task<IActionResult> GetPreview([FromQuery] string url)
    {
        if (!TryValidateOutboundUrl(url, out var validatedUri, out var error))
        {
            return BadRequest(new { error });
        }

        using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(5));
        try
        {
            var response = await _httpClient.GetAsync(validatedUri, cts.Token);

            // Une redirection doit être revalidée explicitement plutôt que suivie
            // automatiquement, sinon la whitelist peut être contournée.
            if (IsRedirect(response.StatusCode))
            {
                return BadRequest(new { error = "Les redirections ne sont pas suivies automatiquement." });
            }

            var content = await response.Content.ReadAsStringAsync(cts.Token);
            return Ok(new { status = (int)response.StatusCode, body = content });
        }
        catch (TaskCanceledException)
        {
            return StatusCode(504, new { error = "Délai dépassé lors de la requête sortante." });
        }
    }

    /// <summary>
    /// Valide le schéma, l'hôte (whitelist) et l'adresse IP résolue avant
    /// d'autoriser une requête sortante. Rejette explicitement les plages
    /// privées, loopback, link-local et réservées (défense en profondeur
    /// même si l'hôte figure dans la whitelist, pour couvrir les cas
    /// d'entrées DNS internes mal configurées).
    /// </summary>
    private static bool TryValidateOutboundUrl(string? rawUrl, out Uri? validatedUri, out string? error)
    {
        validatedUri = null;
        error = null;

        if (string.IsNullOrWhiteSpace(rawUrl) || !Uri.TryCreate(rawUrl, UriKind.Absolute, out var uri))
        {
            error = "URL invalide.";
            return false;
        }

        if (uri.Scheme != Uri.UriSchemeHttps)
        {
            error = "Seul le schéma HTTPS est autorisé.";
            return false;
        }

        if (!AllowedHosts.Contains(uri.Host))
        {
            error = "Hôte non autorisé.";
            return false;
        }

        // Résolution DNS explicite puis validation de l'adresse IP obtenue,
        // pour éviter qu'un hôte whitelisté ne pointe (ou soit reconfiguré
        // pour pointer, cf. DNS rebinding) vers une adresse interne.
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

        if (addresses.Length == 0 || addresses.Any(IsPrivateOrReserved))
        {
            error = "Destination interdite (adresse interne/réservée).";
            return false;
        }

        validatedUri = uri;
        return true;
    }

    private static bool IsPrivateOrReserved(IPAddress ip)
    {
        if (IPAddress.IsLoopback(ip))
        {
            return true;
        }

        var bytes = ip.GetAddressBytes();
        if (ip.AddressFamily == AddressFamily.InterNetwork)
        {
            // 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, 169.254.0.0/16 (link-local, y compris IMDS)
            return bytes[0] == 10
                || (bytes[0] == 172 && bytes[1] is >= 16 and <= 31)
                || (bytes[0] == 192 && bytes[1] == 168)
                || (bytes[0] == 169 && bytes[1] == 254);
        }

        // IPv6 : loopback (::1), link-local (fe80::/10), ULA (fc00::/7)
        return ip.IsIPv6LinkLocal || ip.IsIPv6SiteLocal || (bytes[0] & 0xfe) == 0xfc;
    }

    private static bool IsRedirect(HttpStatusCode code) =>
        code is HttpStatusCode.Moved or HttpStatusCode.Redirect or HttpStatusCode.SeeOther
            or HttpStatusCode.TemporaryRedirect or HttpStatusCode.PermanentRedirect;
}
