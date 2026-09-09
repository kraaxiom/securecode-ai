// Vulnérable : UNION-based SQL Injection (CWE-89)
// L'identifiant produit est concaténé directement dans la requête,
// permettant à un attaquant d'ajouter une clause UNION SELECT pour
// extraire des données d'autres tables (ex: table des utilisateurs).
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
        using var connection = new SqlConnection(_connectionString);
        connection.Open();

        var command = new SqlCommand(
            $"SELECT name, price FROM products WHERE id = {id}", connection);

        using var reader = command.ExecuteReader();
        if (!reader.Read())
            return NotFound();

        return Ok(new { Name = reader.GetString(0), Price = reader.GetDecimal(1) });
    }
}
