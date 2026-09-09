// Corrigé : Contournement de MFA (CWE-287)
// Aucun token pleinement privilégié n'est émis avant validation serveur du
// second facteur. Un token intermédiaire à portée limitée ('mfa_pending')
// est utilisé entre les deux étapes, et n'autorise que l'appel de vérification.
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

        if (user.MfaEnabled)
        {
            // Token intermédiaire à portée limitée : n'autorise que /mfa/verify.
            var partialToken = _userService.IssuePartialToken(user, scope: "mfa_pending");
            return Ok(new { mfaRequired = true, partialToken });
        }

        var fullToken = _userService.IssueFullToken(user);
        return Ok(new { token = fullToken });
    }

    [HttpPost("mfa/verify")]
    [ServiceFilter(typeof(RequirePartialTokenFilter))]
    public async Task<IActionResult> VerifyMfa([FromBody] MfaVerifyRequest request)
    {
        // HttpContext.User a été peuplé par RequirePartialTokenFilter à partir
        // du token intermédiaire, jamais depuis un paramètre client modifiable.
        var user = await _userService.GetCurrentPendingUserAsync(HttpContext);
        if (user == null || !await _userService.VerifyMfaCodeAsync(user, request.Code))
        {
            return Unauthorized(new { error = "Code invalide" });
        }

        // Session/token complet uniquement après validation côté serveur du second facteur.
        var fullToken = _userService.IssueFullToken(user);
        return Ok(new { token = fullToken });
    }
}

public record LoginRequest(string Email, string Password);
public record MfaVerifyRequest(string Code);
