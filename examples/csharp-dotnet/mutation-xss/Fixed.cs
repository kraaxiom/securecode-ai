// Corrigé : Mutation-based Cross-Site Scripting / mXSS (CWE-79)
// Remplacement du filtrage naïf par expression régulière par une bibliothèque
// de sanitisation HTML activement maintenue, tenant compte des quirks de
// parsing du navigateur. La re-sanitisation est appliquée à chaque étape de
// sérialisation (sauvegarde ET affichage), afin de ne jamais faire confiance
// à un contenu déjà "nettoyé" une fois transformé par un aller-retour DOM.
using Ganss.Xss;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("notes")]
public class RichNoteController : ControllerBase
{
    private readonly INoteRepository _notes;
    private readonly HtmlSanitizer _sanitizer;

    public RichNoteController(INoteRepository notes)
    {
        _notes = notes;

        // Bibliothèque de sanitisation HTML maintenue (Ganss.XSS), avec une
        // liste blanche stricte de balises et d'attributs autorisés.
        _sanitizer = new HtmlSanitizer();
        _sanitizer.AllowedTags.Clear();
        _sanitizer.AllowedTags.UnionWith(new[] { "p", "b", "i", "u", "a", "ul", "ol", "li" });
        _sanitizer.AllowedAttributes.Clear();
        _sanitizer.AllowedAttributes.Add("href");
    }

    [HttpPost]
    public IActionResult Save([FromForm] string content)
    {
        var clean = _sanitizer.Sanitize(content);
        _notes.Save(new Note { Content = clean });
        return Ok();
    }

    [HttpGet("{id}")]
    public ContentResult View(int id)
    {
        var note = _notes.GetById(id);

        // Nouvelle passe de sanitisation à l'affichage : même si le contenu a
        // déjà été nettoyé avant stockage, on ne fait pas confiance à une
        // sanitisation antérieure qui pourrait avoir été contournée ou dont
        // le résultat aurait pu muter (bibliothèque mise à jour depuis, etc.).
        var safeHtml = $"<div class='note-content'>{_sanitizer.Sanitize(note.Content)}</div>";
        return new ContentResult { Content = safeHtml, ContentType = "text/html" };
    }
}
