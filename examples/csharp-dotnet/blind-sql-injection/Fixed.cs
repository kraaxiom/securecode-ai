// Corrigé : Blind SQL Injection (CWE-89)
// Utilisation d'une requête paramétrée : la valeur utilisateur est liée en
// tant que paramètre SQL et n'est jamais concaténée dans le texte de la
// requête, ce qui empêche toute inférence booléenne ou temporelle.
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
            "SELECT 1 FROM Users WHERE Username = @Username AND Active = 1", connection);
        command.Parameters.AddWithValue("@Username", username);

        using var reader = command.ExecuteReader();
        // Réponse et temps de traitement uniformes, quel que soit le résultat.
        return Ok(new { exists = reader.HasRows });
    }
}
