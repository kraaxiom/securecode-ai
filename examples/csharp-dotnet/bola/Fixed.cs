// Corrigé : Broken Object Level Authorization (CWE-639)
// La requête filtre directement sur le propriétaire de la ressource, garantissant
// qu'un utilisateur ne peut récupérer que ses propres objets.
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/orders")]
[Authorize]
public class OrdersController : ControllerBase
{
    private readonly AppDbContext _db;

    public OrdersController(AppDbContext db)
    {
        _db = db;
    }

    [HttpGet("{id:int}")]
    public IActionResult GetOrder(int id)
    {
        var currentUserId = User.GetUserId(); // identifiant issu du principal authentifié, pas du client

        // Clause d'appartenance appliquée directement dans la requête de données.
        var order = _db.Orders.FirstOrDefault(o => o.Id == id && o.OwnerId == currentUserId);
        if (order is null)
            return NotFound();

        return Ok(order);
    }
}
