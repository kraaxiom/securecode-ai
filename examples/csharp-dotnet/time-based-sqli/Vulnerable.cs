// Vulnérable : Time-Based Blind SQL Injection (CWE-89)
// La requête est concaténée sans paramétrage, et les erreurs sont masquées
// par une gestion générique, empêchant l'attaquant d'observer directement
// des données mais lui permettant d'inférer des informations via le délai
// de réponse (ex: injection de WAITFOR DELAY).
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
        try
        {
            using var connection = new SqlConnection(_connectionString);
            connection.Open();

            var command = new SqlCommand(
                $"SELECT id, status FROM orders WHERE id = {id}", connection);
            using var reader = command.ExecuteReader();

            return reader.Read() ? Ok(new { Id = reader.GetInt32(0) }) : NotFound();
        }
        catch
        {
            // Réponse générique : ne révèle aucune donnée, mais n'empêche
            // pas la mesure du délai d'exécution induit par l'injection.
            return StatusCode(500, "Erreur interne.");
        }
    }
}
