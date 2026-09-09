// Corrigé : Prompt Injection (CWE-1427)
// Le prompt est construit via les rôles structurés de l'API (system/user),
// la sortie du modèle est traitée comme non fiable avant toute action,
// et les outils exposés sont restreints selon le principe du moindre privilège.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/chat")]
public class ChatController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly IToolRegistry _toolRegistry;
    private readonly IToolSchemaValidator _schemaValidator;
    private readonly IAuditLog _auditLog;

    private const string SystemPrompt =
        "Tu es un assistant interne. Réponds aux questions des employés.";

    // Seuls les outils à faible impact sont disponibles par défaut.
    private static readonly string[] DefaultAllowedTools = { "search_knowledge_base" };
    private static readonly string[] HighImpactTools = { "send_email", "delete_record" };

    public ChatController(
        ILlmClient llmClient,
        IToolRegistry toolRegistry,
        IToolSchemaValidator schemaValidator,
        IAuditLog auditLog)
    {
        _llmClient = llmClient;
        _toolRegistry = toolRegistry;
        _schemaValidator = schemaValidator;
        _auditLog = auditLog;
    }

    [HttpPost("message")]
    public async Task<IActionResult> SendMessage([FromBody] ChatRequest request)
    {
        // Séparation structurée des rôles : l'entrée utilisateur reste dans
        // le rôle "user", distincte du rôle "system" porté par le développeur.
        var response = await _llmClient.ChatAsync(new[]
        {
            new LlmMessage("system", SystemPrompt),
            new LlmMessage("user", request.Message)
        }, _toolRegistry.GetByNames(DefaultAllowedTools));

        foreach (var call in response.ToolCalls)
        {
            // La sortie du modèle est traitée comme une donnée non fiable :
            // validation applicative indépendante contre un schéma strict.
            if (!_schemaValidator.IsValid(call.ToolName, call.Arguments))
            {
                _auditLog.Record("rejected_invalid_tool_call", call.ToolName);
                continue;
            }

            // Moindre privilège : les actions à fort impact exigent une
            // confirmation humaine explicite avant exécution.
            if (HighImpactTools.Contains(call.ToolName))
            {
                _auditLog.Record("high_impact_tool_requires_confirmation", call.ToolName);
                continue;
            }

            var tool = _toolRegistry.Get(call.ToolName);
            await tool.ExecuteAsync(call.Arguments);
        }

        // Journalisation du prompt système et de la sortie utilisée pour
        // des décisions automatisées.
        _auditLog.Record("chat_turn", request.Message, response.Text);

        return Ok(response.Text);
    }
}

public record ChatRequest(string Message);
public record LlmMessage(string Role, string Content);
