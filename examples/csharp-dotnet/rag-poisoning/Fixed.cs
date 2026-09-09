// Corrigé : RAG Poisoning (CWE-349)
// La source du document est validée contre une liste blanche de confiance,
// le contenu est analysé et noté avant indexation, et la provenance de
// chaque document est conservée pour permettre l'audit ultérieur.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/knowledge-base")]
public class IngestionController : ControllerBase
{
    private readonly IVectorStore _vectorStore;
    private readonly IEmbeddingService _embeddingService;
    private readonly ISourceAllowlist _sourceAllowlist;
    private readonly IContentTrustScorer _trustScorer;
    private readonly IAuditLog _auditLog;

    public IngestionController(
        IVectorStore vectorStore,
        IEmbeddingService embeddingService,
        ISourceAllowlist sourceAllowlist,
        IContentTrustScorer trustScorer,
        IAuditLog auditLog)
    {
        _vectorStore = vectorStore;
        _embeddingService = embeddingService;
        _sourceAllowlist = sourceAllowlist;
        _trustScorer = trustScorer;
        _auditLog = auditLog;
    }

    [HttpPost("ingest")]
    public async Task<IActionResult> Ingest([FromBody] IngestRequest request)
    {
        // La source doit figurer dans la liste blanche des sources autorisées.
        if (!_sourceAllowlist.IsTrusted(request.SourceUrl))
        {
            _auditLog.Record("rejected_untrusted_source", request.SourceUrl);
            return BadRequest("Source non autorisée.");
        }

        // Analyse du contenu et attribution d'un score de confiance avant
        // indexation, pour détecter des motifs anormaux (instructions
        // impératives, formatage suspect).
        var trustScore = _trustScorer.Evaluate(request.Content);
        if (trustScore.IsSuspicious)
        {
            _auditLog.Record("rejected_suspicious_content", request.SourceUrl, trustScore.Reason);
            return BadRequest("Contenu rejeté par l'analyse de confiance.");
        }

        var embedding = await _embeddingService.EmbedAsync(request.Content);

        await _vectorStore.UpsertAsync(new VectorDocument
        {
            Id = Guid.NewGuid().ToString(),
            Content = request.Content,
            Embedding = embedding,
            // Traçabilité de la provenance jusqu'à la réponse générée.
            SourceUrl = request.SourceUrl,
            TrustScore = trustScore.Value,
            IngestedAtUtc = DateTime.UtcNow
        });

        _auditLog.Record("document_ingested", request.SourceUrl, trustScore.Value);

        return Ok();
    }
}

public record IngestRequest(string Content, string SourceUrl);
