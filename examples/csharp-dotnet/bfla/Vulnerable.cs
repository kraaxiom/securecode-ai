// Vulnérable : Broken Function Level Authorization (CWE-862)
// L'endpoint d'administration ne vérifie que l'authentification de l'appelant,
// sans contrôler son rôle. Tout utilisateur connecté peut donc invoquer une
// fonction réservée aux administrateurs simplement en connaissant l'URL.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/users")]
[Authorize] // vérifie uniquement que l'utilisateur est authentifié
public class UsersAdminController : ControllerBase
{
    private readonly IUserService _userService;

    public UsersAdminController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpDelete("{id:int}")]
    public IActionResult DeleteUser(int id)
    {
        // Aucune vérification du rôle de l'appelant : un utilisateur standard
        // peut supprimer n'importe quel compte, y compris un administrateur.
        _userService.Delete(id);
        return NoContent();
    }
}
