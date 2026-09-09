// Corrigé : Reflected Cross-Site Scripting (CWE-79)
// Toute valeur issue de la requête (le terme de recherche) est encodée
// contextuellement avant insertion dans le HTML de réponse via HtmlEncoder,
// l'équivalent .NET de l'échappement automatique d'un moteur de templates.
using System.Text.Encodings.Web;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("search")]
public class SearchController : ControllerBase
{
    private readonly HtmlEncoder _htmlEncoder;

    public SearchController(HtmlEncoder htmlEncoder)
    {
        _htmlEncoder = htmlEncoder;
    }

    [HttpGet]
    public ContentResult Search([FromQuery] string q)
    {
        var results = ProductCatalog.Find(q);
        var safeQuery = _htmlEncoder.Encode(q ?? string.Empty);

        var html = $@"
            <!DOCTYPE html>
            <html>
            <body>
                <h1>Résultats pour : {safeQuery}</h1>
                <ul>{string.Join("", results.Select(r => $"<li>{_htmlEncoder.Encode(r.Name)}</li>"))}</ul>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}

// Recommandation : dans une vue Razor MVC, préférer @q (encodage automatique)
// à Html.Raw(q), qui désactive explicitement la protection.
