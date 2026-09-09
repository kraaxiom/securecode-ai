// Vulnérable : LDAP Injection (CWE-90)
// L'identifiant fourni par l'utilisateur est concaténé directement dans le
// filtre de recherche LDAP, permettant de modifier la logique du filtre
// (ex: contournement d'authentification via injection de méta-caractères
// LDAP comme '*', '(', ')').
using Microsoft.AspNetCore.Mvc;
using System.DirectoryServices;

[ApiController]
[Route("api/directory")]
public class DirectorySearchController : ControllerBase
{
    [HttpGet("users")]
    public IActionResult FindUser([FromQuery] string uid)
    {
        using var entry = new DirectoryEntry("LDAP://dc=example,dc=com");
        using var searcher = new DirectorySearcher(entry)
        {
            // Concaténation directe : uid = "*)(uid=*" contourne le filtre attendu.
            Filter = $"(uid={uid})"
        };

        var result = searcher.FindOne();
        if (result == null)
            return NotFound();

        return Ok(result.Properties["cn"][0]?.ToString());
    }
}
