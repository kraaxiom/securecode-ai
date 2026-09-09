// Corrigé : Error-based SQL Injection (CWE-89)
// Requête paramétrée avec identifiant validé/casté, et messages d'erreur
// génériques renvoyés au client ; le détail est journalisé côté serveur.
using Microsoft.AspNetCore.Mvc;
using System.Data.SqlClient;

[ApiController]
[Route("api/orders")]
public class OrderController : ControllerBase
{
    private readonly string _connectionString;
    private readonly ILogger<OrderController> _logger;

    public OrderController(string connectionString, ILogger<OrderController> logger)
    {
        _connectionString = connectionString;
        _logger = logger;
    }

    [HttpGet("{id:int}")]
    public IActionResult GetOrder(int id)
    {
        try
        {
            using var connection = new SqlConnection(_connectionString);
            connection.Open();
            var command = new SqlCommand("SELECT * FROM Orders WHERE Id = @Id", connection);
            command.Parameters.AddWithValue("@Id", id);
            using var reader = command.ExecuteReader();
            if (!reader.Read())
                return NotFound();
            return Ok(new { id = reader["Id"], total = reader["Total"] });
        }
        catch (SqlException ex)
        {
            _logger.LogError(ex, "Erreur lors de la récupération de la commande {Id}", id);
            return StatusCode(500, new { error = "Une erreur est survenue." });
        }
    }
}
