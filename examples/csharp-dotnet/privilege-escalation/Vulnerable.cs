// Vulnérable : Privilege Escalation (CWE-269)
// Le rôle d'un utilisateur est mis à jour depuis une valeur fournie par le
// client, sans vérifier que l'appelant est lui-même autorisé à accorder ce
// niveau de privilège, et sans invalider les sessions existantes.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/users")]
[Authorize] // vérifie seulement que l'appelant est connecté, peu importe son propre rôle
public class UserRolesController : ControllerBase
{
    private readonly AppDbContext _db;

    public UserRolesController(AppDbContext db)
    {
        _db = db;
    }

    public class AssignRoleRequest
    {
        public string Role { get; set; } = string.Empty; // le client contrôle entièrement ce champ
    }

    [HttpPut("{id:int}/role")]
    public IActionResult AssignRole(int id, [FromBody] AssignRoleRequest request)
    {
        // Aucune vérification que l'appelant a le droit d'accorder ce rôle précis :
        // un utilisateur standard authentifié peut se nommer (ou nommer un tiers) "Admin".
        var user = _db.Users.First(u => u.Id == id);
        user.Role = request.Role;
        _db.SaveChanges();

        // Les sessions/tokens déjà émis restent valides avec l'ancien niveau de
        // privilège, mais le nouveau rôle prend effet immédiatement côté données.
        return Ok(user);
    }
}
