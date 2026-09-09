// Corrigé : Agent Hijacking (CWE-1427)
// La couche de décision (LLM) est séparée de la couche d'exécution : les
// outils à fort impact exigent une confirmation humaine explicite, chaque
// étape est revalidée par des règles métier indépendantes du modèle, et
// toutes les décisions/actions sont journalisées pour l'audit.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/agent")]
public class AgentController : ControllerBase
{
    private static readonly HashSet<string> HighImpactTools = new(StringComparer.OrdinalIgnoreCase)
    {
        "send_payment",
        "delete_account"
    };

    private readonly ILlmPlanner _planner;
    private readonly IToolRegistry _tools;
    private readonly IHumanConfirmationService _confirmation;
    private readonly IBusinessRuleValidator _validator;
    private readonly IAuditLog _auditLog;

    public AgentController(
        ILlmPlanner planner,
        IToolRegistry tools,
        IHumanConfirmationService confirmation,
        IBusinessRuleValidator validator,
        IAuditLog auditLog)
    {
        _planner = planner;
        _tools = tools;
        _confirmation = confirmation;
        _validator = validator;
        _auditLog = auditLog;
    }

    [HttpPost("run")]
    public async Task<IActionResult> Run([FromBody] AgentRequest request)
    {
        var plan = await _planner.PlanAsync(request.UserInput, _tools.All);

        var results = new List<object>();
        foreach (var step in plan.Steps)
        {
            if (HighImpactTools.Contains(step.ToolName))
            {
                // Confirmation humaine obligatoire avant toute action irréversible.
                var confirmed = await _confirmation.RequestAsync(request.UserId, step);
                if (!confirmed)
                {
                    await _auditLog.RecordAsync("blocked_high_impact_action", step);
                    continue;
                }
            }

            // Validation métier indépendante du modèle, quel que soit l'outil.
            var validation = await _validator.ValidateAsync(step);
            if (!validation.IsValid)
            {
                await _auditLog.RecordAsync("blocked_invalid_step", step);
                continue;
            }

            var tool = _tools.Resolve(step.ToolName);
            var output = await tool.ExecuteAsync(step.Arguments);
            await _auditLog.RecordAsync("action_executed", step);
            results.Add(output);
        }

        return Ok(results);
    }
}
