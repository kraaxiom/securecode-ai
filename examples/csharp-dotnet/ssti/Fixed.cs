// Corrigé : Server-Side Template Injection (SSTI) (CWE-1336)
// Le template est un texte statique, contrôlé par le développeur, versionné
// séparément. L'entrée utilisateur est passée uniquement comme variable de
// contexte au moteur de rendu, jamais concaténée à la structure du template.
using Microsoft.AspNetCore.Mvc;
using Scriban;

[ApiController]
[Route("api/greeting")]
public class GreetingController : ControllerBase
{
    // Template statique, défini une fois par le développeur.
    private static readonly Template GreetingTemplate = Template.Parse("Bonjour {{ name }}, bienvenue !");

    [HttpGet]
    public IActionResult Greet([FromQuery] string name)
    {
        if (string.IsNullOrWhiteSpace(name))
            return BadRequest("Nom invalide.");

        // "name" transite uniquement comme donnée de contexte.
        var result = GreetingTemplate.Render(new { name });
        return Ok(result);
    }
}
