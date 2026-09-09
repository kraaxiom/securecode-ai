// Vulnérable : Error-based SQL Injection (CWE-89)
// La requête est construite par concaténation, et le message d'exception
// SQL brut est renvoyé au client, permettant l'extraction de données via
// des erreurs de syntaxe provoquées volontairement.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/orders")]
public class OrderController : ControllerBase
{
    private readonly string _connectionString;

    public OrderController(string connectionString) => _connectionString = connectionString;

    [HttpGet("{id}")]
    public IActionResult GetOrder(string id)
    {
        try
        {
            using var connection = new SqlConnection(_connectionString);
            connection.Open();
            var command = new SqlCommand($"SELECT * FROM Orders WHERE Id = {id}", connection);
            using var reader = command.ExecuteReader();
            reader.Read();
            return Ok(new { id = reader["Id"], total = reader["Total"] });
        }
        catch (SqlException ex)
        {
            return StatusCode(500, new { error = ex.Message });
        }
    }
}
