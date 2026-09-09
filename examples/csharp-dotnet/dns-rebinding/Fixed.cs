using System.Net;
using System.Net.Http;
using System.Net.Sockets;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;

namespace SecurityAiSkill.Examples.DnsRebinding;

/// <summary>
/// Exemple CORRIGÉ — Contournement SSRF par DNS rebinding (CWE-918).
/// Applique le "DNS pinning" : l'IP validée est celle effectivement utilisée
/// pour la connexion TCP, via un SocketsHttpHandler.ConnectCallback personnalisé.
/// </summary>
public static class SafeHttpClientRegistration
{
    private static readonly HashSet<string> AllowedHosts = new(StringComparer.OrdinalIgnoreCase)
    {
        "api.partenaire.example.com",
    };

    /// <summary>
    /// Enregistre un HttpClient nommé dont le handler résout le DNS une seule
    /// fois, valide l'adresse IP obtenue, puis force la connexion TCP à se
    /// faire sur CETTE IP précise (pinning) — éliminant la fenêtre TOCTOU
    /// exploitée par le DNS rebinding, où une seconde résolution DNS pourrait
    /// renvoyer une adresse différente (interne) de celle validée.
    /// </summary>
    public static void AddDnsPinnedClient(this IServiceCollection services)
    {
        services.AddHttpClient("DnsPinnedClient", client =>
        {
            client.Timeout = TimeSpan.FromSeconds(5);
        }).ConfigurePrimaryHttpMessageHandler(() => new SocketsHttpHandler
        {
            AllowAutoRedirect = false,
            ConnectCallback = async (context, cancellationToken) =>
            {
                var host = context.DnsEndPoint.Host;

                if (!AllowedHosts.Contains(host))
                {
                    throw new InvalidOperationException("Hôte non autorisé.");
                }

                var addresses = await Dns.GetHostAddressesAsync(host, cancellationToken);
                var validated = addresses.FirstOrDefault(a => !IsPrivateOrReserved(a));

                if (validated is null)
                {
                    throw new InvalidOperationException("Aucune adresse IP valide pour cet hôte.");
                }

                // Connexion TCP effectuée directement sur l'IP validée à
                // l'instant T, sans nouvelle résolution DNS entre la
                // validation et la connexion réelle.
                var socket = new Socket(SocketType.Stream, ProtocolType.Tcp);
                try
                {
                    await socket.ConnectAsync(validated, context.DnsEndPoint.Port, cancellationToken);
                    return new NetworkStream(socket, ownsSocket: true);
                }
                catch
                {
                    socket.Dispose();
                    throw;
                }
            },
        });
    }

    private static bool IsPrivateOrReserved(IPAddress ip)
    {
        if (IPAddress.IsLoopback(ip)) return true;
        var bytes = ip.GetAddressBytes();
        if (ip.AddressFamily != AddressFamily.InterNetwork) return false;
        return bytes[0] == 10
            || (bytes[0] == 172 && bytes[1] is >= 16 and <= 31)
            || (bytes[0] == 192 && bytes[1] == 168)
            || (bytes[0] == 169 && bytes[1] == 254);
    }
}

[ApiController]
[Route("api/webhooks")]
public class WebhookValidatorController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public WebhookValidatorController(IHttpClientFactory httpClientFactory)
    {
        // Client dédié dont le ConnectCallback applique le DNS pinning
        // (voir SafeHttpClientRegistration.AddDnsPinnedClient).
        _httpClient = httpClientFactory.CreateClient("DnsPinnedClient");
    }

    [HttpPost("validate")]
    public async Task<IActionResult> ValidateUrl([FromQuery] string url)
    {
        if (!Uri.TryCreate(url, UriKind.Absolute, out var uri) || uri.Scheme != Uri.UriSchemeHttps)
        {
            return BadRequest("URL invalide.");
        }

        try
        {
            var response = await _httpClient.GetAsync(uri);
            var body = await response.Content.ReadAsStringAsync();
            return Ok(new { status = (int)response.StatusCode, body });
        }
        catch (InvalidOperationException ex)
        {
            return BadRequest(new { error = ex.Message });
        }
    }
}
