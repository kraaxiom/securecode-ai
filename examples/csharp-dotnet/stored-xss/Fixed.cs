// Corrigé : Stored Cross-Site Scripting (CWE-79)
// Toute donnée persistée d'origine utilisateur (nom affiché, biographie) est
// encodée contextuellement à l'affichage. Si du HTML riche doit un jour être
// autorisé dans la biographie, il devra être sanitisé avant stockage via une
// bibliothèque dédiée (ex: HtmlSanitizer), pas simplement encodé à l'affichage.
using System.Text.Encodings.Web;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("profile")]
public class ProfileController : ControllerBase
{
    private readonly IUserRepository _users;
    private readonly HtmlEncoder _htmlEncoder;

    public ProfileController(IUserRepository users, HtmlEncoder htmlEncoder)
    {
        _users = users;
        _htmlEncoder = htmlEncoder;
    }

    [HttpGet("{id}")]
    public ContentResult View(int id)
    {
        var user = _users.GetById(id);

        var safeDisplayName = _htmlEncoder.Encode(user.DisplayName ?? string.Empty);
        var safeBio = _htmlEncoder.Encode(user.Bio ?? string.Empty);

        var html = $@"
            <div class='profile'>
                <h1>{safeDisplayName}</h1>
                <p class='bio'>{safeBio}</p>
            </div>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
