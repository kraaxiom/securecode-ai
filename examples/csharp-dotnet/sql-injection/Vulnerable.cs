// Vulnérable : SQL Injection générique (CWE-89)
// Le nom fourni par l'utilisateur est concaténé directement dans la
// requête SQL, permettant à un attaquant d'altérer la logique de la
// requête exécutée par la base de données.
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
            $"SELECT id, name, email FROM users WHERE name = '{name}'", connection);

        using var reader = command.ExecuteReader();
        var results = new List<object>();
        while (reader.Read())
            results.Add(new { Id = reader.GetInt32(0), Name = reader.GetString(1) });

        return Ok(results);
    }
}
