// Vulnérable : JWT signé avec un secret faible (CWE-326)
// Le secret HMAC utilisé pour signer les tokens est une chaîne courte et
// prévisible, codée en dur dans le service. Un attaquant peut le retrouver
// hors ligne par attaque par dictionnaire puis forger des tokens valides.
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

public class TokenService
{
    // Secret court et prévisible, en dur dans le code source.
    private const string JwtSecret = "monsecret123";

    public string IssueToken(string userId, string role)
    {
        var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(JwtSecret));
        var credentials = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

        var claims = new[]
        {
            new Claim(ClaimTypes.NameIdentifier, userId),
            new Claim(ClaimTypes.Role, role)
        };

        var token = new JwtSecurityToken(
            claims: claims,
            expires: DateTime.UtcNow.AddHours(1),
            signingCredentials: credentials);

        return new JwtSecurityTokenHandler().WriteToken(token);
    }
}
