// Vulnérable : Expression Language Injection (CWE-917)
// L'expression évaluée dynamiquement (via DataTable.Compute, équivalent .NET
// d'un moteur d'expression) est construite par concaténation directe d'une
// entrée utilisateur, permettant l'exécution d'expressions arbitraires.
using Microsoft.AspNetCore.Mvc;
using System.Data;

[ApiController]
[Route("api/rules")]
public class RuleEngineController : ControllerBase
{
    [HttpPost("evaluate")]
    public IActionResult Evaluate([FromBody] RuleRequest request)
    {
        var table = new DataTable();
        // L'expression utilisateur est évaluée telle quelle par le moteur.
        var result = table.Compute(request.Expression, string.Empty);
        return Ok(new { result });
    }
}

public class RuleRequest
{
    public string Expression { get; set; } = string.Empty;
}
