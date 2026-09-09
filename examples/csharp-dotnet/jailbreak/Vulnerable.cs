// Vulnérable : Jailbreak de modèle (CWE-1427)
// Le endpoint de chat transmet directement les messages de l'utilisateur
// au modèle avec pour seule barrière le system prompt. Il n'existe aucune
// couche de modération indépendante en entrée ou en sortie, ni suivi du
// nombre de tours suspects dans une session.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/chat")]
public class ChatController : ControllerBase
{
    private const string SystemPrompt = "Tu es un assistant utile et respectueux des règles de l'entreprise.";

    private readonly ILlmClient _llm;

    public ChatController(ILlmClient llm)
    {
        _llm = llm;
    }

    [HttpPost]
    public async Task<IActionResult> Chat([FromBody] ChatRequest request)
    {
        // Le system prompt est la seule barrière de sécurité. Aucun
        // classifieur indépendant ne vérifie l'entrée ni la sortie du
        // modèle, et rien ne limite le nombre de tours de conversation.
        var messages = new List<ChatMessage>
        {
            new("system", SystemPrompt)
        };
        messages.AddRange(request.History);
        messages.Add(new("user", request.Message));

        var response = await _llm.ChatAsync(messages);

        return Ok(new { response.Content });
    }
}
