// Corrigé : ReDoS - Regular Expression Denial of Service (CWE-1333)
// La regex est réécrite pour supprimer les quantificateurs imbriqués
// (élimine l'ambiguïté de correspondance), une limite de longueur est
// imposée avant tout test, et un timeout d'exécution protège le moteur.
using Microsoft.AspNetCore.Mvc;
using System.Text.RegularExpressions;

[ApiController]
[Route("api/validation")]
public class EmailValidationController : ControllerBase
{
    private const int MaxEmailLength = 254;

    // Regex non ambiguë (pas de groupe répété contenant lui-même un
    // quantificateur) + timeout d'exécution garanti par le moteur .NET.
    private static readonly Regex EmailRegex = new(
        @"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$",
        RegexOptions.Compiled,
        TimeSpan.FromMilliseconds(200));

    [HttpGet("email")]
    public IActionResult ValidateEmail([FromQuery] string email)
    {
        if (string.IsNullOrEmpty(email) || email.Length > MaxEmailLength)
            return BadRequest("Format d'email invalide.");

        try
        {
            bool isValid = EmailRegex.IsMatch(email);
            return Ok(new { valid = isValid });
        }
        catch (RegexMatchTimeoutException)
        {
            // Le moteur a dépassé le budget de temps alloué : entrée rejetée
            // plutôt que de bloquer le thread.
            return BadRequest("Format d'email invalide.");
        }
    }
}
