// Corrigé : Brute Force (CWE-307)
// La tentative de connexion est limitée par compte ET par IP, avec verrouillage
// progressif et journalisation des échecs, conformément à la remédiation
// documentée pour empêcher les essais automatisés de mots de passe.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IUserService _userService;
    private readonly ILoginAttemptTracker _attemptTracker;
    private readonly ILogger<LoginController> _logger;

    private const int MaxAttempts = 5;
    private static readonly TimeSpan LockoutWindow = TimeSpan.FromMinutes(15);

    public LoginController(
        IUserService userService,
        ILoginAttemptTracker attemptTracker,
        ILogger<LoginController> logger)
    {
        _userService = userService;
        _attemptTracker = attemptTracker;
        _logger = logger;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginRequest request)
    {
        var clientIp = HttpContext.Connection.RemoteIpAddress?.ToString() ?? "unknown";
        var accountKey = $"login:account:{request.Email.ToLowerInvariant()}";
        var ipKey = $"login:ip:{clientIp}";

        if (await _attemptTracker.IsLockedAsync(accountKey, MaxAttempts, LockoutWindow) ||
            await _attemptTracker.IsLockedAsync(ipKey, MaxAttempts * 10, LockoutWindow))
        {
            _logger.LogWarning("Tentative de connexion bloquée (verrouillage actif) pour {Ip}", clientIp);
            return StatusCode(429, new { error = "Trop de tentatives. Réessayez plus tard." });
        }

        var user = await _userService.FindByEmailAsync(request.Email);
        var valid = user != null && _userService.VerifyPassword(user, request.Password);

        if (!valid)
        {
            await _attemptTracker.RecordFailureAsync(accountKey, LockoutWindow);
            await _attemptTracker.RecordFailureAsync(ipKey, LockoutWindow);
            _logger.LogWarning("Échec de connexion pour {Email} depuis {Ip}", request.Email, clientIp);
            // Message générique : ne permet pas de distinguer un compte existant d'un compte inexistant.
            return Unauthorized(new { error = "Identifiants invalides" });
        }

        await _attemptTracker.ClearAsync(accountKey);
        var token = _userService.IssueToken(user!);
        return Ok(new { token });
    }
}

public interface ILoginAttemptTracker
{
    Task<bool> IsLockedAsync(string key, int maxAttempts, TimeSpan window);
    Task RecordFailureAsync(string key, TimeSpan window);
    Task ClearAsync(string key);
}

public record LoginRequest(string Email, string Password);
