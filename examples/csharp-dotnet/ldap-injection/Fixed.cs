// Corrigé : LDAP Injection (CWE-90)
// L'entrée utilisateur est échappée avec les règles LDAP (RFC 4515) avant
// d'être insérée dans le filtre, et son format est validé au préalable.
using Microsoft.AspNetCore.Mvc;
using System.DirectoryServices;
using System.Text;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/directory")]
public class DirectorySearchController : ControllerBase
{
    [HttpGet("users")]
    public IActionResult FindUser([FromQuery] string uid)
    {
        // Validation stricte du format attendu (identifiant alphanumérique).
        if (string.IsNullOrEmpty(uid) || !Regex.IsMatch(uid, "^[a-zA-Z0-9._-]{1,64}$"))
            return BadRequest("Identifiant invalide.");

        using var entry = new DirectoryEntry("LDAP://dc=example,dc=com");
        using var searcher = new DirectorySearcher(entry)
        {
            Filter = $"(uid={EscapeLdapFilter(uid)})"
        };

        var result = searcher.FindOne();
        if (result == null)
            return NotFound();

        return Ok(result.Properties["cn"][0]?.ToString());
    }

    // Échappe les caractères spéciaux d'un filtre LDAP selon RFC 4515.
    private static string EscapeLdapFilter(string input)
    {
        var sb = new StringBuilder();
        foreach (var c in input)
        {
            switch (c)
            {
                case '\\': sb.Append(@"\5c"); break;
                case '*': sb.Append(@"\2a"); break;
                case '(': sb.Append(@"\28"); break;
                case ')': sb.Append(@"\29"); break;
                case '\0': sb.Append(@"\00"); break;
                default: sb.Append(c); break;
            }
        }
        return sb.ToString();
    }
}
