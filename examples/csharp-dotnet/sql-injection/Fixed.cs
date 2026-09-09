// Corrigé : SQL Injection générique (CWE-89)
// La requête utilise un paramètre lié : la valeur utilisateur n'est jamais
// concaténée dans le texte SQL, elle est transmise séparément au driver.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/users")]
public class UserSearchController : ControllerBase
{
    private readonly string _connectionString;

    public UserSearchController(IConfiguration config) =>
        _connectionString = config.GetConnectionString("Default")!;

    [HttpGet("search")]
    public IActionResult Search([FromQuery] string name)
    {
        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            "SELECT id, name, email FROM users WHERE name = @name", connection);
        command.Parameters.Add("@name", System.Data.SqlDbType.NVarChar, 100).Value = name;

        using var reader = command.ExecuteReader();
        var results = new List<object>();
        while (reader.Read())
            results.Add(new { Id = reader.GetInt32(0), Name = reader.GetString(1) });

        return Ok(results);
    }
}
