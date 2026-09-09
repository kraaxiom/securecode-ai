// Vulnérable : Stacked Query SQL Injection (CWE-89)
// Le driver SQL Server autorise nativement l'exécution de plusieurs
// instructions séparées par un point-virgule dans un seul appel. La valeur
// utilisateur est concaténée sans paramétrage, permettant à un attaquant
// d'ajouter une instruction SQL complètement distincte (ex: DROP TABLE).
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/users")]
public class UserUpdateController : ControllerBase
{
    private readonly string _connectionString;

    public UserUpdateController(IConfiguration config) =>
        _connectionString = config.GetConnectionString("Default")!;

    [HttpPost("rename")]
    public IActionResult Rename([FromForm] string newName)
    {
        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        // "newName" contenant "x'; DROP TABLE users; --" exécute une
        // seconde instruction distincte.
        var command = new SqlCommand(
            $"UPDATE users SET name = '{newName}' WHERE id = 1", connection);
        command.ExecuteNonQuery();

        return Ok();
    }
}
