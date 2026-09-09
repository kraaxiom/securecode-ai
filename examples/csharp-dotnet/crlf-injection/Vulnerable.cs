// Vulnérable : CRLF Injection (CWE-93)
// La valeur de redirection fournie par l'utilisateur est écrite directement
// dans l'en-tête Location, permettant l'injection de séquences \r\n pour
// falsifier des en-têtes ou scinder la réponse HTTP.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("account")]
public class RedirectController : ControllerBase
{
    [HttpGet("redirect")]
    public IActionResult RedirectTo([FromQuery] string next)
    {
        Response.Headers["Location"] = next;
        return StatusCode(302);
    }
}
