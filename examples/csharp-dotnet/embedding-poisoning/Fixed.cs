// Corrigé : Embedding Poisoning (CWE-349)
// Le contenu est modéré avant génération d'embedding, le volume d'insertion
// par utilisateur/source est limité dans le temps, et chaque insertion est
// journalisée pour permettre une surveillance de la distribution des vecteurs.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/reviews")]
public class ReviewController : ControllerBase
{
    private readonly IEmbeddingClient _embeddingClient;
    private readonly IVectorStore _vectorStore;
    private readonly IContentModerationService _moderation;
    private readonly IRateLimiter _rateLimiter;
    private readonly IAuditLog _auditLog;

    public ReviewController(
        IEmbeddingClient embeddingClient,
        IVectorStore vectorStore,
        IContentModerationService moderation,
        IRateLimiter rateLimiter,
        IAuditLog auditLog)
    {
        _embeddingClient = embeddingClient;
        _vectorStore = vectorStore;
        _moderation = moderation;
        _rateLimiter = rateLimiter;
        _auditLog = auditLog;
    }

    [HttpPost]
    public async Task<IActionResult> SubmitReview([FromBody] ReviewSubmission review)
    {
        // Limite de fréquence/volume par utilisateur pour freiner les
        // campagnes d'empoisonnement automatisées.
        if (!await _rateLimiter.AllowAsync(review.UserId, "review_embedding"))
            return StatusCode(429, "Trop de soumissions, réessayez plus tard.");

        // Modération du contenu avant génération et indexation de l'embedding.
        var moderationResult = await _moderation.CheckAsync(review.Text);
        if (!moderationResult.IsAllowed)
        {
            await _auditLog.RecordAsync("rejected_review_content", review);
            return BadRequest("Contenu rejeté par la modération.");
        }

        var vector = await _embeddingClient.EmbedAsync(review.Text);
        await _vectorStore.UpsertAsync(review.ProductId, vector, review.Text);
        await _auditLog.RecordAsync("embedding_indexed", review);

        return Ok();
    }
}
