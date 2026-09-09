// Vulnérable : JWT Algorithm Confusion (CWE-347)
// La vérification du token laisse la bibliothèque déduire l'algorithme depuis
// le header du JWT plutôt que de l'imposer côté serveur. Un attaquant peut
// forger un token HS256 signé avec la clé publique RSA réutilisée comme secret.
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;

[ApiController]
[Route("api/resource")]
public class ResourceController : ControllerBase
{
    // Clé publique RSA exposée pour la vérification des tokens émis par l'IdP.
    private readonly RsaSecurityKey _publicKey;

    public ResourceController(RsaSecurityKey publicKey)
    {
        _publicKey = publicKey;
    }

    [HttpGet("secure-data")]
    public IActionResult GetSecureData([FromHeader(Name = "Authorization")] string authorization)
    {
        var token = authorization?.Replace("Bearer ", "");
        var handler = new JwtSecurityTokenHandler();

        var validationParameters = new TokenValidationParameters
        {
            IssuerSigningKey = _publicKey,
            ValidateIssuer = false,
            ValidateAudience = false
            // Aucun ValidAlgorithms : l'algorithme est déduit du header du token,
            // ce qui permet de basculer vers HS256 avec la clé publique comme secret.
        };

        var principal = handler.ValidateToken(token, validationParameters, out _);
        return Ok(new { user = principal.Identity?.Name });
    }
}
