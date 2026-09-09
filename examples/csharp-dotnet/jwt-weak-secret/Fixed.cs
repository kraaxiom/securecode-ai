// Corrigé : JWT signé avec un secret faible (CWE-326)
// Le secret est chargé depuis un gestionnaire de secrets (variable d'environnement
// injectée par le vault), et sa longueur minimale (256 bits) est vérifiée
// explicitement au démarrage plutôt que d'utiliser une valeur en dur.
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

public class TokenService
{
    private readonly byte[] _jwtSecret;

    public TokenService(IConfiguration configuration)
    {
        var secret = configuration["JWT_SECRET"];
        if (string.IsNullOrEmpty(secret) || Encoding.UTF8.GetByteCount(secret) < 32)
        {
            // 256 bits minimum requis pour HS256 : échec rapide au démarrage.
            throw new InvalidOperationException("JWT_SECRET manquant ou insuffisant (256 bits minimum requis).");
        }

        _jwtSecret = Encoding.UTF8.GetBytes(secret);
    }

    public string IssueToken(string userId, string role)
    {
        var key = new SymmetricSecurityKey(_jwtSecret);
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
