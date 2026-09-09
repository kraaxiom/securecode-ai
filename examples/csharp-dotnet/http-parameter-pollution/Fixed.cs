// Corrigé : HTTP Parameter Pollution (CWE-235)
// Le nombre d'occurrences du paramètre "role" est vérifié explicitement via
// la query string brute ; toute duplication est rejetée plutôt que résolue
// silencieusement par le binder.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/accounts")]
public class RoleAssignmentController : ControllerBase
{
    [HttpPost("assign-role")]
    public IActionResult AssignRole([FromServices] IUserService users)
    {
        var values = Request.Query["role"];
        if (values.Count != 1)
            return BadRequest("Paramètre 'role' dupliqué ou manquant.");

        users.AssignRole(User.Identity!.Name!, values[0]!);
        return Ok();
    }
}
