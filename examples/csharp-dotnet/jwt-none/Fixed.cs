// Corrigé : JWT `alg: none` (CWE-347)
// ValidateToken est utilisé avec une liste fermée d'algorithmes forts, excluant
// structurellement 'none'. Aucune décision d'autorisation ne repose plus sur un
// décodage non vérifié cryptographiquement.
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;

[ApiController]
[Route("api/resource")]
public class ProfileController : ControllerBase
{
    private readonly SymmetricSecurityKey _signingKey;

    public ProfileController(SymmetricSecurityKey signingKey)
    {
        _signingKey = signingKey;
    }

    [HttpGet("profile")]
    public IActionResult GetProfile([FromHeader(Name = "Authorization")] string authorization)
    {
        var token = authorization?.Replace("Bearer ", "");
        var handler = new JwtSecurityTokenHandler();

        var validationParameters = new TokenValidationParameters
        {
            IssuerSigningKey = _signingKey,
            ValidateIssuer = false,
            ValidateAudience = false,
            // 'none' est structurellement exclu de cette liste fermée.
            ValidAlgorithms = new[] { SecurityAlgorithms.HmacSha256 }
        };

        ClaimsPrincipal principal;
        try
        {
            principal = handler.ValidateToken(token, validationParameters, out _);
        }
        catch (SecurityTokenException)
        {
            return Unauthorized(new { error = "Token invalide ou non signé." });
        }

        var role = principal.FindFirst("role")?.Value;
        if (role == "admin")
        {
            return Ok(new { data = "Informations administrateur sensibles" });
        }

        return Ok(new { data = "Profil standard" });
    }
}
