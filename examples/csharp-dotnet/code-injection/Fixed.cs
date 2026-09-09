// Corrigé : Code Injection (CWE-94)
// Suppression totale de l'évaluation dynamique de code. L'opération demandée
// est résolue via une liste blanche de fonctions autorisées et des
// opérandes numériques validés.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/calc")]
public class FormulaController : ControllerBase
{
    private static readonly Dictionary<string, Func<double, double, double>> AllowedOperations = new()
    {
        ["add"] = (a, b) => a + b,
        ["sub"] = (a, b) => a - b,
        ["mul"] = (a, b) => a * b,
        ["div"] = (a, b) => a / b,
    };

    [HttpPost("evaluate")]
    public IActionResult Evaluate([FromBody] FormulaRequest request)
    {
        if (!AllowedOperations.TryGetValue(request.Operation, out var operation))
            return BadRequest("Opération non autorisée.");

        var result = operation(request.A, request.B);
        return Ok(new { result });
    }
}

public class FormulaRequest
{
    public string Operation { get; set; } = string.Empty;
    public double A { get; set; }
    public double B { get; set; }
}
