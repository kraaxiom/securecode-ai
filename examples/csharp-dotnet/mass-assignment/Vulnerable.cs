// Vulnérable : Mass Assignment (CWE-915)
// Le modèle interne complet, y compris le champ 'IsAdmin', est lié directement
// au corps de la requête entrante sans liste blanche de champs autorisés.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

public class User
{
    public int Id { get; set; }
    public string Name { get; set; } = string.Empty;
    public string Email { get; set; } = string.Empty;
    public bool IsAdmin { get; set; } // attribut sensible porté par le même modèle
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
    public IActionResult UpdateProfile(int id, [FromBody] User payload)
    {
        // Le modèle de requête est le même que le modèle interne complet :
        // un attaquant peut inclure "isAdmin": true dans le JSON envoyé et
        // s'octroyer des privilèges d'administrateur via ce endpoint de profil.
        var user = _db.Users.First(u => u.Id == id);
        user.Name = payload.Name;
        user.Email = payload.Email;
        user.IsAdmin = payload.IsAdmin;
        _db.SaveChanges();

        return Ok(user);
    }
}
