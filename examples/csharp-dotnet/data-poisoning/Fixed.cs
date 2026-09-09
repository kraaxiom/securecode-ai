// Corrigé : Data Poisoning (CWE-349)
// Chaque source de données est validée en provenance/intégrité, les
// retours utilisateurs sont isolés pour un échantillonnage humain avant
// réintégration, et les performances du modèle sont comparées à un jeu de
// référence fixe après chaque cycle d'entraînement.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/training")]
public class TrainingDataController : ControllerBase
{
    private readonly ITrainingDataStore _store;
    private readonly IQuarantineStore _quarantine;
    private readonly IDataProvenanceValidator _provenanceValidator;
    private readonly IAnomalyDetector _anomalyDetector;
    private readonly IFineTuningService _fineTuning;
    private readonly IGoldenSetEvaluator _goldenSetEvaluator;

    public TrainingDataController(
        ITrainingDataStore store,
        IQuarantineStore quarantine,
        IDataProvenanceValidator provenanceValidator,
        IAnomalyDetector anomalyDetector,
        IFineTuningService fineTuning,
        IGoldenSetEvaluator goldenSetEvaluator)
    {
        _store = store;
        _quarantine = quarantine;
        _provenanceValidator = provenanceValidator;
        _anomalyDetector = anomalyDetector;
        _fineTuning = fineTuning;
        _goldenSetEvaluator = goldenSetEvaluator;
    }

    [HttpPost("feedback")]
    public async Task<IActionResult> SubmitFeedback([FromBody] FeedbackExample example)
    {
        // Le retour utilisateur est mis en quarantaine, pas intégré directement.
        await _quarantine.AddAsync(example);
        return Accepted();
    }

    [HttpPost("feedback/review")]
    public async Task<IActionResult> ReviewQuarantine([FromBody] ReviewDecision decision)
    {
        // Validation humaine explicite avant réintégration au jeu d'entraînement.
        var example = await _quarantine.GetAsync(decision.ExampleId);
        if (decision.Approved && !_anomalyDetector.IsAnomalous(example))
        {
            await _store.AddAsync(example);
        }
        await _quarantine.RemoveAsync(decision.ExampleId);
        return Ok();
    }

    [HttpPost("finetune")]
    public async Task<IActionResult> RunFineTuning()
    {
        var dataset = await _store.GetAllAsync();

        // Provenance et intégrité vérifiées pour chaque source figée.
        foreach (var source in dataset.Sources)
        {
            if (!await _provenanceValidator.IsTrustedAsync(source))
                return BadRequest($"Source non vérifiée rejetée : {source.Id}");
        }

        var result = await _fineTuning.RunAsync(dataset);

        // Comparaison avec le jeu de référence pour détecter une dérive suspecte.
        var driftReport = await _goldenSetEvaluator.EvaluateAsync(result.ModelVersion);
        if (driftReport.HasSignificantDrift)
            return Conflict("Dérive de comportement détectée, modèle non déployé.");

        return Ok(result);
    }
}
