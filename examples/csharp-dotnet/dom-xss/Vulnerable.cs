// Vulnérable : DOM-based Cross-Site Scripting (CWE-79)
// Le contrôleur ASP.NET Core sert une page dont le script inline concatène
// directement un paramètre de requête dans une affectation JavaScript, ce qui
// permet à l'attaquant de sortir du contexte chaîne et d'injecter du code
// exécuté côté client (sink DOM non fiable dans le script généré).
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("search")]
public class SearchPageController : ControllerBase
{
    [HttpGet]
    public ContentResult Index([FromQuery] string q)
    {
        var html = $@"
            <!DOCTYPE html>
            <html>
            <body>
                <div id='result'></div>
                <script>
                    // La valeur brute de la requête est injectée dans le script
                    // puis écrite dans le DOM via un sink dangereux côté client.
                    var query = ""{q}"";
                    document.getElementById('result').innerHTML = query;
                </script>
            </body>
            </html>";

        return new ContentResult { Content = html, ContentType = "text/html" };
    }
}
