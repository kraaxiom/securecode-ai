// Corrigé : Model Inversion (CWE-200)
// La granularité des sorties est restreinte, un bruit de confidentialité
// différentielle est appliqué, et les requêtes sont limitées par client
// pour empêcher la reconstruction progressive de données mémorisées.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/clinical-model")]
public class ClinicalModelController : ControllerBase
{
    private readonly IFineTunedModelClient _model;
    private readonly IQueryQuotaService _quotaService;
    private readonly IDifferentialPrivacyService _privacyService;

    public ClinicalModelController(
        IFineTunedModelClient model,
        IQueryQuotaService quotaService,
        IDifferentialPrivacyService privacyService)
    {
        _model = model;
        _quotaService = quotaService;
        _privacyService = privacyService;
    }

    [HttpPost("query")]
    public async Task<IActionResult> Query([FromBody] ClinicalQuery query)
    {
        var clientId = HttpContext.Items["ClientId"] as string
            ?? throw new InvalidOperationException("Client non identifié.");

        // Limitation et surveillance des volumes de requêtes pour empêcher
        // une reconstruction progressive de données mémorisées.
        if (!await _quotaService.TryConsumeAsync(clientId))
            return StatusCode(429, "Quota de requêtes dépassé.");

        var prediction = await _model.PredictAsync(query.Text);

        // Application d'un bruit de confidentialité différentielle et
        // restriction de la sortie à une catégorie agrégée, sans score brut
        // ni probabilité token par token.
        var noisyCategory = _privacyService.ApplyNoiseAndAggregate(prediction.PerClassConfidence);

        return Ok(new { Category = noisyCategory });
    }
}

public record ClinicalQuery(string Text);
