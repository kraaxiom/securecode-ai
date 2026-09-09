// Vulnérable : RAG Poisoning (CWE-349)
// Tout document soumis est indexé directement dans la base vectorielle,
// sans validation de provenance, sans liste blanche de sources de confiance
// et sans analyse de contenu, ce qui permet d'empoisonner le contexte
// utilisé ultérieurement par le modèle.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/knowledge-base")]
public class IngestionController : ControllerBase
{
    private readonly IVectorStore _vectorStore;
    private readonly IEmbeddingService _embeddingService;

    public IngestionController(IVectorStore vectorStore, IEmbeddingService embeddingService)
    {
        _vectorStore = vectorStore;
        _embeddingService = embeddingService;
    }

    [HttpPost("ingest")]
    public async Task<IActionResult> Ingest([FromBody] IngestRequest request)
    {
        // Aucune vérification de la source du document, aucune liste blanche,
        // aucune analyse de contenu avant indexation.
        var embedding = await _embeddingService.EmbedAsync(request.Content);

        await _vectorStore.UpsertAsync(new VectorDocument
        {
            Id = Guid.NewGuid().ToString(),
            Content = request.Content,
            Embedding = embedding
            // Pas de champ de provenance : impossible d'auditer plus tard
            // d'où provient un chunk récupéré par le pipeline RAG.
        });

        return Ok();
    }
}

public record IngestRequest(string Content, string SourceUrl);
