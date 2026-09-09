// Corrigé : Politique de mot de passe faible (CWE-521)
// La longueur minimale requise est portée à 12 caractères, sans règle de
// composition artificielle, et le mot de passe est vérifié contre une liste
// de mots de passe compromis connus (service type "have I been pwned").
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/account")]
public class RegistrationController : ControllerBase
{
    private readonly IUserService _userService;
    private readonly IPwnedPasswordChecker _pwnedChecker;

    private const int MinimumLength = 12;

    public RegistrationController(IUserService userService, IPwnedPasswordChecker pwnedChecker)
    {
        _userService = userService;
        _pwnedChecker = pwnedChecker;
    }

    [HttpPost("register")]
    public async Task<IActionResult> Register([FromBody] RegisterRequest request)
    {
        if (request.Password.Length < MinimumLength)
        {
            return BadRequest(new { error = $"Le mot de passe doit contenir au moins {MinimumLength} caractères" });
        }

        // Vérification via k-anonymity contre une liste de mots de passe compromis connus.
        if (await _pwnedChecker.IsCompromisedAsync(request.Password))
        {
            return BadRequest(new
            {
                error = "Ce mot de passe a été trouvé dans une fuite de données connue, choisissez-en un autre."
            });
        }

        // Aucune règle de composition artificielle ni de rotation périodique forcée.
        await _userService.CreateAccountAsync(request.Email, request.Password);
        return Ok();
    }
}

public interface IPwnedPasswordChecker
{
    Task<bool> IsCompromisedAsync(string password);
}

public record RegisterRequest(string Email, string Password);
