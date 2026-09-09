// Vulnérable : NoSQL Injection (CWE-943)
// Les identifiants sont désérialisés depuis le JSON brut et transmis tels
// quels comme filtre MongoDB, permettant à un attaquant d'injecter un objet
// opérateur (ex: {"$ne": null}) au lieu d'une valeur scalaire attendue.
using Microsoft.AspNetCore.Mvc;
using MongoDB.Bson;
using MongoDB.Driver;

public class LoginRequest
{
    public BsonValue Username { get; set; } = default!;
    public BsonValue Password { get; set; } = default!;
}

[ApiController]
[Route("api/auth")]
public class LoginController : ControllerBase
{
    private readonly IMongoCollection<BsonDocument> _users;

    public LoginController(IMongoCollection<BsonDocument> users) => _users = users;

    [HttpPost("login")]
    public IActionResult Login([FromBody] LoginRequest request)
    {
        // request.Username / request.Password peuvent être des objets JSON
        // arbitraires (ex: {"$ne": null}) et non de simples chaînes.
        var filter = new BsonDocument
        {
            { "username", request.Username },
            { "password", request.Password }
        };

        var user = _users.Find(filter).FirstOrDefault();
        return user == null ? Unauthorized() : Ok(user["username"]);
    }
}
