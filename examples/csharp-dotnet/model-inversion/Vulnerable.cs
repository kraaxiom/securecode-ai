// Vulnérable : Model Inversion (CWE-200)
// Le modèle fine-tuné sur des données internes sensibles renvoie des scores
// de confiance bruts et des sorties verbeuses, sans limite de requêtes,
// ce qui permet une reconstruction progressive de données mémorisées.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/clinical-model")]
public class ClinicalModelController : ControllerBase
{
    private readonly IFineTunedModelClient _model;

    public ClinicalModelController(IFineTunedModelClient model)
    {
        _model = model;
    }

    [HttpPost("query")]
    public async Task<IActionResult> Query([FromBody] ClinicalQuery query)
    {
        // Aucune limite sur le volume ou la structure des requêtes :
        // un attaquant peut interroger le modèle de façon répétée et méthodique.
        var prediction = await _model.PredictAsync(query.Text);

        // Renvoi des scores de confiance bruts par classe/token, ce qui
        // facilite l'inférence d'appartenance et la reconstruction de
        // données d'entraînement mémorisées.
        return Ok(new
        {
            prediction.OutputText,
            RawClassScores = prediction.PerClassConfidence,
            TokenLevelProbabilities = prediction.TokenLogProbs
        });
    }
}

public record ClinicalQuery(string Text);
