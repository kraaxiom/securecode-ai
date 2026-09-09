// Corrigé : Broken Function Level Authorization (CWE-862)
// Un contrôle de rôle explicite et centralisé (policy ASP.NET Core) est appliqué
// sur la fonction sensible, indépendamment de sa visibilité dans l'interface.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/users")]
[Authorize]
public class UsersAdminController : ControllerBase
{
    private readonly IUserService _userService;

    public UsersAdminController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpDelete("{id:int}")]
    [Authorize(Policy = "RequireAdminRole")] // vérification de rôle explicite et centralisée
    public IActionResult DeleteUser(int id)
    {
        _userService.Delete(id);
        return NoContent();
    }
}

// Program.cs / Startup.cs — définition centralisée de la policy, appliquée
// systématiquement à chaque endpoint de niveau administrateur.
// builder.Services.AddAuthorization(options =>
// {
//     options.AddPolicy("RequireAdminRole", policy => policy.RequireRole("Admin"));
// });
