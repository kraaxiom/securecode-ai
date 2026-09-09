// Vulnérable : Model Extraction (CWE-200)
// L'endpoint d'inférence n'applique aucune limite de débit ni quota par client,
// et renvoie les logits/probabilités bruts du modèle, ce qui facilite
// l'interrogation massive nécessaire à la reconstitution d'un modèle équivalent.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/inference")]
public class InferenceController : ControllerBase
{
    private readonly IModelClient _modelClient;

    public InferenceController(IModelClient modelClient)
    {
        _modelClient = modelClient;
    }

    [HttpPost("predict")]
    public async Task<IActionResult> Predict([FromBody] PredictRequest request)
    {
        // Aucune vérification de quota, aucune limitation de débit par clé API.
        // N'importe quel client peut envoyer un volume illimité de requêtes.
        var result = await _modelClient.RunInferenceAsync(request.Input);

        // Les logits complets et les scores de confiance bruts sont renvoyés,
        // au-delà de la simple sortie textuelle nécessaire au cas d'usage.
        return Ok(new
        {
            Text = result.Text,
            Logits = result.RawLogits,
            ConfidenceScores = result.RawProbabilities
        });
    }
}

public record PredictRequest(string Input);
