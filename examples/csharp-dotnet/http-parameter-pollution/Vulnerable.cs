// Vulnérable : HTTP Parameter Pollution (CWE-235)
// Lorsque le paramètre "role" est envoyé plusieurs fois (?role=user&role=admin),
// le binding par défaut de ASP.NET Core peut retenir une valeur différente
// de celle vue par un WAF/proxy en amont, créant une divergence exploitable.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/accounts")]
public class RoleAssignmentController : ControllerBase
{
    [HttpPost("assign-role")]
    public IActionResult AssignRole([FromQuery] string role, [FromServices] IUserService users)
    {
        // Aucune vérification que "role" n'a été fourni qu'une seule fois.
        users.AssignRole(User.Identity!.Name!, role);
        return Ok();
    }
}
