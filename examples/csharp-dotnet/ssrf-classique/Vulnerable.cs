using System.Net.Http;
using Microsoft.AspNetCore.Mvc;

namespace SecurityAiSkill.Examples.SsrfClassique;

/// <summary>
/// Exemple VULNÉRABLE — SSRF classique (CWE-918).
/// Endpoint "aperçu d'URL" qui récupère le contenu d'une URL fournie par le client.
/// </summary>
[ApiController]
[Route("api/preview")]
public class UrlPreviewController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public UrlPreviewController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    // Vulnérable : l'URL provient intégralement de l'utilisateur et est utilisée
    // telle quelle pour effectuer une requête sortante depuis le serveur.
    // Aucune restriction de schéma, d'hôte ou d'adresse IP n'est appliquée.
    // Un attaquant peut fournir http://169.254.169.254/, http://localhost:6379/,
    // ou toute adresse du réseau interne pour faire du serveur un proxy vers
    // des ressources normalement inaccessibles depuis Internet.
    [HttpGet]
    public async Task<IActionResult> GetPreview([FromQuery] string url)
    {
        // Aucune validation : le client HTTP suit aussi les redirections par défaut,
        // ce qui permet de contourner un filtrage superficiel effectué en amont.
        var response = await _httpClient.GetAsync(url);
        var content = await response.Content.ReadAsStringAsync();

        return Ok(new { status = (int)response.StatusCode, body = content });
    }
}
