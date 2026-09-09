// Corrigé : DOM-based Cross-Site Scripting (CWE-79)
// La donnée issue de la requête n'est plus concaténée dans un script inline.
// Elle est sérialisée en JSON échappé pour la sécurité HTML/JS et le script
// client utilise le sink texte (textContent) plutôt qu'un sink HTML (innerHTML).
using System.Text.Encodings.Web;
using System.Text.Json;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("search")]
public class SearchPageController : ControllerBase
{
    private static readonly JsonSerializerOptions JsonOptions = new()
    {
        Encoder = JavaScriptEncoder.Default
    };

    [HttpGet]
    public ContentResult Index([FromQuery] string q)
    {
        // Sérialisation JSON sûre : les caractères spéciaux HTML/JS sont échappés
        // par défaut par l'encodeur strict de System.Text.Json.
        var safeQueryJson = JsonSerializer.Serialize(q ?? string.Empty, JsonOptions);

        var html = $@"
            <!DOCTYPE html>
            <html>
            <body>
                <div id='result'></div>
                <script>
                    var query = {safeQueryJson};
                    // Sink texte : la valeur est traitée comme du texte inerte,
                    // jamais interprétée comme du HTML.
                    document.getElementById('result').textContent = query;
                </script>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
