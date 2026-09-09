// Corrigé : Model Extraction (CWE-200)
// Une limitation de débit stricte par clé API est appliquée, seule la sortie
// textuelle nécessaire au cas d'usage est renvoyée, et les volumes de requêtes
// sont journalisés pour détecter un pattern d'extraction anormal.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/inference")]
public class InferenceController : ControllerBase
{
    private readonly IModelClient _modelClient;
    private readonly IQueryQuotaService _quotaService;
    private readonly IUsageAnomalyLogger _usageLogger;

    public InferenceController(
        IModelClient modelClient,
        IQueryQuotaService quotaService,
        IUsageAnomalyLogger usageLogger)
    {
        _modelClient = modelClient;
        _quotaService = quotaService;
        _usageLogger = usageLogger;
    }

    [HttpPost("predict")]
    public async Task<IActionResult> Predict([FromBody] PredictRequest request)
    {
        var apiKeyId = HttpContext.Items["ApiKeyId"] as string
            ?? throw new InvalidOperationException("Clé API manquante.");

        // Quota et limitation de débit stricts par client/clé API.
        if (!await _quotaService.TryConsumeAsync(apiKeyId))
            return StatusCode(429, "Quota de requêtes dépassé.");

        // Surveillance des patterns de requêtes pour détecter une extraction
        // systématique (volume, régularité, diversité anormale).
        await _usageLogger.RecordQueryAsync(apiKeyId, request.Input);

        var result = await _modelClient.RunInferenceAsync(request.Input);

        // Seule la sortie strictement nécessaire au cas d'usage métier est
        // renvoyée : ni logits, ni probabilités brutes.
        return Ok(new { Text = result.Text });
    }
}

public record PredictRequest(string Input);
