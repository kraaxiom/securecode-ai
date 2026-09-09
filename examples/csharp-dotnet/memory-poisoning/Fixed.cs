// Corrigé : Memory Poisoning (CWE-349)
// Toute écriture durable en mémoire exige une confirmation explicite de
// l'utilisateur, la mémoire est strictement cloisonnée par utilisateur, et
// le contenu réinjecté est traité comme une donnée à revalider plutôt que
// comme une instruction de confiance absolue.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class AssistantController : ControllerBase
{
    private readonly ILlmClient _llm;
    private readonly IMemoryStore _memoryStore;
    private readonly IAuditLog _auditLog;

    public AssistantController(ILlmClient llm, IMemoryStore memoryStore, IAuditLog auditLog)
    {
        _llm = llm;
        _memoryStore = memoryStore;
        _auditLog = auditLog;
    }

    [HttpPost("chat")]
    public async Task<IActionResult> Chat([FromBody] AssistantRequest request)
    {
        // Mémoire strictement cloisonnée par utilisateur, marquée comme
        // donnée à revalider plutôt que comme instruction de confiance.
        var memory = await _memoryStore.GetForUserAsync(request.UserId);

        var messages = BuildMessages(memory, request);
        var response = await _llm.ChatAsync(messages);

        var candidateFacts = await _llm.ExtractFactsAsync(request.Message, response.Content);

        // Les faits proposés sont mis en attente de confirmation explicite
        // de l'utilisateur avant toute écriture durable.
        foreach (var fact in candidateFacts)
        {
            await _memoryStore.QueuePendingAsync(request.UserId, fact);
        }

        return Ok(new
        {
            response.Content,
            PendingMemoryUpdates = candidateFacts.Select(f => f.Summary)
        });
    }

    [HttpPost("memory/confirm")]
    public async Task<IActionResult> ConfirmMemory([FromBody] MemoryConfirmation confirmation)
    {
        // Écriture durable uniquement après confirmation explicite de
        // l'utilisateur concerné.
        if (confirmation.Approved)
        {
            await _memoryStore.CommitAsync(confirmation.UserId, confirmation.FactId);
            await _auditLog.RecordAsync("memory_fact_committed", confirmation);
        }
        else
        {
            await _memoryStore.DiscardPendingAsync(confirmation.UserId, confirmation.FactId);
        }

        return Ok();
    }

    [HttpGet("memory")]
    public async Task<IActionResult> ListMemory([FromQuery] string userId)
    {
        // L'utilisateur peut consulter, corriger et supprimer ses entrées mémoire.
        var memory = await _memoryStore.GetForUserAsync(userId);
        return Ok(memory);
    }

    [HttpDelete("memory/{factId}")]
    public async Task<IActionResult> DeleteMemory(string factId, [FromQuery] string userId)
    {
        await _memoryStore.RemoveAsync(userId, factId);
        await _auditLog.RecordAsync("memory_fact_deleted", new { userId, factId });
        return NoContent();
    }

    private static List<ChatMessage> BuildMessages(IEnumerable<string> memory, AssistantRequest request)
    {
        var messages = new List<ChatMessage>
        {
            // Le contenu mémorisé est présenté comme information contextuelle
            // à revalider, pas comme une instruction impérative.
            new("system", "Informations précédemment mémorisées (à valider si utilisées) : " + string.Join("; ", memory))
        };
        messages.Add(new("user", request.Message));
        return messages;
    }
}
