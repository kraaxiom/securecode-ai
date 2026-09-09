// Corrigé : CRLF Injection (CWE-93)
// La cible de redirection est validée (chemin relatif, sans caractères de
// contrôle) et le mécanisme sûr Redirect() du framework est utilisé plutôt
// qu'une écriture manuelle de l'en-tête.
using Microsoft.AspNetCore.Mvc;
using System.Text.RegularExpressions;

[ApiController]
[Route("account")]
public class RedirectController : ControllerBase
{
    private static readonly Regex SafeRelativePath = new(@"^/[^\r\n]*$", RegexOptions.Compiled);

    [HttpGet("redirect")]
    public IActionResult RedirectTo([FromQuery] string next)
    {
        var safeNext = !string.IsNullOrEmpty(next) && SafeRelativePath.IsMatch(next) ? next : "/";
        return Redirect(safeNext);
    }
}
