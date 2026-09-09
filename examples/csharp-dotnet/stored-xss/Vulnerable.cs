// Vulnérable : Stored Cross-Site Scripting (CWE-79)
// La biographie d'un profil utilisateur, persistée en base de données, est
// réaffichée telle quelle sur les pages consultées par d'autres utilisateurs,
// sans encodage de sortie. Le contenu malveillant reste actif tant qu'il est
// stocké et affecte tout visiteur du profil, y compris des administrateurs.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("profile")]
public class ProfileController : ControllerBase
{
    private readonly IUserRepository _users;

    public ProfileController(IUserRepository users)
    {
        _users = users;
    }

    [HttpGet("{id}")]
    public ContentResult View(int id)
    {
        var user = _users.GetById(id);

        var html = $@"
            <div class='profile'>
                <h1>{user.DisplayName}</h1>
                <p class='bio'>{user.Bio}</p>
            </div>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
