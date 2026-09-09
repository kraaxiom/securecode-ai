// Vulnérable : Data Poisoning (CWE-349)
// Les exemples issus des retours utilisateurs sont ajoutés directement au
// jeu de données de fine-tuning, sans vérification de provenance, sans
// détection d'anomalies statistiques et sans comparaison avec un jeu de
// référence après entraînement.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/training")]
public class TrainingDataController : ControllerBase
{
    private readonly ITrainingDataStore _store;
    private readonly IFineTuningService _fineTuning;

    public TrainingDataController(ITrainingDataStore store, IFineTuningService fineTuning)
    {
        _store = store;
        _fineTuning = fineTuning;
    }

    [HttpPost("feedback")]
    public async Task<IActionResult> SubmitFeedback([FromBody] FeedbackExample example)
    {
        // Le retour utilisateur est directement ajouté au jeu de données,
        // sans validation d'origine ni détection d'anomalie.
        await _store.AddAsync(example);
        return Ok();
    }

    [HttpPost("finetune")]
    public async Task<IActionResult> RunFineTuning()
    {
        var dataset = await _store.GetAllAsync();
        // Aucune comparaison avec un jeu de référence (golden set) pour
        // détecter une dérive de comportement introduite par le dataset.
        var result = await _fineTuning.RunAsync(dataset);
        return Ok(result);
    }
}
