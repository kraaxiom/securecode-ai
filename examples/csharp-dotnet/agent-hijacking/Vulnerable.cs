// Vulnérable : Agent Hijacking (CWE-1427)
// L'agent exécute directement chaque étape planifiée par le LLM, y compris
// les outils à fort impact (paiement, suppression de compte), sans aucune
// validation métier indépendante ni confirmation humaine préalable. Les
// privilèges de l'agent restent identiques même lorsqu'il traite du
// contenu externe non fiable dans la même session.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/agent")]
public class AgentController : ControllerBase
{
    private readonly ILlmPlanner _planner;
    private readonly IToolRegistry _tools;

    public AgentController(ILlmPlanner planner, IToolRegistry tools)
    {
        _planner = planner;
        _tools = tools;
    }

    [HttpPost("run")]
    public async Task<IActionResult> Run([FromBody] AgentRequest request)
    {
        // Le plan peut inclure des outils sensibles (send_payment, delete_account)
        // que le modèle choisit librement en fonction de l'entrée utilisateur
        // ou d'un contenu externe déjà présent dans le contexte.
        var plan = await _planner.PlanAsync(request.UserInput, _tools.All);

        var results = new List<object>();
        foreach (var step in plan.Steps)
        {
            // Exécution directe de l'outil, sans distinction entre actions
            // anodines et actions irréversibles/à fort impact.
            var tool = _tools.Resolve(step.ToolName);
            var output = await tool.ExecuteAsync(step.Arguments);
            results.Add(output);
        }

        return Ok(results);
    }
}
