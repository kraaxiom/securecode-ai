// Vulnérable : Mutation-based Cross-Site Scripting / mXSS (CWE-79)
// Le contenu d'un éditeur riche est "sanitisé" par un filtrage naïf de balises
// (liste noire simpliste) puis stocké et réaffiché tel quel. Ce type de
// filtrage ne tient pas compte des quirks de reparsing HTML du navigateur :
// un markup jugé sûr côté serveur peut être réinterprété différemment une
// fois inséré dans le DOM côté client, faisant réapparaître une charge active.
using System.Text.RegularExpressions;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("notes")]
public class RichNoteController : ControllerBase
{
    private readonly INoteRepository _notes;

    public RichNoteController(INoteRepository notes)
    {
        _notes = notes;
    }

    [HttpPost]
    public IActionResult Save([FromForm] string content)
    {
        // Filtrage naïf par expression régulière : retire uniquement la balise
        // <script> littérale, sans comprendre la grammaire HTML ni les
        // reformulations possibles lors du reparsing par le navigateur.
        var naivelyFiltered = Regex.Replace(content, "<script.*?>.*?</script>", string.Empty,
            RegexOptions.IgnoreCase | RegexOptions.Singleline);

        _notes.Save(new Note { Content = naivelyFiltered });
        return Ok();
    }

    [HttpGet("{id}")]
    public ContentResult View(int id)
    {
        var note = _notes.GetById(id);

        // Le contenu "nettoyé" est réinséré tel quel : aucune re-sanitisation
        // au moment de l'affichage, alors que le navigateur peut muter le
        // markup stocké lors de son insertion dans le DOM (éditeur riche).
        var html = $"<div class='note-content'>{note.Content}</div>";
        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
