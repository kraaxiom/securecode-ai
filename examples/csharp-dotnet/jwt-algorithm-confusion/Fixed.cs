// Corrigé : JWT Algorithm Confusion (CWE-347)
// L'algorithme de vérification attendu (RS256) est explicitement imposé via
// ValidAlgorithms, sans jamais le déduire du token. Toute tentative de signer
// avec HS256 en réutilisant la clé publique RSA comme secret est rejetée.
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;

[ApiController]
[Route("api/resource")]
public class ResourceController : ControllerBase
{
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
            ValidateAudience = false,
            // Liste fermée : seul RS256 est accepté, exclut explicitement HS256.
            ValidAlgorithms = new[] { SecurityAlgorithms.RsaSha256 }
        };

        try
        {
            var principal = handler.ValidateToken(token, validationParameters, out var validatedToken);

            // Double vérification défensive de l'algorithme réellement utilisé.
            if (validatedToken is not JwtSecurityToken jwt ||
                !jwt.Header.Alg.Equals(SecurityAlgorithms.RsaSha256, StringComparison.Ordinal))
            {
                return Unauthorized(new { error = "Algorithme de signature non autorisé." });
            }

            return Ok(new { user = principal.Identity?.Name });
        }
        catch (SecurityTokenException)
        {
            return Unauthorized(new { error = "Token invalide." });
        }
    }
}
