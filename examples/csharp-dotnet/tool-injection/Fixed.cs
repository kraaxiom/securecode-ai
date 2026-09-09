// Corrigé : Tool Injection (CWE-1427)
// Chaque appel d'outil est validé contre un schéma typé strict, restreint
// à une liste blanche contextuelle selon le niveau de confiance du contenu
// traité, et journalisé avant exécution.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/agent")]
public class AgentController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly IToolRegistry _toolRegistry;
    private readonly IToolSchemaValidator _schemaValidator;
    private readonly IAuditLog _auditLog;

    // Liste blanche contextuelle : seuls certains outils sont autorisés
    // lorsque le contenu traité provient d'une source non fiable.
    private static readonly IReadOnlyDictionary<TrustLevel, string[]> AllowedToolsByTrust =
        new Dictionary<TrustLevel, string[]>
        {
            [TrustLevel.Trusted] = new[] { "search_docs", "send_email", "create_ticket" },
            [TrustLevel.Untrusted] = new[] { "search_docs" }
        };

    public AgentController(
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

    [HttpPost("run")]
    public async Task<IActionResult> Run([FromBody] AgentRequest request)
    {
        var trustLevel = request.ContentTrustLevel;
        var allowedTools = _toolRegistry.GetByNames(AllowedToolsByTrust[trustLevel]);

        var response = await _llmClient.CompleteAsync(request.Instruction, allowedTools);

        var results = new List<object>();
        foreach (var call in response.ToolCalls)
        {
            // Liste blanche contextuelle : l'outil doit être autorisé pour
            // le niveau de confiance du contenu en cours de traitement.
            if (!AllowedToolsByTrust[trustLevel].Contains(call.ToolName))
            {
                _auditLog.Record("blocked_tool_not_in_allowlist", call.ToolName, trustLevel.ToString());
                continue;
            }

            // Validation stricte des arguments contre un schéma typé,
            // indépendamment de ce que le modèle a généré.
            if (!_schemaValidator.TryValidate(call.ToolName, call.Arguments, out var validatedArgs))
            {
                _auditLog.Record("rejected_invalid_arguments", call.ToolName);
                continue;
            }

            var tool = _toolRegistry.Get(call.ToolName);
            var result = await tool.ExecuteAsync(validatedArgs);
            _auditLog.Record("tool_executed", call.ToolName, validatedArgs);
            results.Add(result);
        }

        return Ok(results);
    }
}

public enum TrustLevel { Trusted, Untrusted }
public record AgentRequest(string Instruction, TrustLevel ContentTrustLevel);
