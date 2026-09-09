// Vulnérable : Tool Injection (CWE-1427)
// Le nom d'outil et les arguments générés par le modèle sont transmis
// directement à l'exécution, sans validation de schéma ni liste blanche
// contextuelle, ce qui permet à du contenu non fiable de détourner l'appel.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/agent")]
public class AgentController : ControllerBase
{
    private readonly ILlmClient _llmClient;
    private readonly IToolRegistry _toolRegistry;

    public AgentController(ILlmClient llmClient, IToolRegistry toolRegistry)
    {
        _llmClient = llmClient;
        _toolRegistry = toolRegistry;
    }

    [HttpPost("run")]
    public async Task<IActionResult> Run([FromBody] AgentRequest request)
    {
        var response = await _llmClient.CompleteAsync(request.Instruction, _toolRegistry.AllTools);

        var results = new List<object>();
        foreach (var call in response.ToolCalls)
        {
            // Le nom de l'outil et ses arguments, générés par le modèle,
            // sont utilisés tels quels pour construire et exécuter l'appel,
            // sans validation de schéma ni contrôle de privilège.
            var tool = _toolRegistry.Get(call.ToolName);
            var result = await tool.ExecuteAsync(call.Arguments);
            results.Add(result);
        }

        return Ok(results);
    }
}

public record AgentRequest(string Instruction);
