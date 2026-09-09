using System.Net;
using System.Net.Http;
using System.Net.Sockets;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.DnsRebinding;

/// <summary>
/// Exemple VULNÉRABLE — Contournement SSRF par DNS rebinding (CWE-918).
/// La validation semble correcte (résolution DNS + rejet des IP privées) mais
/// reste vulnérable au TOCTOU : l'IP validée n'est pas celle réellement utilisée
/// par la requête HTTP.
/// </summary>
[ApiController]
[Route("api/webhooks")]
public class WebhookValidatorController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public WebhookValidatorController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    // Vulnérable : la résolution DNS effectuée ici pour valider l'hôte est
    // totalement indépendante de celle que HttpClient effectuera lui-même au
    // moment de la connexion réelle. Un attaquant contrôlant le domaine peut
    // répondre avec une IP publique légitime lors de cette première résolution
    // (TTL très court), puis avec une IP interne (127.0.0.1, 169.254.169.254,
    // adresse RFC1918...) lors de la résolution réellement utilisée par le
    // client HTTP quelques instants plus tard : la validation est ainsi
    // totalement contournée (Time-Of-Check / Time-Of-Use).
    [HttpPost("validate")]
    public async Task<IActionResult> ValidateUrl([FromQuery] string url)
    {
        if (!Uri.TryCreate(url, UriKind.Absolute, out var uri) || uri.Scheme != Uri.UriSchemeHttps)
        {
            return BadRequest("URL invalide.");
        }

        var addresses = await Dns.GetHostAddressesAsync(uri.Host);
        if (addresses.Any(IsPrivateOrReserved))
        {
            return BadRequest("Destination interdite.");
        }

        // Entre cette validation et l'appel ci-dessous, HttpClient effectue sa
        // PROPRE résolution DNS du même nom d'hôte : rien ne garantit qu'elle
        // renvoie la même adresse IP que celle qui vient d'être validée.
        var response = await _httpClient.GetAsync(uri);
        var body = await response.Content.ReadAsStringAsync();
        return Ok(new { status = (int)response.StatusCode, body });
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
