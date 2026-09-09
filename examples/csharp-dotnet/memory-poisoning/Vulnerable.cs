// Vulnérable : Memory Poisoning (CWE-349)
// Le modèle extrait automatiquement des "faits" de la conversation et les
// écrit directement dans la mémoire persistante de l'agent, réinjectée
// telle quelle comme contexte de confiance lors des sessions futures, sans
// confirmation de l'utilisateur ni cloisonnement entre utilisateurs.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/assistant")]
public class AssistantController : ControllerBase
{
    private readonly ILlmClient _llm;
    private readonly IMemoryStore _memoryStore;

    public AssistantController(ILlmClient llm, IMemoryStore memoryStore)
    {
        _llm = llm;
        _memoryStore = memoryStore;
    }

    [HttpPost("chat")]
    public async Task<IActionResult> Chat([FromBody] AssistantRequest request)
    {
        // Mémoire partagée sans distinction claire de propriétaire,
        // réinjectée directement comme contexte de confiance.
        var memory = await _memoryStore.GetAsync(request.UserId);

        var messages = BuildMessages(memory, request);
        var response = await _llm.ChatAsync(messages);

        // Le modèle décide seul des faits à mémoriser et l'écriture est
        // effectuée immédiatement, sans validation de l'utilisateur.
        var extractedFacts = await _llm.ExtractFactsAsync(request.Message, response.Content);
        foreach (var fact in extractedFacts)
        {
            await _memoryStore.AppendAsync(request.UserId, fact);
        }

        return Ok(new { response.Content });
    }

    private static List<ChatMessage> BuildMessages(IEnumerable<string> memory, AssistantRequest request)
    {
        var messages = new List<ChatMessage>
        {
            new("system", "Contexte mémorisé : " + string.Join("; ", memory))
        };
        messages.Add(new("user", request.Message));
        return messages;
    }
}
