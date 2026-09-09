// Vulnérable : Boolean-based SQL Injection (CWE-89)
// La clause WHERE est construite par concaténation directe de l'entrée
// utilisateur, permettant à un attaquant d'altérer la logique booléenne
// de la requête (ex: injection d'un OR toujours vrai).
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/products")]
public class ProductSearchController : ControllerBase
{
    private readonly string _connectionString;

    public ProductSearchController(string connectionString) => _connectionString = connectionString;

    [HttpGet("search")]
    public IActionResult Search([FromQuery] string name)
    {
        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            $"SELECT Id, Name, Price FROM Products WHERE Name = '{name}'", connection);

        using var reader = command.ExecuteReader();
        var results = new List<object>();
        while (reader.Read())
            results.Add(new { Id = reader.GetInt32(0), Name = reader.GetString(1) });

        return Ok(results);
    }
}
