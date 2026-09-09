// Corrigé : Token Leakage (CWE-522)
// Les jetons d'authentification sont exclus du contexte transmis au modèle
// et systématiquement masqués dans les journaux et traces de télémétrie.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class PersonalizedAssistantController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly ILogger<PersonalizedAssistantController> _logger;
    private readonly ISensitiveHeaderRedactor _headerRedactor;

    public PersonalizedAssistantController(
        ILlmClient llmClient,
        ILogger<PersonalizedAssistantController> logger,
        ISensitiveHeaderRedactor headerRedactor)
    {
        _llmClient = llmClient;
        _logger = logger;
        _headerRedactor = headerRedactor;
    }

    [HttpPost("ask")]
    public async Task<IActionResult> Ask([FromBody] AskRequest request)
    {
        // La personnalisation se fait via une résolution côté serveur
        // (identité déjà établie par le middleware d'authentification),
        // jamais en transmettant le jeton brut au modèle.
        var userId = User.Identity?.Name
            ?? throw new UnauthorizedAccessException("Utilisateur non authentifié.");

        // Les en-têtes sensibles sont masqués avant toute journalisation.
        _logger.LogInformation("Requête reçue, en-têtes : {Headers}",
            _headerRedactor.Redact(Request.Headers));

        // Le prompt ne contient aucun jeton d'authentification, uniquement
        // les données métier nécessaires à la question posée.
        var prompt = $"Question de l'utilisateur {userId} : {request.Question}";
        var response = await _llmClient.CompleteAsync(prompt);

        return Ok(response.Text);
    }
}

public record AskRequest(string Question);
