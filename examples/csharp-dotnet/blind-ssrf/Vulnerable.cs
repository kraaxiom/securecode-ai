using System.Net.Http;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.BlindSsrf;

/// <summary>
/// Exemple VULNÉRABLE — SSRF aveugle (Blind SSRF, CWE-918).
/// Endpoint d'enregistrement de webhook : l'URL est "vérifiée" en arrière-plan
/// par un job qui n'expose jamais son résultat au client.
/// </summary>
[ApiController]
[Route("api/webhooks")]
public class WebhookController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public WebhookController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    public record RegisterWebhookRequest(string CallbackUrl);

    // Vulnérable : le serveur planifie une requête HTTP "fire and forget" vers
    // l'URL fournie par le client, sans jamais renvoyer le contenu récupéré.
    // Le développeur pense (à tort) que l'absence de retour visible réduit le
    // risque : en réalité, la requête sortante elle-même reste dangereuse
    // (scan du réseau interne, atteinte de services internes) même sans canal
    // de lecture direct — l'exploitation se fait via le délai de réponse ou
    // un callback hors bande (DNS/HTTP) contrôlé par l'attaquant.
    [HttpPost("register")]
    public IActionResult RegisterWebhook([FromBody] RegisterWebhookRequest request)
    {
        // Job asynchrone "fire and forget" sans aucune validation de destination.
        _ = Task.Run(async () =>
        {
            try
            {
                await _httpClient.PostAsync(request.CallbackUrl, content: null);
            }
            catch
            {
                // L'erreur n'est jamais journalisée ni remontée : aucune trace
                // exploitable pour détecter un scan interne a posteriori.
            }
        });

        return Accepted(new { message = "Webhook enregistré, vérification en cours." });
    }
}
