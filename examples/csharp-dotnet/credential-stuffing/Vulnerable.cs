// Vulnérable : Credential Stuffing (CWE-307)
// Aucune corrélation n'est effectuée entre les tentatives de connexion échouées
// sur des comptes différents depuis une même origine, et la MFA n'est jamais
// proposée : des paires identifiant/mot de passe volées peuvent être testées
// en masse sans détection.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IUserService _userService;

    public LoginController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginRequest request)
    {
        var user = await _userService.FindByEmailAsync(request.Email);
        if (user == null || !_userService.VerifyPassword(user, request.Password))
        {
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        // Aucune vérification MFA même si le compte est sensible.
        var token = _userService.IssueToken(user);
        return Ok(new { token });
    }
}

public record LoginRequest(string Email, string Password);
