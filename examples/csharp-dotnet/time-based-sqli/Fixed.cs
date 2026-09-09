// Corrigé : Time-Based Blind SQL Injection (CWE-89)
// Un paramètre lié élimine toute la classe de vulnérabilité (pas seulement
// la variante "time-based"). Un timeout de commande strict limite en outre
// l'impact d'une éventuelle injection résiduelle ailleurs dans le code.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/orders")]
public class OrderController : ControllerBase
{
    private readonly string _connectionString;

    public OrderController(IConfiguration config) =>
        _connectionString = config.GetConnectionString("Default")!;

    [HttpGet("{id}")]
    public IActionResult GetOrder(string id)
    {
        if (!int.TryParse(id, out var orderId))
            return BadRequest("Identifiant invalide.");

        try
        {
            using var connection = new SqlConnection(_connectionString);
            connection.Open();

            var command = new SqlCommand(
                "SELECT id, status FROM orders WHERE id = @id", connection)
            {
                CommandTimeout = 5 // limite l'exécution à 5 secondes
            };
            command.Parameters.Add("@id", System.Data.SqlDbType.Int).Value = orderId;

            using var reader = command.ExecuteReader();
            return reader.Read() ? Ok(new { Id = reader.GetInt32(0) }) : NotFound();
        }
        catch
        {
            return StatusCode(500, "Erreur interne.");
        }
    }
}
