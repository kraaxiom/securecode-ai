// Corrigé : Credential Stuffing (CWE-307)
// Une détection de vélocité globale par IP (tous comptes confondus) complète
// le contrôle par compte, et la MFA est exigée après authentification réussie
// pour neutraliser l'usage de couples identifiant/mot de passe volés.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IUserService _userService;
    private readonly IVelocityTracker _velocityTracker;
    private readonly ISecurityAlertService _alertService;
    private const int SuspiciousIpThreshold = 30;

    public LoginController(
        IUserService userService,
        IVelocityTracker velocityTracker,
        ISecurityAlertService alertService)
    {
        _userService = userService;
        _velocityTracker = velocityTracker;
        _alertService = alertService;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginRequest request)
    {
        var clientIp = HttpContext.Connection.RemoteIpAddress?.ToString() ?? "unknown";

        // Vélocité globale : volume de tentatives depuis cette IP, tous comptes confondus.
        var velocity = await _velocityTracker.GetFailedAttemptVelocityAsync(clientIp);
        if (velocity > SuspiciousIpThreshold)
        {
            await _alertService.FlagSuspiciousSourceAsync(clientIp, "credential_stuffing_suspected");
            return StatusCode(429, new { error = "Trafic anormal détecté, réessayez plus tard." });
        }

        var user = await _userService.FindByEmailAsync(request.Email);
        if (user == null || !_userService.VerifyPassword(user, request.Password))
        {
            await _velocityTracker.RecordFailureAsync(clientIp);
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        if (user.MfaEnabled)
        {
            var challengeId = await _userService.CreateMfaChallengeAsync(user);
            return Ok(new { mfaRequired = true, challengeId });
        }

        var token = _userService.IssueToken(user);
        return Ok(new { token });
    }
}

public interface IVelocityTracker
{
    Task<int> GetFailedAttemptVelocityAsync(string ip);
    Task RecordFailureAsync(string ip);
}

public interface ISecurityAlertService
{
    Task FlagSuspiciousSourceAsync(string ip, string reason);
}

public record LoginRequest(string Email, string Password);
