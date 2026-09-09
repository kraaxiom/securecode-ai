// Corrigé : Expression Language Injection (CWE-917)
// L'expression n'est plus construite depuis l'entrée utilisateur : seule
// une opération choisie dans une liste blanche est exécutée, avec des
// opérandes numériques validés transmis en paramètres, jamais en code.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/rules")]
public class RuleEngineController : ControllerBase
{
    private static readonly Dictionary<string, Func<double, double, double>> AllowedOperations = new()
    {
        ["gt"] = (a, b) => a > b ? 1 : 0,
        ["lt"] = (a, b) => a < b ? 1 : 0,
        ["eq"] = (a, b) => a == b ? 1 : 0,
    };

    [HttpPost("evaluate")]
    public IActionResult Evaluate([FromBody] RuleRequest request)
    {
        if (!AllowedOperations.TryGetValue(request.Operator, out var op))
            return BadRequest("Opérateur non autorisé.");

        var result = op(request.Left, request.Right);
        return Ok(new { result });
    }
}

public class RuleRequest
{
    public string Operator { get; set; } = string.Empty;
    public double Left { get; set; }
    public double Right { get; set; }
}
