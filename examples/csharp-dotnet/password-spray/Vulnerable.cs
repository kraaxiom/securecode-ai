// Vulnérable : Password Spraying (CWE-307)
// La limitation des échecs de connexion n'est appliquée que par compte : un
// attaquant testant un petit nombre de mots de passe courants sur un grand
// nombre de comptes distincts depuis une même source passe inaperçu.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IUserService _userService;
    private readonly IAttemptStore _attemptStore;

    public LoginController(IUserService userService, IAttemptStore attemptStore)
    {
        _userService = userService;
        _attemptStore = attemptStore;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginRequest request)
    {
        // Limitation uniquement par compte, aucune agrégation par IP/source.
        var accountAttempts = await _attemptStore.GetFailedAttemptsAsync(request.Email);
        if (accountAttempts > 5)
        {
            return StatusCode(429, new { error = "Trop de tentatives" });
        }

        var user = await _userService.FindByEmailAsync(request.Email);
        if (user == null || !_userService.VerifyPassword(user, request.Password))
        {
            await _attemptStore.RecordFailureAsync(request.Email);
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        var token = _userService.IssueToken(user);
        return Ok(new { token });
    }
}

public record LoginRequest(string Email, string Password);
