// Corrigé : Password Spraying (CWE-307)
// Une agrégation globale des échecs par IP/source, tous comptes confondus,
// complète la limitation par compte et déclenche une alerte/blocage lorsque
// le volume dépasse un seuil anormal, caractéristique d'un spray distribué.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IUserService _userService;
    private readonly IAttemptStore _attemptStore;
    private readonly ISecurityAlertService _alertService;

    private const int AccountThreshold = 5;
    private const int IpThreshold = 50;

    public LoginController(
        IUserService userService,
        IAttemptStore attemptStore,
        ISecurityAlertService alertService)
    {
        _userService = userService;
        _attemptStore = attemptStore;
        _alertService = alertService;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginRequest request)
    {
        var clientIp = HttpContext.Connection.RemoteIpAddress?.ToString() ?? "unknown";

        var accountAttempts = await _attemptStore.GetFailedAttemptsAsync(request.Email);
        var ipAttempts = await _attemptStore.GetFailedAttemptsByIpAsync(clientIp);

        if (accountAttempts > AccountThreshold || ipAttempts > IpThreshold)
        {
            if (ipAttempts > IpThreshold)
            {
                await _alertService.FlagSuspiciousSourceAsync(clientIp, "password_spray_suspected");
            }
            return StatusCode(429, new { error = "Trop de tentatives." });
        }

        var user = await _userService.FindByEmailAsync(request.Email);
        if (user == null || !_userService.VerifyPassword(user, request.Password))
        {
            await _attemptStore.RecordFailureAsync(request.Email);
            await _attemptStore.RecordFailureByIpAsync(clientIp);
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        var token = _userService.IssueToken(user);
        return Ok(new { token });
    }
}

public record LoginRequest(string Email, string Password);
