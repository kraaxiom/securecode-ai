// Vulnérable : Embedding Poisoning (CWE-349)
// Tout contenu soumis par un utilisateur est directement transformé en
// embedding et indexé dans la base vectorielle, sans modération préalable
// ni limite de fréquence, permettant l'injection en volume de vecteurs
// conçus pour polluer la recherche sémantique.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/reviews")]
public class ReviewController : ControllerBase
{
    private readonly IEmbeddingClient _embeddingClient;
    private readonly IVectorStore _vectorStore;

    public ReviewController(IEmbeddingClient embeddingClient, IVectorStore vectorStore)
    {
        _embeddingClient = embeddingClient;
        _vectorStore = vectorStore;
    }

    [HttpPost]
    public async Task<IActionResult> SubmitReview([FromBody] ReviewSubmission review)
    {
        // Génération et indexation immédiates, sans modération ni contrôle
        // de volume par utilisateur ou par source.
        var vector = await _embeddingClient.EmbedAsync(review.Text);
        await _vectorStore.UpsertAsync(review.ProductId, vector, review.Text);

        return Ok();
    }
}
