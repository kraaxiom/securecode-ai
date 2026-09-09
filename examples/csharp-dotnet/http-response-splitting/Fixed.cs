// Corrigé : HTTP Response Splitting (CWE-113)
// La cible de redirection est validée contre une liste blanche de chemins
// autorisés, et la redirection utilise l'API sûre Redirect() du framework.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("account")]
public class LoginRedirectController : ControllerBase
{
    private static readonly HashSet<string> AllowedPaths = new() { "/dashboard", "/profile", "/account" };

    [HttpGet("post-login")]
    public IActionResult PostLogin([FromQuery] string next)
    {
        var safeNext = AllowedPaths.Contains(next) ? next : "/dashboard";
        return Redirect(safeNext);
    }
}
