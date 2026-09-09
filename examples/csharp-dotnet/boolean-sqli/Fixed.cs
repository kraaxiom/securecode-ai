// Corrigé : Boolean-based SQL Injection (CWE-89)
// La valeur utilisateur est liée via un paramètre SQL, ce qui empêche
// toute altération de la logique booléenne de la requête.
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
            "SELECT Id, Name, Price FROM Products WHERE Name = @Name", connection);
        command.Parameters.AddWithValue("@Name", name);

        using var reader = command.ExecuteReader();
        var results = new List<object>();
        while (reader.Read())
            results.Add(new { Id = reader.GetInt32(0), Name = reader.GetString(1) });

        return Ok(results);
    }
}
