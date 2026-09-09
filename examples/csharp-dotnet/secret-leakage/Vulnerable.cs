// Vulnérable : Secret Leakage via LLM (CWE-200)
// Le prompt système inclut en clair une clé API interne pour que l'agent
// puisse l'utiliser, et les réponses du modèle sont renvoyées et journalisées
// sans aucun filtrage, ce qui peut exposer ce secret.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class AssistantController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly ILogger<AssistantController> _logger;

    // Secret inclus directement dans le template de prompt versionné.
    private const string InternalApiKey = "sk-internal-PLACEHOLDER-0000000000";

    private static readonly string SystemPromptTemplate =
        $"Tu es un assistant interne. Utilise cette clé pour appeler l'API de facturation : {InternalApiKey}";

    public AssistantController(ILlmClient llmClient, ILogger<AssistantController> logger)
    {
        _llmClient = llmClient;
        _logger = logger;
    }

    [HttpPost("ask")]
    public async Task<IActionResult> Ask([FromBody] AskRequest request)
    {
        var response = await _llmClient.CompleteAsync(SystemPromptTemplate + "\n" + request.Question);

        // Le prompt et la réponse complète sont journalisés en clair,
        // sans masquage d'un éventuel secret restitué par le modèle.
        _logger.LogInformation("Prompt: {Prompt} / Réponse: {Response}", SystemPromptTemplate, response.Text);

        return Ok(response.Text);
    }
}

public record AskRequest(string Question);
