// Corrigé : Stacked Query SQL Injection (CWE-89)
// La requête utilise un paramètre lié : la valeur utilisateur ne peut plus
// terminer l'instruction en cours ni en injecter une nouvelle. Le compte
// de connexion applique en complément le moindre privilège (pas de droits
// DDL sur ce compte applicatif).
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
        if (string.IsNullOrWhiteSpace(newName) || newName.Length > 100)
            return BadRequest("Nom invalide.");

        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            "UPDATE users SET name = @name WHERE id = 1", connection);
        command.Parameters.Add("@name", System.Data.SqlDbType.NVarChar, 100).Value = newName;
        command.ExecuteNonQuery();

        return Ok();
    }
}
