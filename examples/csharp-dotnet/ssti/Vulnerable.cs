// Vulnérable : Server-Side Template Injection (SSTI) (CWE-1336)
// Le texte du template lui-même est construit par concaténation d'une
// entrée utilisateur avant compilation par le moteur (Scriban), permettant
// à un attaquant d'injecter de la syntaxe de template interprétée par le
// moteur (au lieu d'une simple variable de contexte).
using Microsoft.AspNetCore.Mvc;
using Scriban;

[ApiController]
[Route("api/greeting")]
public class GreetingController : ControllerBase
{
    [HttpGet]
    public IActionResult Greet([FromQuery] string name)
    {
        // Concaténation directe dans le texte du template avant parsing.
        var templateText = "Bonjour " + name + ", bienvenue !";
        var template = Template.Parse(templateText);

        return Ok(template.Render());
    }
}
