// Corrigé : NoSQL Injection (CWE-943)
// Les champs sont typés strictement en chaîne (aucun objet/opérateur ne
// peut être désérialisé à la place d'une valeur scalaire), et la requête
// utilise le filtre typé du driver plutôt qu'un BsonDocument arbitraire.
using Microsoft.AspNetCore.Mvc;
using MongoDB.Bson;
using MongoDB.Driver;

public class LoginRequest
{
    // Type strict "string" : la désérialisation JSON échoue si le client
    // envoie un objet ({"$ne": null}) au lieu d'une chaîne.
    public string Username { get; set; } = string.Empty;
    public string Password { get; set; } = string.Empty;
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
        if (string.IsNullOrWhiteSpace(request.Username) || string.IsNullOrWhiteSpace(request.Password))
            return BadRequest("Format invalide.");

        // Filtre typé construit avec les builders du driver : les valeurs
        // sont traitées comme des littéraux, jamais comme des opérateurs.
        var filter = Builders<BsonDocument>.Filter.Eq("username", request.Username)
                   & Builders<BsonDocument>.Filter.Eq("password", request.Password);

        var user = _users.Find(filter).FirstOrDefault();
        return user == null ? Unauthorized() : Ok(user["username"]);
    }
}
