// Vulnérable : Brute Force (CWE-307)
// L'endpoint de connexion n'applique aucune limitation du nombre de tentatives
// par compte ni par IP, permettant à un attaquant d'essayer un grand nombre
// de mots de passe de façon automatisée jusqu'à en trouver un valide.
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
            // Aucun compteur d'échecs, aucun verrouillage, aucun rate limiting.
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        var token = _userService.IssueToken(user);
        return Ok(new { token });
    }
}

public record LoginRequest(string Email, string Password);
