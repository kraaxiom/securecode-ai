// Vulnérable : HTTP Response Splitting (CWE-113)
// La cible de redirection fournie par l'utilisateur est écrite directement
// dans l'en-tête Location, permettant l'injection de CR/LF pour scinder
// la réponse HTTP et injecter du contenu supplémentaire.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("account")]
public class LoginRedirectController : ControllerBase
{
    [HttpGet("post-login")]
    public IActionResult PostLogin([FromQuery] string next)
    {
        Response.Headers["Location"] = next;
        return StatusCode(302);
    }
}
