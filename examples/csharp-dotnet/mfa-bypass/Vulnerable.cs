// Vulnérable : Contournement de MFA (CWE-287)
// Un token pleinement privilégié est émis dès que le mot de passe est validé,
// avant toute vérification du second facteur. L'étape MFA suivante n'est
// qu'indicative : le client possède déjà un accès complet.
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

        // Token complet émis immédiatement, indépendamment du statut MFA du compte.
        var fullToken = _userService.IssueFullToken(user);

        return Ok(new
        {
            token = fullToken,
            mfaRequired = user.MfaEnabled // simple indication, non contraignante
        });
    }
}

public record LoginRequest(string Email, string Password);
