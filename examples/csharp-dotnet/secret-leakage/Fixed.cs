// Corrigé : Secret Leakage via LLM (CWE-200)
// Le secret n'est jamais inclus dans le prompt : l'agent utilise une
// référence indirecte résolue côté application uniquement, et un filtrage
// de sortie masque tout motif de secret avant renvoi ou journalisation.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class AssistantController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly ILogger<AssistantController> _logger;
    private readonly ISecretResolver _secretResolver;
    private readonly IOutputSecretScanner _outputScanner;

    // Référence indirecte : le modèle ne voit jamais la valeur réelle du secret.
    private const string SystemPromptTemplate =
        "Tu es un assistant interne. Pour appeler l'API de facturation, utilise l'outil " +
        "'billing_api_call' ; l'authentification est gérée automatiquement côté serveur.";

    public AssistantController(
        ILlmClient llmClient,
        ILogger<AssistantController> logger,
        ISecretResolver secretResolver,
        IOutputSecretScanner outputScanner)
    {
        _llmClient = llmClient;
        _logger = logger;
        _secretResolver = secretResolver;
        _outputScanner = outputScanner;
    }

    [HttpPost("ask")]
    public async Task<IActionResult> Ask([FromBody] AskRequest request)
    {
        var response = await _llmClient.CompleteAsync(SystemPromptTemplate + "\n" + request.Question);

        // Filtrage de sortie détectant les motifs de secrets avant tout
        // renvoi ou journalisation de la réponse du modèle.
        var scanResult = _outputScanner.ScanAndRedact(response.Text);
        if (scanResult.SecretsDetected)
        {
            _logger.LogWarning("Secret potentiel détecté et expurgé de la réponse du modèle.");
        }

        // Les journaux ne conservent jamais le texte brut susceptible de
        // contenir un secret : uniquement la version expurgée.
        _logger.LogInformation("Question: {Question} / Réponse (expurgée): {Response}",
            request.Question, scanResult.RedactedText);

        return Ok(scanResult.RedactedText);
    }

    // L'outil 'billing_api_call' résoudrait la vraie clé côté serveur, hors
    // du contexte transmis au modèle — non représenté ici pour rester concis.
}

public record AskRequest(string Question);
