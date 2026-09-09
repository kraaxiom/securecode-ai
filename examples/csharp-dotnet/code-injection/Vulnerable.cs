// Vulnérable : Code Injection (CWE-94)
// L'application compile et exécute dynamiquement une expression C#
// fournie par l'utilisateur via Roslyn Scripting, lui donnant un contrôle
// total sur le code exécuté côté serveur.
using Microsoft.AspNetCore.Mvc;
using Microsoft.CodeAnalysis.CSharp.Scripting;

[ApiController]
[Route("api/calc")]
public class FormulaController : ControllerBase
{
    [HttpPost("evaluate")]
    public async Task<IActionResult> Evaluate([FromBody] FormulaRequest request)
    {
        // Le contenu de request.Expression est du code C# arbitraire évalué tel quel.
        var result = await CSharpScript.EvaluateAsync<double>(request.Expression);
        return Ok(new { result });
    }
}

public class FormulaRequest
{
    public string Expression { get; set; } = string.Empty;
}
