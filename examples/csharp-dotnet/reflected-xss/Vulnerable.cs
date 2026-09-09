// Vulnérable : Reflected Cross-Site Scripting (CWE-79)
// Le terme de recherche fourni dans la query string est réinjecté directement
// dans la réponse HTML par concaténation de chaînes, sans passer par un
// moteur de templates avec échappement automatique ni encodage explicite.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("search")]
public class SearchController : ControllerBase
{
    [HttpGet]
    public ContentResult Search([FromQuery] string q)
    {
        var results = ProductCatalog.Find(q);

        var html = $@"
            <!DOCTYPE html>
            <html>
            <body>
                <h1>Résultats pour : {q}</h1>
                <ul>{string.Join("", results.Select(r => $"<li>{r.Name}</li>"))}</ul>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
