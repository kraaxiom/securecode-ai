// Vulnérable : Blind SQL Injection (CWE-89)
// La requête SQL est construite par concaténation d'une entrée utilisateur.
// Le résultat n'est pas affiché directement (seulement true/false), mais un
// attaquant peut quand même déduire des informations via des conditions
// booléennes ou des délais de réponse injectés dans la chaîne.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/users")]
public class UserCheckController : ControllerBase
{
    private readonly string _connectionString;

    public UserCheckController(string connectionString) => _connectionString = connectionString;

    [HttpGet("exists")]
    public IActionResult Exists([FromQuery] string username)
    {
        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            $"SELECT 1 FROM Users WHERE Username = '{username}' AND Active = 1", connection);

        using var reader = command.ExecuteReader();
        return Ok(new { exists = reader.HasRows });
    }
}
