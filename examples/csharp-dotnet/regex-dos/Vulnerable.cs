// Vulnérable : ReDoS - Regular Expression Denial of Service (CWE-1333)
// La regex contient des quantificateurs imbriqués ((a+)+), et est appliquée
// à une entrée utilisateur non bornée en taille, sans timeout d'exécution.
// Une chaîne pathologique (ex: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!") peut
// provoquer un temps de calcul exponentiel bloquant le thread appelant.
using Microsoft.AspNetCore.Mvc;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/validation")]
public class EmailValidationController : ControllerBase
{
    private static readonly Regex EmailRegex = new(@"^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$");

    [HttpGet("email")]
    public IActionResult ValidateEmail([FromQuery] string email)
    {
        bool isValid = EmailRegex.IsMatch(email);
        return Ok(new { valid = isValid });
    }
}
