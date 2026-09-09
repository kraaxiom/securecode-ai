// Vulnérable : Mauvaise configuration OAuth (CWE-287)
// Aucun paramètre 'state' n'est généré ni vérifié au retour du callback, et
// l'ID token reçu du fournisseur n'est pas validé (issuer/audience/signature),
// exposant l'application à une CSRF sur le flux d'autorisation.
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/oauth")]
public class OAuthController : ControllerBase
{
    private readonly IOAuthClient _oauthClient;

    public OAuthController(IOAuthClient oauthClient)
    {
        _oauthClient = oauthClient;
    }

    [HttpGet("login")]
    public IActionResult Login()
    {
        // Aucun 'state' généré/stocké avant redirection vers le fournisseur.
        var authorizationUrl = _oauthClient.GetAuthorizationUrl();
        return Redirect(authorizationUrl);
    }

    [HttpGet("callback")]
    public async Task<IActionResult> Callback([FromQuery] string code)
    {
        // 'state' non vérifié, ID token utilisé sans validation stricte.
        var tokenSet = await _oauthClient.ExchangeCodeAsync(code);
        return Ok(new { claims = tokenSet.Claims });
    }
}
