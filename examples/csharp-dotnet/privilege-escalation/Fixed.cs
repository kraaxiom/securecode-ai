// Corrigé : Privilege Escalation (CWE-269)
// L'attribution de rôle vérifie explicitement que l'appelant est autorisé à
// accorder le niveau de privilège demandé, puis invalide les sessions actives
// de l'utilisateur ciblé pour forcer une réauthentification avec le bon rôle.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/users")]
[Authorize]
public class UserRolesController : ControllerBase
{
    private readonly AppDbContext _db;
    private readonly ISessionRevocationService _sessions;

    public UserRolesController(AppDbContext db, ISessionRevocationService sessions)
    {
        _db = db;
        _sessions = sessions;
    }

    public class AssignRoleRequest
    {
        public string Role { get; set; } = string.Empty;
    }

    [HttpPut("{id:int}/role")]
    public IActionResult AssignRole(int id, [FromBody] AssignRoleRequest request)
    {
        // Vérifie que l'appelant possède lui-même un niveau de privilège
        // suffisant pour accorder le rôle demandé (ex. seul un Admin peut
        // créer un autre Admin ; un Manager ne peut promouvoir qu'en dessous).
        if (!User.CanGrantRole(request.Role))
            return Forbid();

        var user = _db.Users.First(u => u.Id == id);
        user.Role = request.Role;
        _db.SaveChanges();

        // Invalide les sessions/tokens actifs de l'utilisateur modifié afin
        // d'éviter la persistance d'un ancien niveau de privilège en mémoire.
        _sessions.RevokeAllForUser(user.Id);

        return Ok(user);
    }
}
