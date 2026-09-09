// Corrigé : UNION-based SQL Injection (CWE-89)
// L'identifiant est validé/typé en entier avant utilisation, et la requête
// utilise un paramètre lié : aucune clause UNION ne peut être injectée.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/products")]
public class ProductController : ControllerBase
{
    private readonly string _connectionString;

    public ProductController(IConfiguration config) =>
        _connectionString = config.GetConnectionString("Default")!;

    [HttpGet("{id}")]
    public IActionResult GetProduct(string id)
    {
        if (!int.TryParse(id, out var productId))
            return BadRequest("Identifiant invalide.");

        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            "SELECT name, price FROM products WHERE id = @id", connection);
        command.Parameters.Add("@id", System.Data.SqlDbType.Int).Value = productId;

        using var reader = command.ExecuteReader();
        if (!reader.Read())
            return NotFound();

        return Ok(new { Name = reader.GetString(0), Price = reader.GetDecimal(1) });
    }
}
