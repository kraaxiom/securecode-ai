// Vulnérable : Token Leakage (CWE-522)
// Le jeton de session de l'utilisateur est inclus dans le contexte transmis
// au modèle, et l'en-tête d'autorisation complet est journalisé en clair
// à chaque requête, exposant les jetons d'authentification.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class PersonalizedAssistantController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly ILogger<PersonalizedAssistantController> _logger;

    public PersonalizedAssistantController(ILlmClient llmClient, ILogger<PersonalizedAssistantController> logger)
    {
        _llmClient = llmClient;
        _logger = logger;
    }

    [HttpPost("ask")]
    public async Task<IActionResult> Ask([FromBody] AskRequest request)
    {
        var sessionToken = Request.Headers["Authorization"].ToString();

        // Journalisation de l'en-tête d'autorisation complet, sans masquage.
        _logger.LogInformation("Requête reçue avec en-tête complet : {Headers}", Request.Headers);

        // Le jeton de session est inclus dans le contexte transmis au modèle
        // pour "personnaliser" la réponse : il devient partie du prompt.
        var prompt = $"Jeton de session utilisateur : {sessionToken}\nQuestion : {request.Question}";
        var response = await _llmClient.CompleteAsync(prompt);

        return Ok(response.Text);
    }
}

public record AskRequest(string Question);
