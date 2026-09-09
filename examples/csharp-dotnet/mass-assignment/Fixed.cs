// Corrigé : Mass Assignment (CWE-915)
// Un DTO d'entrée dédié, distinct du modèle interne, expose une liste blanche
// explicite des champs modifiables par le client. Le champ 'IsAdmin' n'y
// figure jamais et ne peut donc pas transiter par ce endpoint.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = string.Empty;
    public string Email { get; set; } = string.Empty;
    public bool IsAdmin { get; set; }
}

// DTO d'entrée : liste blanche explicite, séparée du modèle de données interne.
public class UpdateUserProfileRequest
{
    public string Name { get; set; } = string.Empty;
    public string Email { get; set; } = string.Empty;
}

[ApiController]
[Route("api/users")]
[Authorize]
public class UsersController : ControllerBase
{
    private readonly AppDbContext _db;

    public UsersController(AppDbContext db)
    {
        _db = db;
    }

    [HttpPut("{id:int}")]
    public IActionResult UpdateProfile(int id, [FromBody] UpdateUserProfileRequest payload)
    {
        var user = _db.Users.First(u => u.Id == id);

        // Seuls les champs explicitement whitelistés sont assignables.
        // 'IsAdmin' n'existe pas dans le DTO : il ne peut jamais être injecté ici.
        user.Name = payload.Name;
        user.Email = payload.Email;
        _db.SaveChanges();

        return Ok(user);
    }
}
