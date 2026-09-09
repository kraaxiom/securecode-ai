// Vulnérable : Prompt Injection (CWE-1427)
// L'entrée utilisateur est concaténée directement au prompt système sous
// forme de texte brut, sans séparation structurelle des rôles, et tout appel
// d'outil renvoyé par le modèle est exécuté sans validation applicative.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/chat")]
public class ChatController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly IToolRegistry _toolRegistry;

    private const string SystemPrompt =
        "Tu es un assistant interne. Réponds aux questions des employés.";

    public ChatController(ILlmClient llmClient, IToolRegistry toolRegistry)
    {
        _llmClient = llmClient;
        _toolRegistry = toolRegistry;
    }

    [HttpPost("message")]
    public async Task<IActionResult> SendMessage([FromBody] ChatRequest request)
    {
        // Concaténation brute : le modèle ne peut pas distinguer
        // structurellement l'instruction développeur de l'entrée utilisateur.
        var fullPrompt = SystemPrompt + "\n" + request.Message;

        // Tous les outils sont disponibles, quel que soit le niveau de
        // confiance du contenu traité (pas de séparation de privilège).
        var response = await _llmClient.CompleteAsync(fullPrompt, _toolRegistry.AllTools);

        // Chaque appel d'outil suggéré par le modèle est exécuté directement,
        // sans validation de schéma ni confirmation pour les actions sensibles.
        foreach (var call in response.ToolCalls)
        {
            var tool = _toolRegistry.Get(call.ToolName);
            await tool.ExecuteAsync(call.Arguments);
        }

        return Ok(response.Text);
    }
}

public record ChatRequest(string Message);
