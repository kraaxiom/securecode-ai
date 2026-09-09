// Vulnérable : JWT `alg: none` (CWE-347)
// Le token est simplement lu (ReadJwtToken) sans jamais appeler ValidateToken,
// donc aucune signature n'est vérifiée. Un attaquant peut forger un token avec
// alg: none, une signature vide, et un contenu de claims arbitraire.
using Microsoft.AspNetCore.Mvc;
using System.IdentityModel.Tokens.Jwt;

[ApiController]
[Route("api/resource")]
public class ProfileController : ControllerBase
{
    [HttpGet("profile")]
    public IActionResult GetProfile([FromHeader(Name = "Authorization")] string authorization)
    {
        var token = authorization?.Replace("Bearer ", "");
        var handler = new JwtSecurityTokenHandler();

        // Décodage sans vérification cryptographique de la signature !
        var jwt = handler.ReadJwtToken(token);
        var role = jwt.Claims.FirstOrDefault(c => c.Type == "role")?.Value;

        if (role == "admin")
        {
            return Ok(new { data = "Informations administrateur sensibles" });
        }

        return Ok(new { data = "Profil standard" });
    }
}
