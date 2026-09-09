// Vulnérable : Politique de mot de passe faible (CWE-521)
// Seule une longueur minimale insuffisante (6 caractères) est exigée, sans
// vérification contre les mots de passe compromis connus, facilitant le
// brute force et le credential stuffing.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/account")]
public class RegistrationController : ControllerBase
{
    private readonly IUserService _userService;

    public RegistrationController(IUserService userService)
    {
        _userService = userService;
    }

    [HttpPost("register")]
    public async Task<IActionResult> Register([FromBody] RegisterRequest request)
    {
        if (request.Password.Length < 6)
        {
            return BadRequest(new { error = "Le mot de passe doit contenir au moins 6 caractères" });
        }

        await _userService.CreateAccountAsync(request.Email, request.Password);
        return Ok();
    }
}

public record RegisterRequest(string Email, string Password);
