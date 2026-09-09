// Corrigé : Mauvaise configuration OAuth (CWE-287)
// Un 'state' unique est généré et stocké côté serveur avant redirection, puis
// strictement vérifié au retour. L'ID token est intégralement validé
// (signature, issuer, audience, expiration) avant toute confiance accordée.
using Microsoft.AspNetCore.Mvc;
using System.Security.Cryptography;

[ApiController]
[Route("api/oauth")]
public class OAuthController : ControllerBase
{
    private readonly IOAuthClient _oauthClient;

    // Redirect_uri enregistré en correspondance exacte, sans wildcard.
    private const string RedirectUri = "https://app.example.com/api/oauth/callback";

    public OAuthController(IOAuthClient oauthClient)
    {
        _oauthClient = oauthClient;
    }

    [HttpGet("login")]
    public IActionResult Login()
    {
        var state = GenerateSecureState();
        HttpContext.Session.SetString("oauth_state", state);

        var authorizationUrl = _oauthClient.GetAuthorizationUrl(RedirectUri, state);
        return Redirect(authorizationUrl);
    }

    [HttpGet("callback")]
    public async Task<IActionResult> Callback([FromQuery] string code, [FromQuery] string state)
    {
        var expectedState = HttpContext.Session.GetString("oauth_state");
        HttpContext.Session.Remove("oauth_state");

        if (string.IsNullOrEmpty(state) || string.IsNullOrEmpty(expectedState) ||
            !CryptographicOperations.FixedTimeEquals(
                System.Text.Encoding.UTF8.GetBytes(state),
                System.Text.Encoding.UTF8.GetBytes(expectedState)))
        {
            return Forbid("State invalide — tentative de CSRF sur le callback OAuth.");
        }

        // La bibliothèque valide intégralement l'ID token : signature, iss, aud, exp.
        var tokenSet = await _oauthClient.ExchangeCodeAsync(code, RedirectUri, validateIdToken: true);
        return Ok(new { claims = tokenSet.Claims });
    }

    private static string GenerateSecureState() =>
        Convert.ToHexString(RandomNumberGenerator.GetBytes(16));
}
