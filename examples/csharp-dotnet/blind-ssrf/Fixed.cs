using System.Net;
using System.Net.Http;
using System.Net.Sockets;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace SecurityAiSkill.Examples.BlindSsrf;

/// <summary>
/// Exemple CORRIGÉ — SSRF aveugle (CWE-918).
/// Applique exactement les mêmes contrôles qu'une SSRF classique (whitelist,
/// validation IP, timeout) au job asynchrone, et journalise chaque tentative
/// pour permettre la détection a posteriori d'un scan de réseau interne.
/// </summary>
[ApiController]
[Route("api/webhooks")]
public class WebhookController : ControllerBase
{
    private static readonly HashSet<string> AllowedHosts = new(StringComparer.OrdinalIgnoreCase)
    {
        "api.partenaire.example.com",
    };

    private readonly HttpClient _httpClient;
    private readonly ILogger<WebhookController> _logger;

    public WebhookController(IHttpClientFactory httpClientFactory, ILogger<WebhookController> logger)
    {
        _httpClient = httpClientFactory.CreateClient("SafeOutboundClient");
        _logger = logger;
    }

    public record RegisterWebhookRequest(string CallbackUrl);

    [HttpPost("register")]
    public IActionResult RegisterWebhook([FromBody] RegisterWebhookRequest request)
    {
        // Validation synchrone AVANT toute planification de tâche : même un
        // traitement "fire and forget" ne doit jamais contourner les contrôles
        // appliqués à une requête sortante synchrone.
        if (!TryValidateOutboundUrl(request.CallbackUrl, out var validatedUri, out var error))
        {
            return BadRequest(new { error });
        }

        _ = Task.Run(async () =>
        {
            using var cts = new CancellationTokenSource(TimeSpan.FromSeconds(5));
            try
            {
                // Journalisation systématique des requêtes sortantes déclenchées
                // par un traitement asynchrone, pour permettre une détection a
                // posteriori d'un scan interne même sans retour visible au client.
                _logger.LogInformation("Vérification webhook vers {Host}", validatedUri!.Host);
                await _httpClient.PostAsync(validatedUri, content: null, cts.Token);
            }
            catch (Exception ex)
            {
                _logger.LogWarning(ex, "Échec de vérification du webhook vers {Host}", validatedUri!.Host);
            }
        });

        return Accepted(new { message = "Webhook enregistré, vérification en cours." });
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

        if (addresses.Length == 0 || addresses.Any(IsPrivateOrReserved))
        {
            error = "Destination interdite.";
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
            return bytes[0] == 10
                || (bytes[0] == 172 && bytes[1] is >= 16 and <= 31)
                || (bytes[0] == 192 && bytes[1] == 168)
                || (bytes[0] == 169 && bytes[1] == 254);
        }

        return ip.IsIPv6LinkLocal || ip.IsIPv6SiteLocal;
    }
}
