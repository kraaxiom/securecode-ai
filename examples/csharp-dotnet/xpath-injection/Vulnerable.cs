// Vulnérable : XPath Injection (CWE-643)
// Les identifiants sont concaténés directement dans l'expression XPath
// utilisée pour l'authentification, permettant à un attaquant de modifier
// la logique de sélection (ex: contournement d'authentification).
using Microsoft.AspNetCore.Mvc;
using System.Xml;

[ApiController]
[Route("api/auth")]
public class XmlAuthController : ControllerBase
{
    private readonly XmlDocument _users = LoadUsers();

    [HttpPost("login")]
    public IActionResult Login([FromForm] string user, [FromForm] string pass)
    {
        // user = "' or '1'='1" contourne la vérification.
        var expr = $"//user[username='{user}' and password='{pass}']";
        var node = _users.SelectSingleNode(expr);

        return node == null ? Unauthorized() : Ok("Authentifié");
    }

    private static XmlDocument LoadUsers()
    {
        var doc = new XmlDocument();
        doc.LoadXml("<users><user><username>alice</username><password>secret</password></user></users>");
        return doc;
    }
}
